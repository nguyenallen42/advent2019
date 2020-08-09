package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Part 1: Just counting blocks
func countBlocks(output []int) int {
	count := 0
	for i := 2; i < len(output); i += 3 {
		if output[i] == 2 {
			count += 1
		}
	}
	return count
}

// Part 2: Simulating the game, returning the high score after breaking all the blocks
type Game struct {
	grid [][]int

	// Current score
	score int

	// Horizontal direction the ball is currently moving in (0 is right, 1 is left)
	hDirection int
	// Vertical direction the ball is currently moving in (0 is down, 1 is up)
	vDirection int

	// Current location of the ball
	ballX int
	ballY int

	// Current location of the paddle
	paddleX int

	// Where we should move our paddle next to hit the ball
	paddleGoalX int
}

const (
	WIDTH  = 50
	HEIGHT = 22

	// How long to wait before animating the next frame
	FRAME_MS = 17
)

func (g *Game) update(output []int) {
	for i := 0; i < len(output); i += 3 {
		x, y, tileId := output[i], output[i+1], output[i+2]

		if tileId == 3 {
			g.paddleX = x
		}

		if tileId == 4 {
			if x > g.ballX {
				g.hDirection = 0
			} else {
				g.hDirection = 1
			}

			if y > g.ballY {
				g.vDirection = 0
			} else {
				g.vDirection = 1
			}

			if g.vDirection == 0 {
				if g.hDirection == 0 {
					// (17, 19) -> (18, 20)
					g.paddleGoalX = x + (HEIGHT - 2 - y - 1)
				} else {
					// (17, 19) -> (16, 20)
					g.paddleGoalX = x - (HEIGHT - 2 - y)
				}
			} else {
				g.paddleGoalX = x
			}

			fmt.Printf("Previous ball: (%d, %d)\n", g.ballX, g.ballY)
			fmt.Printf("Current ball: (%d, %d)\n", x, y)
			fmt.Printf("Paddle pos %d\n", g.paddleX)
			fmt.Printf("Paddle goal %d\n", g.paddleGoalX)

			g.ballX = x
			g.ballY = y
		}

		if x == -1 {
			g.score = tileId
			continue
		}

		g.grid[y][x] = tileId
	}
}

func (g *Game) nextMove() int {
	if g.paddleX < g.paddleGoalX {
		return 1
	} else if g.paddleX > g.paddleGoalX {
		return -1
	}
	return 0
}

func (g *Game) print() {
	// ANSI Escape Sequence (https://en.wikipedia.org/wiki/ANSI_escape_code)
	// This resets cursor back to (0, 0), allowing me to rewrite to the terminal
	fmt.Print("\033[H")

	var tile int
	fmt.Println("Score: ", g.score)
	for i := 0; i < len(g.grid); i++ {
		if i < 10 {
			fmt.Printf("%d :", i)
		} else {
			fmt.Printf("%d:", i)
		}
		for j := 0; j < len(g.grid[i]); j++ {
			tile = g.grid[i][j]
			switch tile {
			case 0:
				// Empty
				fmt.Print(" ")
			case 1:
				// Wall
				fmt.Print("~")
			case 2:
				// Block
				fmt.Print("=")
			case 3:
				// Paddle
				fmt.Print("_")
			case 4:
				// Ball
				fmt.Print("o")
			default:
				fmt.Print("!")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func simulate(computer IntComputer) {
	g := Game{}
	g.grid = make([][]int, HEIGHT)
	for i := range g.grid {
		g.grid[i] = make([]int, WIDTH)
	}

	running := true
	// Playing the game
	for running {
		err := computer.operate()

		// Halt instruction, the game is finished
		if err == nil {
			running = false
		}

		g.update(computer.output)
		g.print()

		// Flush computer output, to wait for next updated frame info
		computer.output = []int{}

		// Expecting the next joystick move
		computer.input = append(computer.input, g.nextMove())

		time.Sleep(FRAME_MS * time.Millisecond)

		// Flushes the screen
		fmt.Print("\033[J")
	}
}

func main() {
	file, err := os.Open("input/13.input")
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

	simulate(computer)
}
