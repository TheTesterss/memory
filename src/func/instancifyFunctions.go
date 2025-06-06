package IDENT

import (
	"memory/src/args"
	"memory/src/core"
)

func instancyFunction(item core.Item) core.Function {
	var args []core.Arg = args.Split(item)
	var correctedSub core.Function

	if item.SubFunction != nil {
		correctedSub = instancyFunction(*item.SubFunction)
	}

	var availableFunctions map[string]core.Function = GetAvailableFunctions()

	return core.Function{
		Name: item.Name,
		Args: args,
		ArgsCount: float32(len(args)),
		SubFunction: &correctedSub,
		Execute: availableFunctions[item.Name].Execute,
	}
}