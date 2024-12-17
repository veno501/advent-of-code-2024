package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type State struct {
	A                  int
	B                  int
	C                  int
	instructionPointer int
}

func getComboOperand(state State, operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return state.A
	case 5:
		return state.B
	case 6:
		return state.C
	}
	return -1
}

func runProgram(A int, B int, C int, program []int) string {
	var output []string
	r := State{A, B, C, 0}
	for r.instructionPointer < len(program) {
		opcode := program[r.instructionPointer]
		operand := program[r.instructionPointer+1]
		switch opcode {
		case 0:
			// adv
			denominator := 1 << getComboOperand(r, operand)
			r.A = r.A / denominator
		case 1:
			// bxl
			r.B = r.B ^ operand
		case 2:
			// bst
			r.B = getComboOperand(r, operand) % 8
		case 3:
			// jnz
			if r.A != 0 {
				r.instructionPointer = operand
				continue
			}
		case 4:
			// bxc
			r.B = r.B ^ r.C
		case 5:
			// out
			next := getComboOperand(r, operand) % 8
			output = append(output, strconv.Itoa(next))
		case 6:
			// bdv
			denominator := 1 << getComboOperand(r, operand)
			r.B = r.A / denominator
		case 7:
			// cdv
			denominator := 1 << getComboOperand(r, operand)
			r.C = r.A / denominator
		default:
		}
		r.instructionPointer += 2
	}
	return strings.Join(output, ",")
}

func main() {
	cwd, _ := os.Getwd()
	filePath := filepath.Join(cwd, "input_program.txt")
	inputBytes, _ := os.ReadFile(filePath)
	input := bytes.Split(inputBytes, []byte("\r\n\r\n"))
	regs := strings.Split(string(input[0]), "\r\n")
	A, _ := strconv.Atoi(regs[0][strings.Index(regs[0], ": ")+2:])
	B, _ := strconv.Atoi(regs[1][strings.Index(regs[1], ": ")+2:])
	C, _ := strconv.Atoi(regs[2][strings.Index(regs[2], ": ")+2:])
	programStr := strings.Split(strings.Split(string(input[1]), ": ")[1], ",")
	program := make([]int, len(programStr))
	for i := 0; i < len(programStr); i++ {
		program[i], _ = strconv.Atoi(programStr[i])
	}

	fmt.Printf("Program output: %s\n", runProgram(A, B, C, program))

	iter := 100
	cmp := strings.Join(programStr, ",")
	for true {
		str := runProgram(iter, B, C, program)
		if str == cmp {
			break
		}
		if strings.HasSuffix(cmp, str) {
			iter *= 8
		} else {
			iter += 1
		}
	}
	fmt.Printf("Register A initial value that results in the program outputting itself: %d\n", iter)
}
