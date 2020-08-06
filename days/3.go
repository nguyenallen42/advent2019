package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Point struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Part 1
func process(first_wire, second_wire []string) {
	grid := make(map[Point]int)
	overlap := make(map[Point]int)

	// Populate our grid by traversing our first wire, starting at (0, 0)
	origin, distance := Point{0, 0}, 0
	for _, direction := range first_wire {
		origin, distance = processDirection(direction, origin, distance, grid, nil)
	}

	// Traverse our second wire, tracking positions that overlap
	origin, distance = Point{0, 0}, 0
	for _, direction := range second_wire {
		origin, distance = processDirection(direction, origin, distance, grid, overlap)
	}

	// Return the closest Point to the origin, based on Manhattan distance
	minDistance := math.MaxUint32
	for _, distance := range overlap {
		if distance < minDistance {
			minDistance = distance
		}
	}

	fmt.Println(minDistance)
}

func processDirection(direction string, origin Point, distance int, grid, overlap map[Point]int) (Point, int) {
	orientation, size := utf8.DecodeRuneInString(direction)
	length, _ := strconv.Atoi(direction[size:])

	var new_point Point
	for i := 1; i <= length; i++ {
		switch orientation {
		case 'R':
			new_point = Point{origin.x + i, origin.y}
		case 'U':
			new_point = Point{origin.x, origin.y + i}
		case 'D':
			new_point = Point{origin.x, origin.y - i}
		case 'L':
			new_point = Point{origin.x - i, origin.y}
		default:
			panic(errors.New("Invalid direction"))
		}

		if overlap != nil {
			first_distance, ok := grid[new_point]
			if ok {
				overlap[new_point] = distance + i + first_distance
			}
		} else {
			grid[new_point] = distance + i
		}
	}
	return new_point, distance + length
}

func main() {
	file, err := os.Open("input/3.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	first_wire := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	second_wire := strings.Split(scanner.Text(), ",")

	fmt.Println(first_wire)
	fmt.Println(second_wire)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	process(first_wire, second_wire)
}
