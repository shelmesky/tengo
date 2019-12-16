package objects_test

import (
	"testing"

	"github.com/shelmesky/tengo/assert"
	"github.com/shelmesky/tengo/compiler/token"
	"github.com/shelmesky/tengo/objects"
)

func testBinaryOp(t *testing.T, lhs objects.Object, op token.Token, rhs objects.Object, expected objects.Object) bool {
	t.Helper()

	actual, err := lhs.BinaryOp(op, rhs)

	return assert.NoError(t, err) && assert.Equal(t, expected, actual)
}

func boolValue(b bool) objects.Object {
	if b {
		return objects.TrueValue
	}

	return objects.FalseValue
}
