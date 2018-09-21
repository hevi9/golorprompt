package sys

import "github.com/rs/zerolog"

type errors struct {
	Count int
}

// Errors Number of errors countered
var Errors = &errors{}

// Run zerolog hook for counting errors
func (er *errors) Run(e *zerolog.Event, l zerolog.Level, msg string) {
	if l >= zerolog.ErrorLevel {
		er.Count++
	}
}
