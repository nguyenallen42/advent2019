package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func explore(instructions []int, expected_output int) (int, int) {
	/* Explores by:
	- modifying the instructions using nouns and verbs
	- operating until we reach a halt
	- stopping once the 0th instruction is equivalent to the expected_output
	*/
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			new_instructions := make([]int, len(instructions))
			copy(new_instructions, instructions)

			new_instructions[1], new_instructions[2] = noun, verb
			operate(new_instructions)

			if new_instructions[0] == expected_output {
				fmt.Printf("(noun=%d, verb=%d)\n", noun, verb)
				return noun, verb
			}
		}
	}
	return -1, -1
}

func main() {
	file, err := os.Open("input/2.input")
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
