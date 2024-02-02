package v2

import (
	"strconv"
)

// Rule of the game.
type Rule func(number int) (string, bool)

// NewGame with rules.
func NewGame(rules []Rule) *Game {
	return &Game{rules: rules}
}

// Game to play.
type Game struct {
	rules []Rule
}

// Play the game.
func (g *Game) Play(total uint) []string {
	words := make([]string, total)

	for number := 1; number <= int(total); number++ {
		for _, rule := range g.rules {
			if word, ok := rule(number); ok {
				words[number-1] = word

				break
			}
		}
	}

	return words
}

func number(number int) (string, bool) {
	return strconv.Itoa(number), true
}
