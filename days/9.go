package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input/9.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := convert(strings.Split(scanner.Text(), ","))

	instructionsWithMemory := make([]int, len(instructions)+MAX_DATA)
	copy(instructionsWithMemory, instructions)
	computer := IntComputer{
		instructions: instructionsWithMemory,
		input:        []int{2},
	}

	computer.operate()

	fmt.Println(computer.output)
}
