package main

import (
	"fmt"
)

type Moon struct {
	// Position
	x int
	y int
	z int

	// Velocites
	xv int
	yv int
	zv int
}

func (m *Moon) getCoord(c string) int {
	switch c {
	case "x":
		return m.x
	case "y":
		return m.y
	case "z":
		return m.z
	}
	return 0
}

func (m *Moon) updateCoord(c string, val int) {
	switch c {
	case "x":
		m.x += val
	case "y":
		m.y += val
	case "z":
		m.z += val
	}
}

func (m *Moon) getVel(v string) int {
	switch v {
	case "xv":
		return m.xv
	case "yv":
		return m.yv
	case "zv":
		return m.zv
	}
	return 0
}

func (m *Moon) updateVel(v string, val int) {
	switch v {
	case "xv":
		m.xv += val
	case "yv":
		m.yv += val
	case "zv":
		m.zv += val
	}
}

// Part 1
func simulate(moons []Moon) {
	for _, moon := range moons {
		fmt.Printf(
			"pos <%d, %d, %d>, vel <%d, %d, %d>\n",
			moon.x, moon.y, moon.z,
			moon.xv, moon.yv, moon.zv,
		)
	}
	fmt.Println()

	for i := 1; i <= 1000; i++ {
		applyGravity(moons)
		updatePositions(moons)

		fmt.Println("Iteration", i)
		for _, moon := range moons {
			fmt.Printf(
				"pos <%d, %d, %d>, vel <%d, %d, %d>\n",
				moon.x, moon.y, moon.z,
				moon.xv, moon.yv, moon.zv,
			)
		}
		fmt.Println()
	}
}

func applyGravity(moons []Moon) {
	for i := 0; i < len(moons); i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].x > moons[j].x {
				moons[i].xv -= 1
				moons[j].xv += 1
			} else if moons[i].x < moons[j].x {
				moons[j].xv -= 1
				moons[i].xv += 1
			}

			if moons[i].y > moons[j].y {
				moons[i].yv -= 1
				moons[j].yv += 1
			} else if moons[i].y < moons[j].y {
				moons[j].yv -= 1
				moons[i].yv += 1
			}

			if moons[i].z > moons[j].z {
				moons[i].zv -= 1
				moons[j].zv += 1
			} else if moons[i].z < moons[j].z {
				moons[j].zv -= 1
				moons[i].zv += 1
			}
		}
	}
}

func updatePositions(moons []Moon) {
	for i := 0; i < len(moons); i++ {
		moons[i].x += moons[i].xv
		moons[i].y += moons[i].yv
		moons[i].z += moons[i].zv
	}
}

func computeEnergy(moons []Moon) int {
	totalEnergy := 0

	for _, moon := range moons {
		potential := Abs(moon.x) + Abs(moon.y) + Abs(moon.z)
		kinetic := Abs(moon.xv) + Abs(moon.yv) + Abs(moon.zv)
		totalEnergy += potential * kinetic
	}

	return totalEnergy
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Part 2
func simulateByCoords(moons []Moon) {
	fmt.Println("x: ", computeCycleForCoord(moons, "x", "xv"))
	fmt.Println("y: ", computeCycleForCoord(moons, "y", "yv"))
	fmt.Println("z: ", computeCycleForCoord(moons, "z", "zv"))
}

func computeCycleForCoord(moons []Moon, c, v string) int {
	origCoords, origVel := []int{}, []int{}
	for _, moon := range moons {
		origCoords = append(origCoords, moon.getCoord(c))
		origVel = append(origVel, moon.getVel(v))
	}

	done := false
	count := 0
	for !done {
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				if moons[i].getCoord(c) > moons[j].getCoord(c) {
					moons[i].updateVel(v, -1)
					moons[j].updateVel(v, 1)
				} else if moons[i].getCoord(c) < moons[j].getCoord(c) {
					moons[i].updateVel(v, 1)
					moons[j].updateVel(v, -1)
				}
			}
		}

		for i := 0; i < len(moons); i++ {
			moons[i].updateCoord(c, moons[i].getVel(v))
		}

		done = true
		for i := range moons {
			if moons[i].getCoord(c) != origCoords[i] {
				done = false
			}
			if moons[i].getVel(v) != origVel[i] {
				done = false
			}
		}
		count += 1
	}
	return count
}

func main() {
	// Because screw reading from a file
	moons := []Moon{
		Moon{x: 19, y: -10, z: 7},
		Moon{x: 1, y: 2, z: -3},
		Moon{x: 14, y: -4, z: 1},
		Moon{x: 8, y: 7, z: -6},
		// Moon{x:-1, y:0, z:2},
		// Moon{x:2, y:-10, z:-7},
		// Moon{x:4, y:-8, z:8},
		// Moon{x:3, y:5, z:-1},
	}

	// After simulating for each coordinate, find the LCM
	simulateByCoords(moons)
}
