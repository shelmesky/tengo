package stdlib

import (
	"encoding/hex"
	"github.com/shelmesky/tengo/objects"
)

var hexModule = map[string]objects.Object{
	"encode": &objects.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &objects.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
