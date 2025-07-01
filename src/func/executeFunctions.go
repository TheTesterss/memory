package IDENT

import (
	"fmt"
	"memory/src/types"
	"memory/src/maps"
	"os"
	"memory/src/registry"
)

var functions_D map[string]types.Function = maps.GetAvailableFunctions()
var _ []string = registry.GetAvailableFunctionsNames()

func ExecuteFunctions(items []types.Item) {
	for i, element := range items {
		if element.Name != "" {
			ExecuteFunction(i, InstancyFunction(element), items)
		}
	}
}

func ExecuteFunction(i int, function types.Function, items []types.Item) {
	if function.ArgsCount < functions_D[function.Name].ArgsCount && functions_D[function.Name].ArgsCount != -1 {
			fmt.Printf("[73402] - At line %d - %s doesn't count as much arguments as required (demanded = %.0f/gave = %.0f).\n", items[i].Line, function.Name, function.ArgsCount, functions_D[function.Name].ArgsCount)
			os.Exit(1)
	}

	if functions_D[function.Name].ArgsCount != -1 {
		for j, arg := range function.Args {
			if arg.T != functions_D[function.Name].Args[j].T && functions_D[function.Name].Args[j].T != "any" {
				fmt.Printf("[73402] - At line %d - %s can't match different types (demanded = %s/gave = %s).\n", items[i].Line, function.Name, arg.T, functions_D[function.Name].Args[j].T)
				os.Exit(1)
			}
		}
	} else if functions_D[function.Name].ArgsCount > 0 {
		for _, arg := range function.Args {
			if arg.T != functions_D[function.Name].Args[0].T && functions_D[function.Name].Args[0].T != "any" {
				fmt.Printf("[73402] - At line %d - %s can't match different types (demanded = %s/gave = %s).\n", items[i].Line, function.Name, arg.T, functions_D[function.Name].Args[0].T)
				os.Exit(1)
			}
		}
	}

	functions_D[function.Name].Execute(&function)
}