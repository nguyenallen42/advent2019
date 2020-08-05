package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func generatePasswords(low, upper int) int {
	count := 0
	for i := low; i <= upper; i++ {
		if validateAdjDigits(i) && validateNonDecreaseDigits(i) {
			count += 1
		}
	}
	return count
}

func validateAdjDigits(candidate int) bool {
	str := strconv.Itoa(candidate)
	valid := false

	for len(str) > 1 {
		r, size := utf8.DecodeRuneInString(str)
		next_r, _ := utf8.DecodeRuneInString(str[size:])
		if r == next_r {
			valid = true
		}
		str = str[size:]
	}

	return valid
}

func validateNonDecreaseDigits(candidate int) bool {
	str := strconv.Itoa(candidate)
	valid := true

	for len(str) > 1 {
		r, size := utf8.DecodeRuneInString(str)
		next_r, _ := utf8.DecodeRuneInString(str[size:])
		if int(r - '0') > int(next_r - '0') {
			valid = false
		}
		str = str[size:]
	}

	return valid
}

func main() {
	file, err := os.Open("../input/4.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputRange := strings.Split(scanner.Text(), "-")

	lower, _ := strconv.Atoi(inputRange[0])
	upper, _ := strconv.Atoi(inputRange[1])

	fmt.Println(generatePasswords(lower, upper))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
