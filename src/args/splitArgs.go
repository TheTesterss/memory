package args

import (
	"fmt"
	"memory/src/core/conditions"
	"memory/src/types"
	"memory/src/util"
	"os"
)

func Split(f types.Item) []types.Arg {
	if f.Listed_args == "" {
		return []types.Arg{}
	}
	var r []types.Arg = []types.Arg{
		{T: "", Value: ""},
	}
	var inString bool = false
	var canOpenString bool = true

	for i, char := range f.Listed_args {

		switch string(char) {

		case "\"":
			arg := &r[len(r)-1]
			if i > 0 && string(f.Listed_args[i-1]) == "\\" {
				arg.Value+="\""
				continue
			}
			if !canOpenString && !inString {
				fmt.Printf("[73402] - At line %d - Couldn't open a string twice in the same argument.\n", f.Line)
				os.Exit(1)
			}
			inString = !inString
			if !inString && len(f.Listed_args)-1 > i && string(f.Listed_args[i+1]) != ";" {
				fmt.Printf("[73402] - At line %d - Closing a string but not the argument.\n", f.Line)
				os.Exit(1)
			}
			canOpenString = false
			if inString {
				arg.T = "str"
			}

		case ";":
			if inString {
				arg := &r[len(r)-1]
				arg.Value+=";"
			}
			v := &r[len(r)-1]
			v.Value = conditions.ResolveValue(v.Value)
			if util.IsNumber(v.Value) {
				if v.T == "str" { // Can't use two types for a single argument.
					fmt.Printf("[73402] - At line %d - %s can't be both string and int.\n", f.Line, v.Value)
					os.Exit(1)
				}
				v.T = "int" // The value is an integer or a float.
			} else if v.Value == "nil" {
				if v.T == "str" { // Can't use two types for a single argument.
					fmt.Printf("[73402] - At line %d - %s can't be both string and nil.\n", f.Line, v.Value)
					os.Exit(1)
				}
				v.T = "nil" // The value is nil without "" that means the nil type is just simply called.
			} else if util.IsBoolean(v.Value) {
				if v.T == "str" { // Can't use two types for a single argument.
					fmt.Printf("[73402] - At line %d - %s can't be both string and bool.\n", f.Line, v.Value)
					os.Exit(1)
				}
				v.T = "bool" // The value is true/false.
			} else if conditions.IsCondition(v.Value) {
				result := conditions.EvaluateConditions(v.Value)
				v.Value = fmt.Sprintf("%v", result) // Changes the value by true/false
				v.T = "bool"
			}
			if v.T == "str" && inString { // Never closed the string.
				fmt.Printf("[73402] - At line %d - %s is an opened string but never closed.\n", f.Line, v.Value)
				os.Exit(1)
			}
			if v.T == "" { // No accorded type which means its not valid.
				fmt.Printf("[73402] - At line %d - %s is not accorded to any working type (int/nil/bool/str).\n", f.Line, v.Value)
				os.Exit(1)
			}

			r = append(r, types.Arg{})
			canOpenString = true
			inString = false

		default:
			arg := &r[len(r)-1]
			arg.Value+=string(char)
		}
	}
	v := &r[len(r)-1]
	v.Value = conditions.ResolveValue(v.Value)
	if util.IsNumber(v.Value) {
		if v.T == "str" { // Can't use two types for a single argument.
			fmt.Printf("[73402] - At line %d - %s can't be both string and int.\n", f.Line, v.Value)
			os.Exit(1)
		}
		v.T = "int" // The value is an integer or a float.
	} else if v.Value == "nil" {
		if v.T == "str" { // Can't use two types for a single argument.
			fmt.Printf("[73402] - At line %d - %s can't be both string and nil.\n", f.Line, v.Value)
			os.Exit(1)
		}
		v.T = "nil" // The value is nil without "" that means the nil type is just simply called.
	} else if util.IsBoolean(v.Value) {
		if v.T == "str" { // Can't use two types for a single argument.
			fmt.Printf("[73402] - At line %d - %s can't be both string and bool.\n", f.Line, v.Value)
			os.Exit(1)
		}
		v.T = "bool" // The value is true/false.
	} else if conditions.IsCondition(v.Value) {
		result := conditions.EvaluateConditions(v.Value)
		v.Value = fmt.Sprintf("%v", result) // Changes the value by true/false
		v.T = "bool"
	}
	if v.T == "str" && inString {
		fmt.Printf("[73402] - At line %d - %s is an opened string but never closed.\n", f.Line, v.Value)
		os.Exit(1)
	}
	if v.T == "" {
		fmt.Printf("[73402] - At line %d - %s is not accorded to any working type (int/nil/bool/str).\n", f.Line, v.Value)
		os.Exit(1)
	}
	return r
}