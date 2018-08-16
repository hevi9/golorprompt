package sys

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"plugin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// App Application instance running shared data and state
type App struct {
	errors int
}

// NewApp Create App instance
func NewApp() *App {
	return &App{}
}

// Errors Nunber of errors encountered. From zerolog hook
func (a *App) Errors() int {
	return a.errors
}

// Run zerolog hook for counting errors
func (a *App) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	if l >= zerolog.ErrorLevel {
		a.errors++
	}
}

func (a *App) resolvePluginPath(name string) (string, error) {
	suffixes := []string{".so", ""}
	for _, p := range Config.Paths {
		for _, s := range suffixes {
			fullpath := path.Join(p, name) + s
			if _, err := os.Stat(fullpath); err == nil {
				return fullpath, nil
			}
		}
	}
	return "", fmt.Errorf("%s not found", name)
}

// NewSegmentByNameJSON Create new segment by name with given json data
func (a *App) NewSegmentByNameJSON(name string, jsonBuf []byte) (Segment, error) {
	info, ok := segmentRegistry[name]
	if !ok {
		return nil, fmt.Errorf("%s does not exists", name)
	}
	// load segment plugin on demand, that updates registry
	segment, err := info.newWithJSONFunc(jsonBuf)
	if err != nil {
		return nil, err
	}
	log.Info().Str("seg", info.name).Msg("new segment")
	return segment, nil
}

func (a *App) buildFromJson(jsonBuf []byte) ([]*widgetS, error) {
	widgets := make([]*widgetS, 0)
	type SegmentSpec struct {
		Seg  string
		Args json.RawMessage
	}
	specs := []SegmentSpec{}
	err := json.Unmarshal(jsonBuf, &specs)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal")
		return nil, err
	}
	for _, s := range specs {
		jsonBuf, err := s.Args.MarshalJSON()
		if err != nil {
			panic("IMPOSSIBLE: cannot marshal segment args json")
		}
		segment, err := a.NewSegmentByNameJSON(s.Seg, jsonBuf)
		if err != nil {
			log.Error().Err(err).Msg("NewSegmentByNameJSON")
			continue
		}
		widgets = append(widgets, &widgetS{
			segment: segment,
		})
	}
	return widgets, nil
}

func (a *App) buildFromJsonSYMBOLVERSION(jsonBuf []byte) ([]*widgetS, error) {
	widgets := make([]*widgetS, 0)
	type SegmentSpec struct {
		Seg  string
		Args json.RawMessage
	}
	segments := []SegmentSpec{}
	err := json.Unmarshal(jsonBuf, &segments)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal")
		return nil, err
	}
	for _, s := range segments {
		pluginFile, err := a.resolvePluginPath(s.Seg)
		if err != nil {
			log.Error().
				Str("Seg", s.Seg).
				Err(err).
				Msg("resolvePluginPath")
			continue
		}
		pluginLib, err := plugin.Open(pluginFile)
		if err != nil {
			log.Error().Err(err).Msg("Open")
			continue
		}
		symbol, err := pluginLib.Lookup(SegmentEntrySymbolName)
		if err != nil {
			log.Error().Err(err).
				Str("file", pluginFile).
				Str("symbol", SegmentEntrySymbolName).
				Msg("Lookup")
			continue
		}
		newFunc := symbol.(func([]byte) Segment)
		argsBuf, _ := s.Args.MarshalJSON()
		segment := newFunc(argsBuf)
		widgets = append(widgets, &widgetS{
			segment: segment,
		})
	}
	return widgets, nil
}