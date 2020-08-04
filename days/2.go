package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Part 1
func operate(instructions []int) ([]int, error) {
	// Operates on an array of instructions
	for pos := 0; pos < len(instructions); pos += 4 {
		if instructions[pos] == 99 {
			return instructions, nil
		}

		var pos1, pos2, result_pos = instructions[pos+1], instructions[pos+2], instructions[pos+3]
		if instructions[pos] == 1 {
			instructions[result_pos] = instructions[pos1] + instructions[pos2]
		} else if instructions[pos] == 2 {
			instructions[result_pos] = instructions[pos1] * instructions[pos2]
		} else {
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

// Part 2
func explore(instructions []int, expected_output int) (int, int) {
	/* Explores by:
	- modifying the instructions using nouns and verbs
	- operating until we reach a halt
	- stopping once the 0th instruction is equivalent to the expected_output
	*/
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			var new_instructions = make([]int, len(instructions))
			copy(new_instructions, instructions)

			new_instructions[1], new_instructions[2] = noun, verb
			new_instructions, err := operate(new_instructions)

			if err == nil && new_instructions[0] == expected_output {
				fmt.Printf("(noun=%d, verb=%d)\n", noun, verb)
				return noun, verb
			}
		}
	}
	return -1, -1
}

func main() {
	file, err := os.Open("../input/2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := strings.Split(scanner.Text(), ",")
	if err != nil {
		log.Fatal(err)
	}

	noun, verb := explore(convert(instructions), 19690720)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(100*noun + verb)
}
