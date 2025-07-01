package IDENT

import (
	"fmt"
	"memory/src/args"
	"memory/src/types"
	"os"
	"memory/src/maps"
	"memory/src/registry"
)

func InstancyFunction(item types.Item) types.Function {
	var args []types.Arg = args.Split(item)
	var correctedSub types.Function

	if item.SubFunction != nil {
		correctedSub = InstancyFunction(*item.SubFunction)
	}

	var availableFunctions map[string]types.Function = maps.GetAvailableFunctions()

	if !Contains(item.Name, registry.GetAvailableFunctionsNames()) {
		fmt.Printf("[73402] - At line %d - %s is not valid.\n", item.Line, item.Name)
		os.Exit(1)
	}

	return types.Function{
		Name: item.Name,
		Args: args,
		ArgsCount: float32(len(args)),
		SubFunction: &correctedSub,
		Execute: availableFunctions[item.Name].Execute,
	}
}

func Contains(t string, l []string) bool {
	for i := range l {
		if l[i] == t {
			return true
		}
	}
	return false
}