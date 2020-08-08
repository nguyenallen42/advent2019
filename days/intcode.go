package main

import (
	"errors"
	"fmt"
	"math"
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

const (
	MAX_DATA = 40000000
)

// IntComputer represents a computer that processes IntCode
type IntComputer struct {
	// Represents a list of instructions (opcodes) along with data
	// Note: instructions are mutable and can extend further than the original list of instructions
	instructions []int

	// Represents the instruction position that we're currently processing
	pos int

	// Represents potential input for the INPUT instruction. After an input
	// is processed, it is removed from this list
	input []int

	// Represents output from the OUTPUT instruction
	output []int

	// Stores the relative base of the computer, used for relative base parameter mode
	relativeBase int

	// Determines if we should run in debug mode (outputing instructions, etc.)
	debug bool
}

// operate runs until certain conditions:
//  - we encounter a HALT instruction
//  - we encounter an INPUT instruction when we do not have sufficient input
//  - we encounter an unknown opcode
func (ic *IntComputer) operate() error {
	for {
		instruction, modes := parseInstruction(ic.instructions[ic.pos])
		if ic.debug {
			fmt.Printf("Attempting instruction %d at pos %d w/ base %d\n", instruction.opcode, ic.pos, ic.relativeBase)
		}

		switch instruction {
		case ADD:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			val2 := ic.instructions[ic.getAddress(ic.pos+2, modes[1])]
			addr := ic.getAddress(ic.pos+3, modes[2])
			ic.instructions[addr] = val1 + val2
			ic.pos += instruction.offset
		case MULT:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			val2 := ic.instructions[ic.getAddress(ic.pos+2, modes[1])]
			addr := ic.getAddress(ic.pos+3, modes[2])
			ic.instructions[addr] = val1 * val2
			ic.pos += instruction.offset
		case INPUT:
			if ic.debug {
				fmt.Println("input: ", ic.input)
			}
			if len(ic.input) == 0 {
				return &inputError{}
			}
			addr := ic.getAddress(ic.pos+1, modes[0])
			ic.instructions[addr] = ic.input[0]
			ic.input = ic.input[1:]
			ic.pos += instruction.offset
		case OUTPUT:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			ic.output = append(ic.output, val1)
			ic.pos += instruction.offset
		case JUMP_IF_TRUE:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			val2 := ic.instructions[ic.getAddress(ic.pos+2, modes[1])]
			if val1 != 0 {
				ic.pos = val2
			} else {
				ic.pos += instruction.offset
			}
		case JUMP_IF_FALSE:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			val2 := ic.instructions[ic.getAddress(ic.pos+2, modes[1])]
			if val1 == 0 {
				ic.pos = val2
			} else {
				ic.pos += instruction.offset
			}
		case LESS_THAN:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			val2 := ic.instructions[ic.getAddress(ic.pos+2, modes[1])]
			addr := ic.getAddress(ic.pos+3, modes[2])
			if val1 < val2 {
				ic.instructions[addr] = 1
			} else {
				ic.instructions[addr] = 0
			}
			ic.pos += instruction.offset
		case ADJ_REL:
			ic.relativeBase += ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			ic.pos += instruction.offset
		case EQUAL_TO:
			val1 := ic.instructions[ic.getAddress(ic.pos+1, modes[0])]
			val2 := ic.instructions[ic.getAddress(ic.pos+2, modes[1])]
			addr := ic.getAddress(ic.pos+3, modes[2])
			if val1 == val2 {
				ic.instructions[addr] = 1
			} else {
				ic.instructions[addr] = 0
			}
			ic.pos += instruction.offset
		case HALT:
			return nil
		default:
			return errors.New("Invalid instruction")
		}
	}
}

func (ic *IntComputer) getAddress(pos, mode int) int {
	switch mode {
	case POSITION:
		return ic.instructions[pos]
	case IMMEDIATE:
		return pos
	case RELATIVE_BASE:
		return ic.instructions[pos] + ic.relativeBase
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
	// Adjusts the relative base by its only parameter
	ADJ_REL = Instruction{opcode: 9, offset: 2}
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
	ADJ_REL,
	HALT,
}

// Parameter Modes
const (
	// Uses the value at the position of the parameter
	POSITION = 0
	// Uses the value of the parameter itself
	IMMEDIATE = 1
	// Uses the value at the position of the parameter, offset by the current relativeBase
	RELATIVE_BASE = 2
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
				case RELATIVE_BASE:
					modes = append(modes, RELATIVE_BASE)
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

type inputError struct {
}

func (e *inputError) Error() string {
	return "Missing more input"
}
