package bloc

import (
	"memory/src/types"
)

func While_D() types.Function {
	return types.Function{
		Name:    "while",
		ReturnT: "nil",
		ArgsCount: 1,
		Args: []types.Arg{
			{
				Name: "condition",
				T:    "bool",
			},
		},
		Execute: While,
	}
}

func While(f *types.Function) any {
	return nil
}