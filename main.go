package main

import (
	"fmt"
	"memory/src/types"
	"memory/src/func"
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
	var functions []types.Function = []types.Function{}
	for _, element := range result {
		functions = append(functions, IDENT.InstancyFunction(element))
	}

	IDENT.ExecuteFunctions(functions, result)
}
