package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input/5.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := convert(strings.Split(scanner.Text(), ","))

	// Part 1
	tmpInstructions := make([]int, len(instructions))
	copy(tmpInstructions, instructions)
	computer := IntComputer{
		instructions: tmpInstructions,
		input:        []int{1},
	}
	computer.operate()
	fmt.Println(computer.output)

	// Part 2
	copy(tmpInstructions, instructions)
	computer = IntComputer{
		instructions: tmpInstructions,
		input:        []int{5},
	}
	computer.operate()
	fmt.Println(computer.output)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
