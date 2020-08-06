package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part 1
func determineFuel(mass int) int {
	fuel := mass/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

// Part 2
func determineTotalFuel(mass int) int {
	var totalFuel = 0
	fuel := determineFuel(mass)
	totalFuel += fuel

	for fuel > 0 {
		fuel = determineFuel(fuel)
		totalFuel += fuel
	}

	return totalFuel
}

func main() {
	file, err := os.Open("input/1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalFuel = 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		totalFuel += determineTotalFuel(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total fuel required: %d\n", totalFuel)
}
