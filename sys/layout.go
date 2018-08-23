package sys

import (
	"strings"

	"github.com/rs/zerolog/log"
)

// Return lines of widgets
func makeLayout(widgets []Widget) [][]Widget {

	// remove nil widgets
	tmp := make([]Widget, 0)
	for _, w := range widgets {
		if w.Chunks() != nil {
			tmp = append(tmp, w)
		} else {
			log.Debug().Str("name", w.Name()).Msg("discard")
		}
	}
	widgets = tmp

	// compress space
	tmp = make([]Widget, 0)
	prevName := ""
	for _, w := range widgets {
		if w.Name() == "space" {
			if prevName != "space" {
				tmp = append(tmp, w)
			}
		} else {
			tmp = append(tmp, w)
		}
		prevName = w.Name()
	}
	widgets = tmp

	// split widgets to lines
	lines := make([][]Widget, 0)
	line := make([]Widget, 0)
	hasWidgets := len(widgets)
	currentWidgetIdx := 0
	currentLineLen := 0
	for hasWidgets > 0 {
		w := widgets[currentWidgetIdx]
		currentLineLen += w.Len()
		// or newline segment
		if currentLineLen < GetWidth() {
			line = append(line, w)
		} else {
			lines = append(lines, line)
			line = make([]Widget, 0)
			line = append(line, w)
			currentLineLen = 0
		}
		currentWidgetIdx++
		hasWidgets--
	}
	lines = append(lines, line)

	// fill lines
	for i := range lines {
		fillCnt := GetWidth() - widgetsLen(lines[i])
		lines[i] = append(lines[i], &segmentWidget{
			chunks: []Chunk{
				Chunk{
					Text: strings.Repeat("^", maxInt(fillCnt, 1)),
				},
			},
		})

		log.Debug().Int("line len", widgetsLen(line)).Msg("")
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
