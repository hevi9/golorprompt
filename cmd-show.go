package main

import (
	"reflect"
	"fmt"
	"golang.org/x/text/width"
)

func listSigns(signs Signs) {
	typ := reflect.TypeOf(signs)
	val := reflect.ValueOf(signs)
	for i := 0; i < typ.NumField(); i++ {
		s := val.Field(i).String()
		fmt.Printf("  %10s = | M%sM | %+q\n",
			typ.Field(i).Name, s, s)
	}
	fmt.Printf("\n")
}

func kind(kind width.Kind) string {
	switch kind {
	case width.Neutral:
		return "neutral"
	case width.EastAsianAmbiguous:
		return "east asian ambiguous"
	case width.EastAsianWide:
		return "east asian wide"
	case width.EastAsianNarrow:
		return "east asian narrow"
	case width.EastAsianFullwidth:
		return "east asian full width"
	case width.EastAsianHalfwidth:
		return "east asian half width"
	default:
		return "unknown"
	}
}

func listMarks(marks Marks) {
	typ := reflect.TypeOf(marks)
	val := reflect.ValueOf(marks)
	for i := 0; i < typ.NumField(); i++ {
		r := rune(val.Field(i).Int())
		props := width.LookupRune(r)
		fmt.Printf(
			"  %10s = | M%cM | %U [ %c %c %c ] %s\n",
			typ.Field(i).Name, r, r,
			props.Narrow(), props.Wide(), props.Folded(),
			kind(props.Kind()))
	}
	fmt.Printf("\n")
}

func printCurrentTheme() {
	fmt.Printf("Used signs:\n")
	listSigns(sign)
	fmt.Printf("Experimental marks:\n")
	listMarks(mark)
}

