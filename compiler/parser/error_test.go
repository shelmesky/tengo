package parser_test

import (
	"testing"

	"github.com/shelmesky/tengo/assert"
	"github.com/shelmesky/tengo/compiler/parser"
	"github.com/shelmesky/tengo/compiler/source"
)

func TestError_Error(t *testing.T) {
	err := &parser.Error{Pos: source.FilePos{Offset: 10, Line: 1, Column: 10}, Msg: "test"}
	assert.Equal(t, "Parse Error: test\n\tat 1:10", err.Error())
}
