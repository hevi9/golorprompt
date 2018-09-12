package sys

import (
	"strings"

	"github.com/rs/zerolog/log"
)

// Return lines of slots
func makeLayout(slots []Slot) [][]Slot {

	// remove empty slots
	tmp := make([]Slot, 0)
	for _, s := range slots {
		if s.Chunks() != nil {
			tmp = append(tmp, s)
		} else {
			log.Debug().Str("name", s.Name()).Msg("discard")
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
