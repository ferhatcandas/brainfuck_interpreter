package brainfuck

func (i *interPreter) compile(input []string) {

	for _, char := range input {
		if i.commands[char] != nil {
			i.processes = append(i.processes, processType{name: char, fnc: i.commands[char]})
		}
		//else ignore
	}
}
