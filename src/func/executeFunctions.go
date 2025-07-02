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

func ExecuteFunctions(items []types.Item) {
    i := 0
    for i < len(items) {
        if items[i].Name != "" {
            instanciedElement := InstancyFunction(items[i])
            switch instanciedElement.Name {

            case "if":
                blockEnd := findMatchingEnd(items, i)
                branchToExec := -1
				
                j := i
                for j < blockEnd {
                    if items[j].Name == "if" || items[j].Name == "elseif" {
                        inst := InstancyFunction(items[j])
                        if !conditions.IsCondition(inst.Args[0].Value) {
                            fmt.Printf("[73402] - At line %d - Can't open a condition block if the condition is not valid.\n", items[j].Line)
                            os.Exit(1)
                        }
                        if conditions.EvaluateConditions(inst.Args[0].Value) {
                            branchToExec = j
                            break
                        }
                    } else if items[j].Name == "else" {
                        branchToExec = j
                        break
                    }
                    j++
                }
                if branchToExec != -1 {
    				k := branchToExec + 1
    				for k < blockEnd && items[k].Name != "elseif" && items[k].Name != "else" && items[k].Name != "end" {
        				inst := InstancyFunction(items[k])
        				switch items[k].Name {

        				case "if", "while":
        				    subBlockEnd := findMatchingEnd(items, k)
        				    ExecuteFunctions(items[k : subBlockEnd+1])
        				    k = subBlockEnd
        				default:
        				    ExecuteFunction(k, inst, items)
        				}
        				k++
    				}
				}
                i = blockEnd

            case "while":
                blockStart := i
                blockEnd := findMatchingEnd(items, i)
                for {
                    condItem := InstancyFunction(items[blockStart])
                    if !conditions.IsCondition(condItem.Args[0].Value) {
                        fmt.Printf("[73402] - At line %d - Can't open a while block if the condition is not valid.\n", items[blockStart].Line)
                        os.Exit(1)
                    }
                    if !conditions.EvaluateConditions(condItem.Args[0].Value) {
                        break
                    }
                    k := blockStart + 1
					for k < blockEnd && items[k].Name != "end" {
    					inst := InstancyFunction(items[k])
    					switch items[k].Name {
    					
						case "if", "while":
        					subBlockEnd := findMatchingEnd(items, k)
        					ExecuteFunctions(items[k : subBlockEnd+1])
        					k = subBlockEnd
    					default:
        					ExecuteFunction(k, inst, items)
    					}
    					k++
					}
                }
                i = blockEnd
            default:
                ExecuteFunction(i, instanciedElement, items)
            }
        }
        i++
    }
}

func ExecuteFunction(i int, function types.Function, items []types.Item) {
	required := int(functions_D[function.Name].ArgsCount)
	given := int(function.ArgsCount)
	argsDef := functions_D[function.Name].Args

	if required > 0 && given < required {
			fmt.Printf("[73402] - At line %d - %s doesn't count as much arguments as required (demanded = %.0f/gave = %.0f).\n", items[i].Line, function.Name, function.ArgsCount, functions_D[function.Name].ArgsCount)
			os.Exit(1)
	}

	for j := 0; j < required && j < given && j < len(argsDef); j++ {
    	expectedType := argsDef[j].T
    	givenType := function.Args[j].T
    	if expectedType != "any" && expectedType != givenType {
        	fmt.Printf("[73402] - At line %d - %s can't match different types (demanded = %s/gave = %s).\n", items[i].Line, function.Name, expectedType, givenType)
        	os.Exit(1)
    	}
	}
	
	if required == -1 && len(argsDef) > 0 {
        expectedType := argsDef[0].T
        for j := 0; j < given; j++ {
            givenType := function.Args[j].T
            if expectedType != "any" && expectedType != givenType {
                fmt.Printf("[73402] - At line %d - %s can't match different types (demanded = %s/gave = %s).\n", items[i].Line, function.Name, expectedType, givenType)
                os.Exit(1)
            }
        }
    }

	functions_D[function.Name].Execute(&function)
}

func findMatchingEnd(items []types.Item, pos int) int {
    depth := 0
    for i := pos; i < len(items); i++ {
        if items[i].Name == "if" || items[i].Name == "while" {
            depth++
        } else if items[i].Name == "end" {
            depth--
            if depth == 0 {
                return i
            }
        }
    }
    fmt.Printf("[73402] - No matching end found for block at %d\n", pos)
    os.Exit(1)
    return -1
}