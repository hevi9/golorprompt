package sys

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

var jsonText = `
[
	{ "seg": "cwd" },
	{ "seg": "stub" },
	{ "seg": "hostname", "args": {"showifenv": "HOME"}}
]
`

func CommandPrompt(app *App) error {
	// build widges from json spec
	widgets, err := app.buildFromJson([]byte(jsonText))
	if err != nil {
		log.Error().Err(err).Msg("Cannot build from json spec")
		app.AddError(err)
		return err
	}

	// render widgets concurrently
	wg := sync.WaitGroup{}
	for _, widgetElem := range widgets {
		wg.Add(1)
		go func(widgetElem *widgetS) {
			defer wg.Done()
			widgetElem.chunks = widgetElem.segment.Render(app)
		}(widgetElem)
	}
	wg.Wait()

	// make layout

	// print widgets
	buf := bytes.Buffer{}
	buf.WriteString(Bg(Config.BgLine))
	for _, widgetElem := range widgets {
		// fmt.Printf("%#v\n", widgetElem)
		for _, chunk := range widgetElem.chunks {
			buf.WriteString(Fg(chunk.Fg))
			buf.WriteString(shellEscapeZsh(chunk.Text))
		}
	}
	buf.WriteString(Rz())
	fmt.Printf("%s", buf.String())

	return nil
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
// 	buf.WriteString("\n")
// 	for _, w := range widgetsStart {
// 		for _, c := range w.Chunks() {
// 			buf.WriteString(fg(c.fg))
// 			buf.WriteString(shellEscape(c.text))
// 		}
// 	}
// 	buf.WriteString(rz())
// }
