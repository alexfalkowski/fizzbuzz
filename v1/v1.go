package v1

import (
	"strconv"
)

// FizzBuzz game.
func FizzBuzz(n uint) []string {
	r := make([]string, n)

	for i := 1; i <= int(n); i++ {
		var s string

		switch {
		case i%15 == 0:
			s = "Fizz Buzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		default:
			s = strconv.Itoa(i)
		}

		r[i-1] = s
	}

	return r
}
