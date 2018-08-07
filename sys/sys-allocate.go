package main

import (
	"math"
)

// emptyLen is length of space not used in allocation
func allocateOld(widgets []Widget, maxLen int) (emptyLen int) {
	for _, w := range widgets {
		w.Allocate(maxLen)
	}
	lenWidgets := widgetsLen(widgets)
	if lenWidgets > maxLen {
		frac := float64(maxLen) / float64(lenWidgets)
		for _, w := range widgets {
			w.Allocate(int(math.Floor(frac * float64(w.Len()))))
			if widgetsLen(widgets) <= maxLen {
				break
			}
		}
	}
	return maxLen - widgetsLen(widgets)
}

// Simple allocation of widgets/segments, does not shrink/drop widgets
func allocateSimple(widgets []Widget, maxLen int) (emptyLen int) {
	prevLen := -1
	for _, widget := range widgets {
		widget.Allocate(maxLen)
		//log.Printf("widget.Len: %d", widget.Len())
		switch w := widget.(type) {
		// If there is no previous widget don't print space
		case *spaceWidget:
			if prevLen == 0 {
				w.Allocate(0)
			}
		}
		prevLen = widget.Len()
	}
	return maxLen - widgetsLen(widgets)
}
