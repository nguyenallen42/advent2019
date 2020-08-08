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
	tmp := make([]int, len(instructions))
	copy(tmp, instructions)
	output, _ := operate(tmp, 1)
	fmt.Println(output)

	// Part 2
	copy(tmp, instructions)
	output, _ = operate(instructions, 5)
	fmt.Println(output)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
