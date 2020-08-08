package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// IntComputer represents a computer that processes IntCode
type IntComputer struct {
	// Represents a list of instructions (opcodes) along with data
	// Note: instructions are mutable
	instructions []int

	// Represents the instruction position that we're currently processing
	pos int

	// Represents potential input for the INPUT instruction. After an input
	// is processed, it is removed from this list
	input []int

	// Represents output from the OUTPUT instruction
	output []int

	// Determines if we should run in debug mode (outputing instructions, etc.)
	debug bool
}

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

// operate runs until certain conditions:
//  - we encounter a HALT instruction
//  - we encounter an INPUT instruction when we do not have sufficient input
//  - we encounter an unknown opcode
func (ic *IntComputer) operate() error {
	for {
		instruction, modes := parseInstruction(ic.instructions[ic.pos])
		if ic.debug {
			fmt.Printf("Attempting instruction %d at ic.pos %d\n", instruction.opcode, ic.pos)
			fmt.Println(ic.instructions)
		}

		switch instruction {
		case ADD:
			val1 := ic.getVal(ic.pos+1, modes[0])
			val2 := ic.getVal(ic.pos+2, modes[1])
			result := ic.instructions[ic.pos+3]
			ic.instructions[result] = val1 + val2
			ic.pos += instruction.offset
		case MULT:
			val1 := ic.getVal(ic.pos+1, modes[0])
			val2 := ic.getVal(ic.pos+2, modes[1])
			result := ic.instructions[ic.pos+3]
			ic.instructions[result] = val1 * val2
			ic.pos += instruction.offset
		case INPUT:
			result := ic.instructions[ic.pos+1]
			ic.instructions[result] = ic.input[0]
			ic.input = ic.input[1:]
			ic.pos += instruction.offset
		case OUTPUT:
			val1 := ic.instructions[ic.pos+1]
			ic.output = append(ic.output, ic.instructions[val1])
			ic.pos += instruction.offset
		case JUMP_IF_TRUE:
			val1 := ic.getVal(ic.pos+1, modes[0])
			val2 := ic.getVal(ic.pos+2, modes[1])
			if val1 != 0 {
				ic.pos = val2
			} else {
				ic.pos += instruction.offset
			}
		case JUMP_IF_FALSE:
			val1 := ic.getVal(ic.pos+1, modes[0])
			val2 := ic.getVal(ic.pos+2, modes[1])
			if val1 == 0 {
				ic.pos = val2
			} else {
				ic.pos += instruction.offset
			}
		case LESS_THAN:
			val1 := ic.getVal(ic.pos+1, modes[0])
			val2 := ic.getVal(ic.pos+2, modes[1])
			result := ic.instructions[ic.pos+3]
			if val1 < val2 {
				ic.instructions[result] = 1
			} else {
				ic.instructions[result] = 0
			}
			ic.pos += instruction.offset
		case EQUAL_TO:
			val1 := ic.getVal(ic.pos+1, modes[0])
			val2 := ic.getVal(ic.pos+2, modes[1])
			result := ic.instructions[ic.pos+3]
			if val1 == val2 {
				ic.instructions[result] = 1
			} else {
				ic.instructions[result] = 0
			}
			ic.pos += instruction.offset
		case HALT:
			return nil
		default:
			return errors.New("Invalid instruction")
		}
	}
}

func (ic *IntComputer) getVal(pos, mode int) int {
	switch mode {
	case POSITION:
		return ic.instructions[ic.instructions[pos]]
	case IMMEDIATE:
		return ic.instructions[pos]
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

// Instructions
var (
	// Adds two parameters and stores at the third parameter
	ADD = Instruction{opcode: 1, offset: 4}
	// Multiplies two parameters and stores at the third parameter
	MULT = Instruction{opcode: 2, offset: 4}
	// Stores input at the parameter
	INPUT = Instruction{opcode: 3, offset: 2}
	// Outputs the parameter
	OUTPUT = Instruction{opcode: 4, offset: 2}
	// Jumps the instruction pointer to the second paramter if the first parameter is true
	JUMP_IF_TRUE = Instruction{opcode: 5, offset: 3}
	// Jumps the instruction pointer to the second paramter if the first parameter is false
	JUMP_IF_FALSE = Instruction{opcode: 6, offset: 3}
	// Stores 1 if first parameter is less than second parameter, else 0
	LESS_THAN = Instruction{opcode: 7, offset: 4}
	// Stores 1 if first parameter is equal to second parameter, else 0
	EQUAL_TO = Instruction{opcode: 8, offset: 4}
	// Halts the program
	HALT = Instruction{opcode: 99, offset: 1}
)

var validInstructions = []Instruction{
	ADD,
	MULT,
	INPUT,
	OUTPUT,
	JUMP_IF_TRUE,
	JUMP_IF_FALSE,
	LESS_THAN,
	EQUAL_TO,
	HALT,
}

// Parameter Modes
const (
	// Uses the value at the position of the parameter
	POSITION = 0
	// Uses the value of the parameter itself
	IMMEDIATE = 1
)

func parseInstruction(instruction int) (Instruction, []int) {
	opcode := instruction % 100
	modes := []int{}

	for _, ins := range validInstructions {
		if ins.opcode == opcode {
			for i := 1; i < ins.offset; i++ {
				// For each parameter, the mode to read it is defined in
				// the 100th place, then the 1000th place, and so on
				mode := (instruction / int(math.Pow(10, float64(i+1)))) % 10
				switch mode {
				case POSITION:
					modes = append(modes, POSITION)
				case IMMEDIATE:
					modes = append(modes, IMMEDIATE)
				default:
					panic(errors.New("Undefined mode"))
				}
			}

			return ins, modes
		}
	}

	fmt.Println("Undefined opcode", opcode)
	panic(errors.New("Undefined opcode"))
}
