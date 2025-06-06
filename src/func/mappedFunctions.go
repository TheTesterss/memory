package IDENT

import (
	"memory/src/core"
	"memory/src/core/functions/console"
)

func GetAvailableFunctions() map[string]core.Function {
	return map[string]core.Function{
		"$print": console.Print_D(),
	}
}