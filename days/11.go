package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

var directionLeft = map[string]string{
	"n": "w",
	"w": "s",
	"s": "e",
	"e": "n",
}

var directionRight = map[string]string{
	"n": "e",
	"e": "s",
	"s": "w",
	"w": "n",
}

const (
	WIDTH  = 43
	HEIGHT = 8
)

func paint(computer IntComputer) {
	// 0 = black, 1 = white
	hull := make([][]int, HEIGHT*2+1)
	for i := range hull {
		hull[i] = make([]int, WIDTH*2+1)
	}
	hull[HEIGHT][WIDTH] = 1

	currentPoint := Point{x: 0, y: 0}
	direction := "n"

	for {
		currentColor := hull[currentPoint.y+HEIGHT][currentPoint.x+WIDTH]

		computer.input = append(computer.input, currentColor)

		err := computer.operate()

		if err == nil {
			break
		}

		colorToPaint := computer.output[len(computer.output)-2]
		directionToTurn := computer.output[len(computer.output)-1]

		hull[currentPoint.y+HEIGHT][currentPoint.x+WIDTH] = colorToPaint

		if directionToTurn == 0 {
			direction = directionLeft[direction]
		} else {
			direction = directionRight[direction]
		}

		switch direction {
		case "n":
			currentPoint = Point{x: currentPoint.x, y: currentPoint.y + 1}
		case "w":
			currentPoint = Point{x: currentPoint.x - 1, y: currentPoint.y}
		case "s":
			currentPoint = Point{x: currentPoint.x, y: currentPoint.y - 1}
		case "e":
			currentPoint = Point{x: currentPoint.x + 1, y: currentPoint.y}
		default:
			panic("Bad direction input")
		}

	}

	for i := len(hull) - 1; i >= 0; i-- {
		for j := 0; j < len(hull[i]); j++ {
			if hull[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("input/11.input")
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
	}

	paint(computer)
}
