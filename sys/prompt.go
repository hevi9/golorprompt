// TODO: IDEA @ as a alone hostname hash colored

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

type appS struct {
	pluginPaths []string
	errors      int
}

func (a *appS) resolvePluginPath(name string) (string, error) {
	suffixes := []string{".so", ""}
	for _, p := range a.pluginPaths {
		for _, s := range suffixes {
			fullpath := path.Join(p, name) + s
			if _, err := os.Stat(fullpath); err == nil {
				return fullpath, nil
			}
		}
	}
	return "", fmt.Errorf("%s not found", name)
}

func (a *appS) buildFromJson(jsonBuf []byte) ([]*widgetS, error) {
	widgets := make([]*widgetS, 0)
	type SegmentSpec struct {
		Seg  string
		Args json.RawMessage
	}
	segments := []SegmentSpec{}
	err := json.Unmarshal(jsonBuf, &segments)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal")
		a.errors++
		return nil, err
	}
	for _, s := range segments {
		pluginFile, err := a.resolvePluginPath(s.Seg)
		if err != nil {
			log.Error().Err(err).Msg("")
			a.errors++
			continue
		}
		pluginLib, err := plugin.Open(pluginFile)
		if err != nil {
			log.Error().Err(err).Msg("")
			a.errors++
			continue
		}
		symbol, err := pluginLib.Lookup(SegmentEntrySymbolName)
		if err != nil {
			log.Error().Err(err).
				Str("file", pluginFile).
				Str("symbol", SegmentEntrySymbolName).
				Msg("")
			a.errors++
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

var jsonText = `
[
	{
		"seg": "cwd"
	}
]
`

func CommandPrompt(app *appS) {
	// widgets, err := app.buildFromJson([]byte(jsonText))
}

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
// 	buf := bytes.Buffer{}
// 	buf.WriteString(bg(config.BgLine))
// 	for _, w := range widgetsLine {
// 		for _, c := range w.Chunks() {
// 			buf.WriteString(fg(c.fg))
// 			buf.WriteString(shellEscape(c.text))
// 		}
// 	}
// 	buf.WriteString(rz())
// 	buf.WriteString("\n")
// 	for _, w := range widgetsStart {
// 		for _, c := range w.Chunks() {
// 			buf.WriteString(fg(c.fg))
// 			buf.WriteString(shellEscape(c.text))
// 		}
// 	}
// 	buf.WriteString(rz())
// 	fmt.Printf("%s", buf.String())
// }
