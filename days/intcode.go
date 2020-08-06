package main

import (
	"errors"
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
		instruction, _ := parseInstruction(instructions[pos])
		switch instruction {
		case ADD:
			var pos1, pos2, result_pos = instructions[pos+1], instructions[pos+2], instructions[pos+3]
			instructions[result_pos] = instructions[pos1] + instructions[pos2]
		case MULT:
			var pos1, pos2, result_pos = instructions[pos+1], instructions[pos+2], instructions[pos+3]
			instructions[result_pos] = instructions[pos1] * instructions[pos2]
		case INPUT:
		case OUTPUT:
		case HALT:
			return instructions, nil
		default:
			return instructions, errors.New("Invalid instruction")
		}

		offset = instruction.offset
	}
	return instructions, nil
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
	if instruction == 1 {
		return ADD, []int{POSITION, POSITION, POSITION}
	} else if instruction == 2 {
		return MULT, []int{POSITION, POSITION, POSITION}
	} else {
		return HALT, nil
	}
}
