package vars

import (
	"memory/src/registry"
	"memory/src/types"
)

var variables map[string]types.Variable = map[string]types.Variable{}

func SetAvailableVariables(variable *types.Variable, redeclared bool) {
	if !redeclared {
		registry.VariableNames = append(registry.GetAvailableVariablesNames(), variable.Name)
	}
	variables[variable.Name] = *variable
}

func GetAvailableVariables() map[string]types.Variable {
	return variables
}