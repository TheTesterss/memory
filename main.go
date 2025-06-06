package main

import (
	"fmt"
	"memory/src/func"
	"memory/src/args"
	"memory/src/core"
	"os"
)

func main() {
	content, err := os.ReadFile("test.tev");
	if err != nil {
		fmt.Println("error while opening file");
		os.Exit(1);
	}

	result := core.Tokenise(string(content));
	var functions []core.Function = []core.Function{}
	for _, element := range result {
		functions = append(functions, instancyFunction(element))
	}

	fmt.Println(functions)

}

func instancyFunction(item core.Item) core.Function {
	var args []core.Arg = args.Split(item)
	var correctedSub core.Function

	if item.SubFunction != nil {
		correctedSub = instancyFunction(*item.SubFunction)
	}

	var availableFunctions map[string]core.Function = IDENT.GetAvailableFunctions()

	return core.Function{
		Name: item.Name,
		Args: args,
		ArgsCount: float32(len(args)),
		SubFunction: &correctedSub,
		Execute: availableFunctions[item.Name].Execute,
	}
}