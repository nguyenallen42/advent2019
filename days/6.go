package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type ObjAndDepth struct {
	obj string
	depth int
}

// Part 1
func traverse(space map[string][]string) int {
	visited := make(map[string]bool)
	total := 0
	queue := []ObjAndDepth{
		ObjAndDepth{
			obj: "COM",
			depth: 0,
		},
	}

	for len(queue) > 0 {
		objAndDepth := queue[0]
		queue = queue[1:]

		_, ok := visited[objAndDepth.obj]

		if ok {
			continue
		}

		visited[objAndDepth.obj] = true
		total += objAndDepth.depth
		for _, neighbor := range space[objAndDepth.obj] {
			queue = append(queue, ObjAndDepth{
				obj: neighbor,
				depth: objAndDepth.depth+1,
			})
		}
	}

	return total
}

// Part 2
func dfsPath(space map[string][]string, goal string) []string {
	visited := make(map[string]bool)

	for _, neighbor := range space["COM"] {
		result, ok := dfsHelper(space, neighbor, goal, visited)
		if ok {
			return append([]string{"COM"}, result...)
		}
	}

	return nil
}

func dfsHelper(space map[string][]string, current, goal string, visited map[string]bool) ([]string, bool) {
	if current == goal {
		return []string{current}, true
	}

	_, ok := visited[current]
	if ok {
		return nil, false
	}

	visited[current] = true
	for _, neighbor := range space[current] {
		result, ok := dfsHelper(space, neighbor, goal, visited)
		if ok {
			return append([]string{current}, result...), true
		}
	}

	return nil, false
}

func transition(pathToSan, pathToYou []string) int {
	minPathLength := len(pathToSan)
	if len(pathToYou) < len(pathToSan) {
		minPathLength = len(pathToYou)
	}

	i := 0
	for ; i < minPathLength; i++ {
		if pathToSan[i] != pathToYou[i] {
			break
		}
	}

	fmt.Println(i)

	return (len(pathToSan) - i - 1) + (len(pathToYou) - i - 1)
}

func main() {
	file, err := os.Open("input/6.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Track our entire system
	space := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		orbits := strings.Split(scanner.Text(), ")")
		if err != nil {
			log.Fatal(err)
		}

		space[orbits[0]] = append(space[orbits[0]], orbits[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	pathToSan := dfsPath(space, "SAN")
	pathToYou := dfsPath(space, "YOU")

	fmt.Println(pathToSan)
	fmt.Println(pathToYou)

	fmt.Println(transition(pathToSan, pathToYou))
}
