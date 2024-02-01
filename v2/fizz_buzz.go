package v2

// NewFizzBuzz game.
func NewFizzBuzz() *Game {
	return NewGame([]Rule{fizzBuzz, fizz, buzz, none})
}

func fizz(n int) (string, bool) {
	return "Fizz", n%3 == 0
}

func buzz(n int) (string, bool) {
	return "Buzz", n%5 == 0
}

func fizzBuzz(n int) (string, bool) {
	return "Fizz Buzz", n%15 == 0
}
