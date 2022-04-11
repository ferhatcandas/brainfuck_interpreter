package brainfuck

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Interpreter_Output(t *testing.T) {
	ip := NewInterPreter()
	ip.AddCommand('s', func(pointer, paren *int, memory map[int]uint8, index *int) {
		memory[*pointer] = memory[*pointer] * memory[*pointer]
		*index++
	})
	tests := []struct {
		name   string
		code   string
		output string
	}{
		{
			name:   "WHEN CODE RUNS STDOUT MUST BE Hello World!",
			code:   "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.",
			output: "Hello World!",
		},
		{
			name:   "WHEN CODE RUNS STDOUT MUST BE F, ALSO TESTED NEW COMMAND s(square the current cell)",
			code:   "++++++++s++++++.",
			output: "F",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			ip.Execute(tt.code)
			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = rescueStdout
			assert.Equal(t, tt.output, string(out))
		})
	}
}
