package bloc

import (
	"memory/src/types"
)

func If_D() types.Function {
	return types.Function{
		Name:    "if",
		ReturnT: "nil",
		ArgsCount: 1,
		Args: []types.Arg{
			{
				Name: "condition",
				T:    "bool",
			},
		},
		Execute: If,
	}
}

func If(f *types.Function) any {
	return nil
}

func End_D() types.Function {
	return types.Function{
		Name:    "end",
		ReturnT: "nil",
		ArgsCount: 1,
		Args: []types.Arg{},
		Execute: End,
	}
}

func End(f *types.Function) any {
	return nil
}

func Elseif_D() types.Function {
	return types.Function{
		Name:    "elseif",
		ReturnT: "nil",
		ArgsCount: 0,
		Args: []types.Arg{
			{
				Name: "condition",
				T:    "bool",
			},
		},
		Execute: Elseif,
	}
}

func Elseif(f *types.Function) any {
	return nil
}

func Else_D() types.Function {
	return types.Function{
		Name:    "else",
		ReturnT: "nil",
		ArgsCount: 0,
		Args: []types.Arg{},
		Execute: Else,
	}
}

func Else(f *types.Function) any {
	return nil
}
