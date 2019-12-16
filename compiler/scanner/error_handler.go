package scanner

import "github.com/shelmesky/tengo/compiler/source"

// ErrorHandler is an error handler for the scanner.
type ErrorHandler func(pos source.FilePos, msg string)
