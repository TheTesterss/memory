package resolvers

import (
	"memory/src/vars"
	"strings"
)

func ReplaceVariablesInExpr(expr string) string {
    variables := vars.GetAvailableVariables()
    var names []string
    for name := range variables {
        names = append(names, name)
    }
    for i := 0; i < len(names); i++ {
        for j := i + 1; j < len(names); j++ {
            if len(names[j]) > len(names[i]) {
                names[i], names[j] = names[j], names[i]
            }
        }
    }
    for _, name := range names {
        val := variables[name].Value
        expr = strings.ReplaceAll(expr, name, val)
    }
    return expr
}