package args

import (
	"fmt"
	"memory/src/core"
	"memory/src/util"
	"os"
)

func Split(f core.Item) []core.Arg {
	var r []core.Arg = []core.Arg{
		{},
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
			// Handles integers & floats.
			if util.IsNumber(v.Value) {
				v.T = "int"
			}
			// Handles conditions + true & false.
			if util.IsBoolean(v.Value) {
				v.T = "bool"
			}
			if v.Value == "nil" {
				v.T = "nil"
			}
			if v.T == "" {
				v.T = "any"
			}
			r = append(r, core.Arg{})
			canOpenString = true

		default:
			arg := &r[len(r)-1]
			arg.Value+=string(char)
		}
	}
	return r
}