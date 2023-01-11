package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type code int

const (
	codeInc code = iota
	codeDec
	codeAdd
	codeSub
	codeIn
	codeOut
	codeJnz
	codeJmp
)

type instruction struct {
	code
	op int
}

type instructions []*instruction

func (i *instructions) addCode(code code) {
	i.addCodeWithOp(code, -1)
}

func (i *instructions) addCodeWithOp(code code, op int) {
	*i = append(*i, &instruction{code, op})
}

func compile(program string) (instructions, error) {
	var lookup Stack

	instructions := make(instructions, 0)

	for _, char := range program {
		switch char {
		case '>':
			instructions.addCode(codeInc)
		case '<':
			instructions.addCode(codeDec)
		case '+':
			instructions.addCode(codeAdd)
		case '-':
			instructions.addCode(codeSub)
		case ',':
			instructions.addCode(codeIn)
		case '.':
			instructions.addCode(codeOut)
		case '[':
			lookup.Push(len(instructions))
			instructions.addCode(codeJnz)
		case ']':
			if val, ok := lookup.Pop(); !ok {
				return nil, errors.New("no open loop found")
			} else {
				instructions.addCodeWithOp(codeJmp, val)
				instructions[val].op = len(instructions)
			}
		default:
			continue
		}
	}

	return instructions, nil
}

func Run(program string, memorySize int) error {
	var ptr int

	memory := make([]byte, memorySize)

	instructions, err := compile(program)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(instructions); {
		instruction := instructions[i]
		switch instruction.code {
		case codeInc:
			ptr++
		case codeDec:
			ptr--
		case codeAdd:
			memory[ptr]++
		case codeSub:
			memory[ptr]--
		case codeOut:
			if _, err := fmt.Printf("%c", memory[ptr]); err != nil {
				return nil
			}
		case codeIn:
			if val, err := reader.ReadByte(); err != nil {
				return err
			} else {
				memory[ptr] = val
			}
		case codeJnz:
			if memory[ptr] == 0 {
				i = instruction.op
				continue
			}
		case codeJmp:
			i = instruction.op
			continue
		}
		i++
	}

	return nil
}
