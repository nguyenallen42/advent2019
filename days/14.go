package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

type Ingredient struct {
	name string
	quantity int
}

var (
	FUEL = Ingredient{
		name: "FUEL",
		quantity: 1,
	}
)

func parseIngredient(ingredient string) Ingredient {
	result := strings.Split(ingredient, " ")
	resultQuantity, _ := strconv.Atoi(result[0])
	return Ingredient{
		name: result[1],
		quantity: resultQuantity,
	}
}

func determineOre(
	reactions map[string][]Ingredient,
	extra map[string]int,
	initialAmount int,
) int {
	initialFuel := Ingredient{
		name: "FUEL",
		quantity: initialAmount,
	}
	queue := make([]Ingredient, 0, len(reactions))
	queue = append(queue, initialFuel)
	count := 0

	for len(queue) > 0 {
		ingredient := queue[0]
		queue = queue[1:]

		amt, ok := extra[ingredient.name]
		if ok {
			ingredient.quantity -= amt
			delete(extra, ingredient.name)
		}

		if ingredient.name == "ORE" {
			count += ingredient.quantity
			continue
		}

		ingredients := reactions[ingredient.name]
		multiples := 0
		for ingredient.quantity > 0 {
			ingredient.quantity -= ingredients[0].quantity
			multiples += 1
		}
		for i := 1; i < len(ingredients); i++ {
			newIngredient := Ingredient{
				name: ingredients[i].name,
				quantity: multiples * ingredients[i].quantity,
			}
			queue = append(queue, newIngredient)
		}

		if ingredient.quantity < 0 {
			amt, ok := extra[ingredient.name]
			if !ok {
				extra[ingredient.name] = 0
			}
			extra[ingredient.name] = amt - ingredient.quantity
		}
	}

	return count
}

func main() {
	file, err := os.Open("input/14.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Stores ingredient => [result ingredient, other ingredients...]
	reactions := make(map[string][]Ingredient)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		recipe := strings.Split(scanner.Text(), " => ")
		ingredients := strings.Split(recipe[0], ", ")
		ingredientList := make([]Ingredient, len(ingredients) + 1)

		result := parseIngredient(recipe[1])
		ingredientList[0] = result

		for i, ingredient := range ingredients {
			ingredientList[i + 1] = parseIngredient(ingredient)
		}

		reactions[result.name] = ingredientList
	}

	extraIngredients := make(map[string]int)
	oreTotal := 1000 * 1000 * 1000 * 1000
	var oreNeeded, i int
	for i = 0; oreTotal > 0; i++ {
		oreNeeded = determineOre(reactions, extraIngredients, 1)
		oreTotal -= oreNeeded

		if len(extraIngredients) == 0 {
			fmt.Println("No more extras", i)
		}
	}
	fmt.Println("Iterations: ", i-1)
}
