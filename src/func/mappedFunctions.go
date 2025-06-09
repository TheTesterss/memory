package IDENT

import (
	"memory/src/core"
	"memory/src/core/functions/console"
)

var names []string = []string{
	"$print",
}
var functions map[string]core.Function = map[string]core.Function{
	"$print": console.Print_D(),
}

func GetAvailableFunctions() map[string]core.Function {
	return functions
}

func GetAvailableFunctionsNames() []string {
	return names
}