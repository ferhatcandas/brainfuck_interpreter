package brainfuck

func (interPreter *interPreter) run() {
	for i := 0; i < len(interPreter.processes); {
		curr := interPreter.processes[i]
		curr.fnc(&interPreter.pointer, &interPreter.paren, interPreter.memory, &i)
	}
}
