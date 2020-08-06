package main

import (
	"errors"
	"fmt"
	"strconv"
)

func convert(instructions []string) []int {
	// Converts instructions from []string to []int, for ease of use
	var converted_instructions = []int{}
	for _, i := range instructions {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		converted_instructions = append(converted_instructions, j)
	}
	return converted_instructions
}

func operate(instructions []int) ([]int, error) {
	var offset int

	for pos := 0; pos < len(instructions); pos += offset {
		instruction, modes := parseInstruction(instructions[pos])
		switch instruction {
		case ADD:
			val1 := getVal(instructions, instructions[pos+1], modes[0])
			val2 := getVal(instructions, instructions[pos+2], modes[1])
			result := instructions[pos+3]
			instructions[result] = val1 + val2
		case MULT:
			val1 := getVal(instructions, instructions[pos+1], modes[0])
			val2 := getVal(instructions, instructions[pos+2], modes[1])
			result := instructions[pos+3]
			instructions[result] = val1 * val2
		case INPUT:
			fmt.Print("Input: ")
			var inputVal int
			fmt.Scanf("%d", &inputVal)
			result := instructions[pos+1]
			instructions[result] = inputVal
		case OUTPUT:
			val1 := instructions[pos+1]
			fmt.Println(instructions[val1])
		case HALT:
			return instructions, nil
		default:
			return instructions, errors.New("Invalid instruction")
		}

		offset = instruction.offset
	}
	return instructions, nil
}

func getVal(instructions []int, parameter, mode int) int {
	switch mode {
	case POSITION:
		return instructions[parameter]
	case IMMEDIATE:
		return parameter
	default:
		panic(errors.New("Invalid mode"))
	}

}

type Instruction struct {
	// The opcode that defines this instruction
	opcode int
	// How many to offset the instruction pointer by when determining
	// the next instruction to read. The offset should always be
	// 1 (for the opcode itself) + the number of parameters
	offset int
}

// Valid Instructions
var (
	// Adds two parameters and stores at the third parameter
	ADD = Instruction{opcode: 1, offset: 4}
	// Multiplies two parameters and stores at the third parameter
	MULT = Instruction{opcode: 2, offset: 4}
	// Stores input at the parameter
	INPUT = Instruction{opcode: 3, offset: 2}
	// Outputs the parameter
	OUTPUT = Instruction{opcode: 4, offset: 2}
	// Halts the program
	HALT = Instruction{opcode: 99, offset: 1}
)

// Parameter Modes
const (
	// Uses the value at the position of the parameter
	POSITION = 0
	// Uses the value of the parameter itself
	IMMEDIATE = 1
)

func parseInstruction(instruction int) (Instruction, []int) {
	opcode := instruction % 100

	switch opcode {
	case ADD.opcode:
		return ADD, []int{
			(instruction / 100) % 10,
			(instruction / 1000) % 10,
			(instruction / 10000) % 10,
		}
	case MULT.opcode:
		return MULT, []int{
			(instruction / 100) % 10,
			(instruction / 1000) % 10,
			(instruction / 10000) % 10,
		}
	case INPUT.opcode:
		return INPUT, []int{
			(instruction / 100) % 10,
		}
	case OUTPUT.opcode:
		return OUTPUT, []int{
			(instruction / 100) % 10,
		}
	case HALT.opcode:
		return HALT, []int{}
	default:
		fmt.Println("Undefined opcode", opcode)
		panic(errors.New("Undefined opcode"))
	}
}
