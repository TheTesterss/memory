package resolvers

import (
	"memory/src/maps"
	"memory/src/registry"
	"memory/src/util"
	"memory/src/vars"
	"strings"
)

func ResolveValue(s string) (string, string) {
	s = strings.TrimSpace(s)

	if util.Contains(s, registry.GetAvailableFunctionsNames()) {
		function_D := maps.GetAvailableFunctions()[s]
		result := function_D.Execute(&function_D)
		if val, isString := result.(string); isString {
			return val, function_D.ReturnT
		} else {
			return "", "any"
		}
	} else if util.Contains(s, registry.GetAvailableVariablesNames()) {
		variable := vars.GetAvailableVariables()[s]
		return variable.Value, variable.T
	}

	return s, ""
}