package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Generate permutations using Heap's algorithm: https://en.wikipedia.org/wiki/Heap%27s_algorithm
func generatePermutations(arr []int) [][]int {
	var helper func(int, []int)
	result := [][]int{}

	helper = func(k int, arr []int) {
		if k == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
			return
		}

		helper(k-1, arr)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				tmp := arr[i]
				arr[i] = arr[k-1]
				arr[k-1] = tmp
			} else {
				tmp := arr[0]
				arr[0] = arr[k-1]
				arr[k-1] = tmp
			}

			helper(k-1, arr)
		}
	}

	helper(len(arr), arr)
	return result
}

// Part 1
func determineMaxSetting(instructions []int) int {
	maxOutput := 0

	for _, settings := range generatePermutations([]int{0, 1, 2, 3, 4}) {
		output := 0

		for _, setting := range settings {
			tmp := make([]int, len(instructions))
			copy(tmp, instructions)
			computer := IntComputer{
				instructions: tmp,
				input:        []int{setting, output},
			}
			computer.operate()
			output = computer.output[0]
		}

		if output > maxOutput {
			maxOutput = output
		}
	}

	return maxOutput
}

// Part 2
func determineSettingsForFeedbackLoop(instructions []int) int {
	maxOutput := 0
	for _, settings := range generatePermutations([]int{5, 6, 7, 8, 9}) {
		output := runFeedbackLoop(instructions, settings)
		if output > maxOutput {
			maxOutput = output
		}
	}
	return maxOutput
}

func runFeedbackLoop(instructions, settings []int) int {
	computers := make([]IntComputer, 5)

	// Initialize computers
	for i, setting := range settings {
		tmp := make([]int, len(instructions))
		copy(tmp, instructions)
		computer := IntComputer{
			instructions: tmp,
			input:        []int{setting},
		}
		computer.operate()

		computers[i] = computer
	}

	// Feed the output from the previous computer repeatedly until they all halt
	output := 0
	numHalted := 0
	for i := 0; numHalted < 5; i = (i + 1) % 5 {
		computers[i].input = append(computers[i].input, output)
		err := computers[i].operate()
		if err == nil {
			numHalted += 1
		}
		output = computers[i].output[len(computers[i].output)-1]
	}

	return computers[4].output[len(computers[4].output)-1]
}

func main() {
	file, err := os.Open("input/7.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := convert(strings.Split(scanner.Text(), ","))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(determineSettingsForFeedbackLoop(instructions))
}
