package console

import (
	"fmt"
	"memory/src/core"
)

func Print_D() core.Function {
	return core.Function{
		Name:    "$print",
		ReturnT: "nil",
		ArgsCount: -1,
		Args: []core.Arg{
			{
				Name: "content",
				T:    "any",
			},
		},
		Execute: Print,
	}
}

func Print(f *core.Function) any {
	for _, arg := range f.Args {
		fmt.Println(arg)
		fmt.Println(arg.Value);
	}
	return nil;
}