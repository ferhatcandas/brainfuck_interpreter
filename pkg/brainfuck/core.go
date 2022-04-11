package brainfuck

import (
	"bufio"
	"os"
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
	i.clear()
}
func (i *interPreter) clear() {
	i.paren = 0
	i.memory = make(map[int]uint8)
	i.pointer = 0
	i.processes = []processType{}
}

func (i *interPreter) ExecuteFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}
	i.Execute(input)
}
