package sys

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// CommandPrompt Build prompt from given json spec
func CommandPrompt(app *App, jsonBuf []byte) error {
	// set shell
	Config.Shell = noneShell

	// build widges from json spec
	rootSlot, err := buildFromJSON(jsonBuf)
	if err != nil {
		log.Error().Err(err).Msg("Cannot build from json spec")
		return err
	}

	// render widgets concurrently
	wg := sync.WaitGroup{}
	var render func(s Slot)
	render = func(s Slot) {
		wg.Add(1)
		go func(s2 Slot) {
			defer wg.Done()
			s2.Render(app, 100)
		}(s)
		for _, s1 := range s.Slots() {
			render(s1)
		}
	}
	render(rootSlot)
	wg.Wait()

	// // make layout
	// lines := makeLayout(slots)

	// // print widgets
	// buf := bytes.Buffer{}
	// for idx, line := range lines {
	// 	log.Debug().Int("idx", idx).Msg("line")
	// 	buf.WriteString(Bg(Config.BgLine)) // TODO move out
	// 	for _, widgetElem := range line {
	// 		// fmt.Printf("%#v\n", widgetElem)
	// 		for _, chunk := range widgetElem.Chunks() {
	// 			buf.WriteString(Fg(chunk.Fg))
	// 			buf.WriteString(Config.Shell.escapeFunc(chunk.Text))
	// 		}
	// 	}
	// 	buf.WriteString(Rz())
	// }
	// fmt.Printf("%s", buf.String())

	return nil
}

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
		return nil, fmt.Errorf("name '%s' does not exists", name)
	}
	segment := info.newSegmentFunc()
	// log.Info().Str("segment", info.name).Msg("new segment")
	return segment, nil
}

func buildFromJSON(jsonBuf []byte) (Slot, error) {
	return &segmentSlot{}, nil
	// slots := make([]Slot, 0)

	// segmentMsgs := []json.RawMessage{}
	// err := json.Unmarshal(jsonBuf, &segmentMsgs)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Unmarshal segment messages")
	// 	return nil, err
	// }

	// for _, rawMsg := range segmentMsgs {
	// 	aSegmentSlot := segmentSlot{} // use slot
	// 	err := json.Unmarshal(rawMsg, &aSegmentSlot)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Unmarshall common segment")
	// 		continue
	// 	}
	// 	// decode bg color
	// 	segment, err := a.NewSegmentByNameJSON(aSegmentSlot.Name(), []byte(""))
	// 	if err != nil {
	// 		log.Error().Str("segment", aSegmentSlot.Name()).Err(err).Msg("NewSegmentByNameJSON")
	// 		continue
	// 	}
	// 	err = json.Unmarshal(rawMsg, &segment)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Unmarshall specific segment")
	// 		continue
	// 	}

	// 	aSegmentSlot.segment = segment
	// 	log.Debug().
	// 		Str("slot", fmt.Sprintf("%#v", aSegmentSlot)).
	// 		Str("segment", fmt.Sprintf("%#v", segment)).
	// 		Msg("new")
	// 	slots = append(slots, &aSegmentSlot)
	// }
	// return slots, nil
}

// func (a *App) buildFromJSON(jsonBuf []byte) ([]Widget, error) {
// 	widgets := make([]Widget, 0)
// 	type SegmentSpec struct {
// 		Seg    string          // segment identication
// 		Adjust int             // width adjust for unregular unicode  runes
// 		Args   json.RawMessage // segment specific args
// 	}
// 	specs := []SegmentSpec{}
// 	err := json.Unmarshal(jsonBuf, &specs)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Unmarshal")
// 		return nil, err
// 	}
// 	for _, s := range specs {
// 		jsonBuf, err := s.Args.MarshalJSON()
// 		if err != nil {
// 			panic("IMPOSSIBLE: cannot marshal segment args json")
// 		}
// 		segment, err := a.NewSegmentByNameJSON(s.Seg, jsonBuf)
// 		if err != nil {
// 			log.Error().Err(err).Msg("NewSegmentByNameJSON")
// 			continue
// 		}
// 		widgets = append(widgets, &segmentWidget{
// 			name:    s.Seg,
// 			adjust:  s.Adjust,
// 			segment: segment,
// 		})
// 	}
// 	return widgets, nil
// }

// func (a *App) buildFromJsonSYMBOLVERSION(jsonBuf []byte) ([]*widgetS, error) {
// 	widgets := make([]*widgetS, 0)
// 	type SegmentSpec struct {
// 		Seg  string
// 		Args json.RawMessage
// 	}
// 	segments := []SegmentSpec{}
// 	err := json.Unmarshal(jsonBuf, &segments)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Unmarshal")
// 		return nil, err
// 	}
// 	for _, s := range segments {
// 		pluginFile, err := a.resolvePluginPath(s.Seg)
// 		if err != nil {
// 			log.Error().
// 				Str("Seg", s.Seg).
// 				Err(err).
// 				Msg("resolvePluginPath")
// 			continue
// 		}
// 		pluginLib, err := plugin.Open(pluginFile)
// 		if err != nil {
// 			log.Error().Err(err).Msg("Open")
// 			continue
// 		}
// 		symbol, err := pluginLib.Lookup(SegmentEntrySymbolName)
// 		if err != nil {
// 			log.Error().Err(err).
// 				Str("file", pluginFile).
// 				Str("symbol", SegmentEntrySymbolName).
// 				Msg("Lookup")
// 			continue
// 		}
// 		newFunc := symbol.(func([]byte) Segment)
// 		argsBuf, _ := s.Args.MarshalJSON()
// 		segment := newFunc(argsBuf)
// 		widgets = append(widgets, &widgetS{
// 			segment: segment,
// 		})
// 	}
// 	return widgets, nil
// }

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// )

// type MakeSegFunc func() Segment

// var makeSegFromJson = map[string]MakeSegFunc{}

// var segmentDesc = map[string]string{}

// func SegRegister(name string, desc string, makeFn MakeSegFunc) {
// 	makeSegFromJson[name] = makeFn
// 	segmentDesc[name] = desc
// }

// func buildWidgetsFromJson(msgs []json.RawMessage) []Widget {
// 	widgets := make([]Widget, 0)
// 	for i, msg := range msgs {
// 		spec := PartsConfigSegType{}
// 		if err := json.Unmarshal(msg, &spec); err != nil {
// 			log.Printf("Error: no seg: %s", err)
// 			os.Exit(1)
// 		}
// 		fn, exists := makeSegFromJson[spec.Seg]
// 		if !exists {
// 			log.Printf("Error: seg %s does not exists", spec.Seg)
// 			continue
// 		}
// 		segment := fn()
// 		if err := json.Unmarshal(msg, segment); err != nil {
// 			log.Printf("Error: cannot unmarshal segment %s params: %s", spec.Seg, err)
// 			continue
// 		}
// 		//log.Printf("segment %#v", segment)
// 		// Add spaces between segments
// 		if i != 0 && i < len(msgs) {
// 			widgets = append(widgets, Space())
// 		}
// 		widgets = append(widgets, SegmentWidget(segment))
// 	}
// 	return widgets
// }

// func printPrompt() {
// 	// Build widgets from json
// 	partsConfig := PartsConfig{}
// 	if err := json.Unmarshal([]byte(defaultJson), &partsConfig); err != nil {
// 		log.Printf("Error: json.Unmarshal([]byte(defaultJson), &partsConfig): %s", err)
// 		os.Exit(1)
// 	}
// 	widgetsLeft := buildWidgetsFromJson(partsConfig.Left)
// 	widgetsRight := buildWidgetsFromJson(partsConfig.Right)
// 	widgetsStart := buildWidgetsFromJson(partsConfig.Start)

// 	// Make infoline
// 	widgetsLine := make([]Widget, 0)
// 	widgetsLine = append(widgetsLine, Space()) // margin
// 	widgetsLine = append(widgetsLine, widgetsLeft...)
// 	filler := Filler()
// 	widgetsLine = append(widgetsLine, filler)
// 	widgetsLine = append(widgetsLine, widgetsRight...)
// 	widgetsLine = append(widgetsLine, Space()) // margin

// 	// Make start part
// 	widgetsStart = append(widgetsStart, Space())

// 	// Allocate
// 	empty := allocateSimple(widgetsLine, GetWidth())
// 	filler.Allocate(empty)
// 	allocateSimple(widgetsStart, GetWidth())

// 	// Print line
// 	// TODO: Shell escape $'s in output
// 	// TODO: Shell escape \'s in output
// 	//fmt.Printf("%s\n", strings.Repeat("-", lenTerm))
// 	buf.WriteString("\n")
// 	for _, w := range widgetsStart {
// 		for _, c := range w.Chunks() {
// 			buf.WriteString(fg(c.fg))
// 			buf.WriteString(shellEscape(c.text))
// 		}
// 	}
// 	buf.WriteString(rz())
// }
