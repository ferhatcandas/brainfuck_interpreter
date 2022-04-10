package main

import (
	bf "brainfuck_interpreter/pkg/brainfuck"
)

func main() {
	t := bf.NewInterPreter()
	t.Execute("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
}
