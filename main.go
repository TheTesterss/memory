package main

import (
	"fmt"
	"memory/src/func"
	"memory/src/core"
	"os"
)

func main() {
	content, err := os.ReadFile("test.ru");
	if err != nil {
		fmt.Println("error while opening file");
		os.Exit(1);
	}

	result := core.Tokenise(string(content));
	IDENT.ExecuteFunctions(result)
}
