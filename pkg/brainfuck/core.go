package brainfuck

import (
	"strings"
)

type processType struct {
	name string
	fnc  func(pointer *int, paren *int, memory map[int]uint8, index *int)
}

type interPreter struct {
	commands   map[string]func(pointer *int, paren *int, memory map[int]uint8, index *int)
	paren      int
	pointer    int
	processes  []processType
	memory     map[int]uint8
	AddCommand func(char byte, f func(pointer *int, paren *int, memory map[int]uint8, index *int))
}

func NewInterPreter() *interPreter {
	interPreter := &interPreter{
		commands: make(map[string]func(pointer *int, paren *int, memory map[int]uint8, index *int)),
		pointer:  0,
		memory:   map[int]uint8{0: 0},
		paren:    0,
	}
	interPreter.AddCommand = func(char byte, f func(pointer *int, paren *int, memory map[int]uint8, index *int)) {
		if interPreter.commands[string(char)] == nil {
			interPreter.commands[string(char)] = f
		}
	}

	interPreter.initializeCommands()
	return interPreter
}

func (i *interPreter) Execute(code string) {
	input := strings.Split(code, "")
	i.compile(input)
	i.run()
}
