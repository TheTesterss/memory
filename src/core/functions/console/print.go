package console

import (
	"fmt"
	"memory/src/types"
)

func Print_D() types.Function {
	return types.Function{
		Name:    "$print",
		ReturnT: "nil",
		ArgsCount: -1,
		Args: []types.Arg{
			{
				Name: "content",
				T:    "any",
			},
		},
		Execute: Print,
	}
}

func Print(f *types.Function) any {
	for _, arg := range f.Args {
		fmt.Println(arg.Value);
	}
	return nil
}