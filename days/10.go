package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Point struct {
	x int
	y int
}

// Part 1
func findBestLocation(space map[Point]bool) int {
	var maxAsteroid Point
	maxCount := 0
	for asteroid := range space {
		count := countVisibleAsteroids(asteroid, space)
		if count > maxCount {
			maxAsteroid = asteroid
			maxCount = count
		}
	}

	fmt.Println(maxAsteroid)

	return maxCount
}

func countVisibleAsteroids(asteroid Point, space map[Point]bool) int {
	count := 0
	for otherAsteroid := range space {
		if otherAsteroid == asteroid {
			continue
		}
		if isVisible(asteroid, otherAsteroid, space) {
			count += 1
		}
	}
	return count
}

// A terribly inefficient way of determining if any points are between asteroid and otherAsteroid, by looking
// at every other point
func isVisible(asteroid, otherAsteroid Point, space map[Point]bool) bool {
	totalDistance := distance(asteroid, otherAsteroid)

	for a := range space {
		if a == asteroid || a == otherAsteroid {
			continue
		}

		if floatEquals(distance(asteroid, a)+distance(a, otherAsteroid), totalDistance) {
			return false
		}
	}
	return true
}

func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2.0) + math.Pow(float64(a.y-b.y), 2.0))
}

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

// Part 2
func findNthVisible(asteroid Point, space map[Point]bool, n int) int {
	// First, find all visible other asteroids from this one
	visibleAsteroids := []Point{}
	for otherAsteroid := range space {
		if asteroid == otherAsteroid {
			continue
		}

		if isVisible(asteroid, otherAsteroid, space) {
			visibleAsteroids = append(visibleAsteroids, otherAsteroid)
		}
	}

	// Then, group the visible asteroids into quadrants
	//  0 | 1
	//  2 | 3
	quadrants := make([][]Point, 4)
	for _, visible := range visibleAsteroids {
		if visible.x >= asteroid.x && visible.y <= asteroid.y {
			quadrants[1] = append(quadrants[1], visible)
		} else if visible.x >= asteroid.x && visible.y > asteroid.y {
			quadrants[3] = append(quadrants[3], visible)
		} else if visible.x < asteroid.x && visible.y <= asteroid.y {
			quadrants[0] = append(quadrants[0], visible)
		} else {
			quadrants[2] = append(quadrants[2], visible)
		}
	}

	// Then, iterate through quadrants until we find our asteroid
	quadrantI := 0
	for _, i := range []int{1, 3, 2, 0} {
		fmt.Println(i)
		asteroids := quadrants[i]
		fmt.Println(i, len(asteroids), n, asteroids)
		fmt.Println("----")
		if len(asteroids) < n {
			n -= len(asteroids)
			continue
		} else {
			quadrantI = i
		}
	}

	// Then, sort that quadrant by slope, to find the nth asteroid
	asteroids := quadrants[quadrantI]
	sort.Slice(asteroids, func(i, j int) bool {
		return slope(asteroids[i], asteroid) < slope(asteroids[j], asteroid)
	})

	return asteroids[n-1].x*100 + asteroids[n-1].y
}

func slope(asteroid, otherAsteroid Point) float64 {
	return float64(otherAsteroid.y-asteroid.y) / float64(otherAsteroid.x-asteroid.x)
}

func main() {
	file, err := os.Open("input/10.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	space := make(map[Point]bool)

	row := 0
	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			if char == '#' {
				space[Point{x: pos, y: row}] = true
			}
		}
		row += 1
	}

	// fmt.Println(findBestLocation(space))
	fmt.Println(findNthVisible(Point{x: 20, y: 21}, space, 200))
}
