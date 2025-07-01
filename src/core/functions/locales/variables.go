package locales

import (
	"fmt"
	"memory/src/types"
	"memory/src/vars"
	"os"
	"strings"
	"memory/src/registry"
	"memory/src/util"
)

func SetVar_D() types.Function {
	return types.Function{
		Name:    "setVar",
		ReturnT: "nil",
		ArgsCount: 2,
		Args: []types.Arg{
			{
				Name: "name",
				T:    "str",
			},
			{
				Name: "value",
				T:    "any",
			},
			{
				Name: "type",
				T:    "any",
			},
		},
		Execute: SetVar,
	}
}

func SetVar(f *types.Function) any {
	name := f.Args[0]
	value := f.Args[1]
	t := f.Args[2]
	if util.Contains(name.Value, registry.GetAvailableFunctionsNames()) {
		fmt.Printf("[73402] - A function already exist with this name: %s.\n", name.Value)
		os.Exit(1)
	}
	if !util.Contains(name.Value, registry.GetAvailableVariablesNames()) {
		if(!util.Contains(t.Value, []string{"any", "bool", "nil", "int", "str"})) {
			fmt.Printf("[73402] - %s, the types is not correct: %s.\n", name.Value, t.Value)
			os.Exit(1)
		}
		if t.Value == "" {
			t.Value = "any"
		}
		vars.SetAvailableVariables(&types.Variable{ Name: name.Value, T: t.Value, Value: value.Value }, false)
	} else { // Already declared.
		t.Value = vars.GetAvailableVariables()[name.Value].T
		if (t.Value == "bool" && !util.IsBoolean(value.Value)) ||
			(t.Value == "int" && !util.IsNumber(value.Value)) ||
			(t.Value == "str" && (!strings.HasPrefix(value.Value, "\"") || !strings.HasSuffix(value.Value, "\"") || strings.Count(value.Value, "\"") > 2)) ||
			(t.Value == "nil" && value.Value != "nil") ||
			(!util.Contains(t.Value, []string{"any", "bool", "nil", "int", "str"})) {
				fmt.Printf("[73402] - %s, the types are not matching: %s, %s.\n", name.Value, value.Value, t.Value)
				os.Exit(1)
		}

		vars.SetAvailableVariables(&types.Variable{ Name: name.Value, T: t.Value, Value: value.Value }, true)
	}
	return nil
}

