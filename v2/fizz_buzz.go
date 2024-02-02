package v2

// NewFizzBuzz game.
func NewFizzBuzz() *Game {
	return NewGame([]Rule{fizzBuzz, fizz, buzz, number})
}

func fizz(number int) (string, bool) {
	return "Fizz", number%3 == 0
}

func buzz(number int) (string, bool) {
	return "Buzz", number%5 == 0
}

func fizzBuzz(number int) (string, bool) {
	return "Fizz Buzz", number%15 == 0
}
