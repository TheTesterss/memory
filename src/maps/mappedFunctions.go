package maps

import (
	"memory/src/core/functions/bloc"
	"memory/src/core/functions/console"
	"memory/src/core/functions/locales"
	"memory/src/types"
)

var functions map[string]types.Function = map[string]types.Function{
	"print": console.Print_D(),
	"setVar": locales.SetVar_D(),
	"if": bloc.If_D(),
	"end": bloc.End_D(),
	"else": bloc.Else_D(),
	"elseif": bloc.Elseif_D(),
}

func GetAvailableFunctions() map[string]types.Function {
	return functions
}