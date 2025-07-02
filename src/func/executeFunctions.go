package IDENT

import (
	"fmt"
	"memory/src/core/conditions"
	"memory/src/maps"
	"memory/src/registry"
	"memory/src/types"
	"os"
)

var functions_D map[string]types.Function = maps.GetAvailableFunctions()
var _ []string = registry.GetAvailableFunctionsNames()
var depth []string = []string{}

func ExecuteFunctions(items []types.Item) {
	var i int = 0
	for i < len(items) {
		if items[i].Name != "" {
			instanciedElement := InstancyFunction(items[i])
			switch instanciedElement.Name {
			
			case "if":
				depth = append(depth, "if")
				if !conditions.IsCondition(instanciedElement.Args[0].Value) {
					fmt.Printf("[73402] - At line %d - Can't open an if block if the condition is not valid.\n", items[i].Line)
					os.Exit(1)
				}
				if(conditions.EvaluateConditions(instanciedElement.Args[0].Value)) {
					i++
                    for i < len(items) && !(items[i].Name == "elseif" || items[i].Name == "else" || items[i].Name == "end") {
                        ExecuteFunction(i, InstancyFunction(items[i]), items)
                        i++
                    }
					for i < len(items) && items[i].Name != "end" {
                        i++
                    }
				} else {
					found := false
					for i < len(items) && items[i].Name != "end" {
                        if items[i].Name == "elseif" {
							instanciedElement2 := InstancyFunction(items[i])
                            if !conditions.IsCondition(instanciedElement2.Args[0].Value) {
                                fmt.Printf("[73402] - At line %d - Can't open an elseif block if the condition is not valid.\n", items[i].Line)
                                os.Exit(1)
                            }
                            if conditions.EvaluateConditions(instanciedElement2.Args[0].Value) {
    							found = true
    							i++
    							for i < len(items) && !(items[i].Name == "elseif" || items[i].Name == "else" || items[i].Name == "end") {
        							ExecuteFunction(i, InstancyFunction(items[i]), items)
        							i++
    							}
								for i < len(items) && items[i].Name != "end" {
        							i++
    							}
    							break
							}
                        } else if items[i].Name == "else" {
                            found = true
                            i++
                            for i < len(items) && items[i].Name != "end" {
                                ExecuteFunction(i, InstancyFunction(items[i]), items)
                                i++
                            }
                            break
                        }
                        i++
                    }
                    if !found {
                        for i < len(items) && items[i].Name != "end" {
                            i++
                        }
                    }
				}

				if len(depth) > 0 {
                    depth = depth[:len(depth)-1]
                }
			case "elseif", "else":
				for i < len(items) && items[i].Name != "end" {
                    i++
                }
			case "while":
				depth = append(depth, "while")
			case "end":
				if len(depth) > 0 {
                    depth = depth[:len(depth)-1]
                } else {
                    fmt.Printf("[73402] - At line %d - Can't end block because no block opened found.\n", items[i].Line)
                    os.Exit(1)
                }
			case "func":
				depth = append(depth, "func")
			default:
				ExecuteFunction(i, instanciedElement, items)
			}
		}
		i++
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