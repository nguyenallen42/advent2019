package main

import (
	"errors"
	"strconv"
)


func operate(instructions []int) ([]int, error) {
	// Operates on an array of instructions
	for pos := 0; pos < len(instructions); pos += 4 {
		switch instructions[pos] {
		case 1:
			var pos1, pos2, result_pos = instructions[pos+1], instructions[pos+2], instructions[pos+3]
			instructions[result_pos] = instructions[pos1] + instructions[pos2]
		case 2:
			var pos1, pos2, result_pos = instructions[pos+1], instructions[pos+2], instructions[pos+3]
			instructions[result_pos] = instructions[pos1] * instructions[pos2]
		case 99:
			return instructions, nil
		default:
			return instructions, errors.New("Invalid instruction")
		}
	}
	return instructions, nil
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
