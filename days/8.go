package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"unicode/utf8"
)

type Layer struct {
	id int
	pixels [][]int
}

const (
	WIDTH = 25
	HEIGHT = 6
)

func constructLayers(input string) []Layer {
	numLayers := len(input) / (WIDTH * HEIGHT)
	layers := make([]Layer, numLayers)

	for i := 0; i < numLayers; i++ {
		layer := Layer{
			id: i + 1,
			pixels: [][]int{},
		}
		for h := 0; h < HEIGHT; h++ {
			row := make([]int, WIDTH)
			for w := 0; w < WIDTH; w++ {
				r, size := utf8.DecodeRuneInString(input)
				row[w] = int(r - '0')
				input = input[size:]
			}
			layer.pixels = append(layer.pixels, row)
		}
		layers[i] = layer
	}

	return layers
}

// Part 1
func determineFewestZeros(layers []Layer) int {
	fewestLayerIdx := 0
	fewestZeros := math.MaxUint32

	for i, layer := range layers {
		zeros := countDigits(layer, 0)
		if zeros < fewestZeros {
			fewestZeros = zeros
			fewestLayerIdx = i
		}
	}

	return countDigits(layers[fewestLayerIdx], 1) * countDigits(layers[fewestLayerIdx], 2)
}

func countDigits(layer Layer, digit int) int {
	count := 0
	for h := 0; h < HEIGHT; h++ {
		for w := 0; w < WIDTH; w++ {
			if layer.pixels[h][w] == digit {
				count += 1
			}
		}
	}

	return count
}

// Part 2
func computeImage(layers []Layer) Layer {
	pixels := make([][]int, HEIGHT)
	for i := range pixels {
		pixels[i] = make([]int, WIDTH)
	}
	resultLayer := Layer{
		id: 0,
		pixels: pixels,
	}

	for h := 0; h < HEIGHT; h++ {
		for w := 0; w < WIDTH; w++ {
			resultLayer.pixels[h][w] = 2
			for i := range layers {
				if layers[i].pixels[h][w] != 2 {
					resultLayer.pixels[h][w] = layers[i].pixels[h][w]
					break
				}
			}
		}
	}

	return resultLayer
}

func main() {
	file, err := os.Open("input/8.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	layers := constructLayers(input)

	image := computeImage(layers)

	for row := range image.pixels {
		fmt.Println(image.pixels[row])
	}

}
