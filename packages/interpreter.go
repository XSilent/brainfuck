package packages

import (
	"fmt"
	"io/ioutil"
	"os"
)

const cellInc = 43      // +
const cellDec = 45      // -
const pointerInc = 62   // >
const pointerDec = 60   // <
const cellOutput = 46   // .
const cellInput = 44    // ,
const jumpForward = 91  // [
const jumpBackward = 93 // ]

type Interpreter struct {
	source string
}

var band [30000]byte
var pointer byte

func (i *Interpreter) Run(source []byte) {

	pointer = 0
	j := 0

	for j < len(source) {
		switch source[j] {
		case cellInc:
			band[pointer]++
		case cellDec:
			band[pointer]--
		case pointerInc:
			pointer++
		case pointerDec:
			pointer--
		case cellInput:
			//
		case cellOutput:
			fmt.Print(string(band[pointer]))
		case jumpForward:
			if band[pointer] == 0 {
				for source[j+1] != jumpBackward {
					j++
				}
			}
		case jumpBackward:
			if band[pointer] != 0 {
				for source[j-1] != jumpForward {
					j--
				}
			}
		}

		j++
	}

	fmt.Println("")
}

func (i *Interpreter) RunFromFile(file string) {
	code, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Source file not found!")
		os.Exit(-1)
	}

	i.Run(code)
}
