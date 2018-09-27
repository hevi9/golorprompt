package sys

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

// CommandPrompt Build prompt from given json spec
func CommandPrompt(jsonBuf []byte) error {
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
			s2.Render()
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

func resolvePluginPath(name string) (string, error) {
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
func NewSegmentByNameJSON(name string, jsonBuf []byte) (Segment, error) {
	info, ok := segmentRegistry[name]
	if !ok {
		return nil, fmt.Errorf("name '%s' does not exists", name)
	}
	segment := info.newSegmentFunc()
	// log.Info().Str("segment", info.name).Msg("new segment")
	return segment, nil
}

func buildFromJSON(jsonBuf []byte) (Slot, error) {

	type JsonSlot struct {
		Segment  string
		Segments []json.RawMessage
	}

	var recurse func(rmsg json.RawMessage) (Slot, error)
	recurse = func(rmsg json.RawMessage) (Slot, error) {
		jslot := &JsonSlot{}
		err := json.Unmarshal(rmsg, jslot)
		if err != nil {
			log.Error().Err(err).Msg("Unmarshal recursive")
			return nil, err
		}
		slot := segmentSlot{}
		// log.Printf("%#v", slot)
		for _, subRawMsg := range jslot.Segments {
			recurse(subRawMsg)
		}
		return slot, nil
	}
	root := json.RawMessage{}
	err := json.Unmarshal(jsonBuf, &root)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal root")
		return nil, err
	}
<<<<<<< HEAD
	log.Printf("%#v", root)
=======
	recurse(root)
>>>>>>> 42b5acd6fec1250a083168f69dfc386e4468de2d

	// ..
	return &segmentSlot{}, nil

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

// Return lines of slots
func makeLayout(slots []Slot) [][]Slot {

	// remove empty slots
	tmp := make([]Slot, 0)
	for _, s := range slots {
		if s.Chunks() != nil {
			tmp = append(tmp, s)
		} else {
			log.Debug().Str("name", s.Name()).Msg("discard empty segment")
		}
	}
	slots = tmp

	// compress space
	tmp = make([]Slot, 0)
	prevName := ""
	for _, s := range slots {
		if s.Name() == "space" {
			if prevName != "space" {
				tmp = append(tmp, s)
			}
		} else {
			tmp = append(tmp, s)
		}
		prevName = s.Name()
	}
	slots = tmp

	// split widgets to lines
	lines := make([][]Slot, 0)
	line := make([]Slot, 0)
	hasWidgets := len(slots)
	currentWidgetIdx := 0
	currentLineLen := 0
	for hasWidgets > 0 {
		s := slots[currentWidgetIdx]
		currentLineLen += s.Len()
		// or newline segment
		if currentLineLen < GetWidth() && s.Name() != "newline" {
			line = append(line, s)
		} else {
			lines = append(lines, line)
			line = make([]Slot, 0)
			line = append(line, s)
			currentLineLen = 0
		}
		currentWidgetIdx++
		hasWidgets--
	}
	lines = append(lines, line)

	// fill lines
	for i := 0; i < len(lines)-1; i++ {
		fillCnt := GetWidth() - slotsLen(lines[i])
		lines[i] = append(lines[i], &segmentSlot{
			chunks: []Chunk{
				Chunk{
					Text: strings.Repeat("^", maxInt(fillCnt, 1)),
				},
			},
		})
		log.Debug().Int("line len", slotsLen(line)).Msg("")
	}

	// log.Debug().
	// 	Int("widgets width", widgetsLen(widgets)).
	// 	Int("terminal width", GetWidth()).
	// 	Msg("widths")

	// debug
	// for _, w := range widgets {
	// 	log.Debug().
	// 		Str("name", w.Name()).
	// 		Msg("apply segment")
	// }

	return lines
}

// import (
// 	"math"
// )

// // emptyLen is length of space not used in allocation
// func allocateOld(widgets []Widget, maxLen int) (emptyLen int) {
// 	for _, w := range widgets {
// 		w.Allocate(maxLen)
// 	}
// 	lenWidgets := widgetsLen(widgets)
// 	if lenWidgets > maxLen {
// 		frac := float64(maxLen) / float64(lenWidgets)
// 		for _, w := range widgets {
// 			w.Allocate(int(math.Floor(frac * float64(w.Len()))))
// 			if widgetsLen(widgets) <= maxLen {
// 				break
// 			}
// 		}
// 	}
// 	return maxLen - widgetsLen(widgets)
// }

// // Simple allocation of widgets/segments, does not shrink/drop widgets
// func allocateSimple(widgets []Widget, maxLen int) (emptyLen int) {
// 	prevLen := -1
// 	for _, widget := range widgets {
// 		widget.Allocate(maxLen)
// 		//log.Printf("widget.Len: %d", widget.Len())
// 		switch w := widget.(type) {
// 		// If there is no previous widget don't print space
// 		case *spaceWidget:
// 			if prevLen == 0 {
// 				w.Allocate(0)
// 			}
// 		}
// 		prevLen = widget.Len()
// 	}
// 	return maxLen - widgetsLen(widgets)
// }

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
