package main

import (
	"fmt"
	"os"

	bf "github.com/ferhatcandas/brainfuck_interpreter/pkg/brainfuck"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments missing")
		os.Exit(1)
	}
	if os.Args[1] == "exec" {
		if len(os.Args) < 3 {
			fmt.Println("Code is missing")
			os.Exit(1)
		}
		bf.NewInterPreter().Execute(os.Args[2])
	} else if os.Args[1] == "file" {
		if len(os.Args) < 3 {
			fmt.Println("File path is missing")
			os.Exit(1)
		}
		bf.NewInterPreter().ExecuteFromFile(os.Args[2])
	}
}
