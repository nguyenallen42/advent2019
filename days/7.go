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

func determineMaxSetting(instructions []int) int {
	maxOutput := 0

	for _, settings := range generatePermutations([]int{0, 1, 2, 3, 4}) {
		tmpInstructions := make([]int, len(instructions))
		copy(tmpInstructions, instructions)
		output := 0

		for _, setting := range settings {
			outputs, _ := operate(tmpInstructions, setting, output)
			output = outputs[0]
		}

		if output > maxOutput {
			maxOutput = output
		}
	}

	return maxOutput
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

	fmt.Println(determineMaxSetting(instructions))
}
