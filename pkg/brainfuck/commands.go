package brainfuck

import (
	"fmt"
	"os"
)

func (i *interPreter) initializeCommands() {
	i.AddCommand('>', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		*pointer++
		*index++
	})
	i.AddCommand('<', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		*pointer--
		*index++
	})
	i.AddCommand('+', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		memory[*pointer]++
		*index++
	})
	i.AddCommand('-', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		memory[*pointer]--
		*index++
	})
	i.AddCommand('[', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		if memory[*pointer] == 0 {
			*paren++
			for {
				*index++
				if *index >= len(i.processes) {
					fmt.Println("Loop end not found")
					os.Exit(1)
				}
				if i.processes[*index].name == "[" {
					*paren++
				} else if i.processes[*index].name == "]" {
					*paren--
				}
				if *paren == 0 {
					break
				}
			}
		} else {
			*index++
		}
	})
	i.AddCommand(']', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		if memory[*pointer] != 0 {
			*paren--
			for {
				*index--
				if *index < 0 {
					fmt.Println("Loop start not found")
					os.Exit(1)
				}
				if i.processes[*index].name == "[" {
					*paren++
				} else if i.processes[*index].name == "]" {
					*paren--
				}
				if *paren == 0 {
					break
				}
			}
		} else {
			*index++
		}
	})
	i.AddCommand('.', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		_, err := fmt.Printf("%c", memory[*pointer])
		if err != nil {
			panic(err)
		}
		*index++
	})
	i.AddCommand(',', func(pointer *int, paren *int, memory map[int]uint8, index *int) {
		var input uint8
		fmt.Scanf("%c", &input)
		memory[*pointer] = input
		*index++
	})
}
