package parser_test

import (
	"testing"

	"github.com/shelmesky/tengo/compiler/ast"
	"github.com/shelmesky/tengo/compiler/token"
)

func TestBoolean(t *testing.T) {
	expect(t, "true", func(p pfn) []ast.Stmt {
		return stmts(
			exprStmt(
				boolLit(true, p(1, 1))))
	})

	expect(t, "false", func(p pfn) []ast.Stmt {
		return stmts(
			exprStmt(
				boolLit(false, p(1, 1))))
	})

	expect(t, "true != false", func(p pfn) []ast.Stmt {
		return stmts(
			exprStmt(
				binaryExpr(
					boolLit(true, p(1, 1)),
					boolLit(false, p(1, 9)),
					token.NotEqual,
					p(1, 6))))
	})

	expect(t, "!false", func(p pfn) []ast.Stmt {
		return stmts(
			exprStmt(
				unaryExpr(
					boolLit(false, p(1, 2)),
					token.Not,
					p(1, 1))))
	})
}
