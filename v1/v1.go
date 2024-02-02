package v1

import (
	"strconv"
)

// FizzBuzz game.
func FizzBuzz(total uint) []string {
	words := make([]string, total)

	for number := 1; number <= int(total); number++ {
		var word string

		switch {
		case number%15 == 0:
			word = "Fizz Buzz"
		case number%3 == 0:
			word = "Fizz"
		case number%5 == 0:
			word = "Buzz"
		default:
			word = strconv.Itoa(number)
		}

		words[number-1] = word
	}

	return words
}
