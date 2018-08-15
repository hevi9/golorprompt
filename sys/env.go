package sys

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"plugin"

	"github.com/rs/zerolog/log"
)

type widgetS struct {
	segment Segment
	chunks  []Chunk
}

type App struct {
	errors int
}

func NewApp() *App {
	return &App{}
}

func (a *App) Errors() int {
	return a.errors
}

func (a *App) AddError(err error) Environment {
	a.errors++
	return a
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

func (a *App) buildFromJson(jsonBuf []byte) ([]*widgetS, error) {
	widgets := make([]*widgetS, 0)
	type SegmentSpec struct {
		Seg  string
		Args json.RawMessage
	}
	segments := []SegmentSpec{}
	err := json.Unmarshal(jsonBuf, &segments)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal")
		a.AddError(err)
		return nil, err
	}
	for _, s := range segments {
		pluginFile, err := a.resolvePluginPath(s.Seg)
		if err != nil {
			log.Error().
				Str("Seg", s.Seg).
				Err(err).
				Msg("resolvePluginPath")
			a.AddError(err)
			continue
		}
		pluginLib, err := plugin.Open(pluginFile)
		if err != nil {
			log.Error().Err(err).Msg("Open")
			a.AddError(err)
			continue
		}
		symbol, err := pluginLib.Lookup(SegmentEntrySymbolName)
		if err != nil {
			log.Error().Err(err).
				Str("file", pluginFile).
				Str("symbol", SegmentEntrySymbolName).
				Msg("Lookup")
			a.AddError(err)
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
