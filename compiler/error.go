package compiler

import (
	"fmt"

	"github.com/shelmesky/tengo/compiler/ast"
	"github.com/shelmesky/tengo/compiler/source"
)

// Error represents a compiler error.
type Error struct {
	fileSet *source.FileSet
	node    ast.Node
	error   error
}

func (e *Error) Error() string {
	filePos := e.fileSet.Position(e.node.Pos())
	return fmt.Sprintf("Compile Error: %s\n\tat %s", e.error.Error(), filePos)
}
