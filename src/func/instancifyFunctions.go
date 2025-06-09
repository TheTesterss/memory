package IDENT

import (
	"fmt"
	"memory/src/args"
	"memory/src/core"
	"os"
)

func InstancyFunction(item core.Item) core.Function {
	var args []core.Arg = args.Split(item)
	var correctedSub core.Function

	if item.SubFunction != nil {
		correctedSub = InstancyFunction(*item.SubFunction)
	}

	var availableFunctions map[string]core.Function = GetAvailableFunctions()

	if !Contains(item.Name, GetAvailableFunctionsNames()) {
		fmt.Printf("[73402] - At line %d - %s is not valid.\n", item.Line, item.Name)
		os.Exit(1)
	}

	return core.Function{
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