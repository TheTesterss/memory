package registry

var FunctionNames []string = []string{
	"print",
	"setVar",
	"if",
	"end",
	"elseif",
	"else",
}

var VariableNames []string = []string{}

func GetAvailableFunctionsNames() []string {
	return FunctionNames
}

func GetAvailableVariablesNames() []string {
	return VariableNames
}