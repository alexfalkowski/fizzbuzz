package v2

import (
	"strconv"
)

// Rule of the game.
type Rule func(n int) (string, bool)

// NewGame with rules.
func NewGame(rules []Rule) *Game {
	return &Game{rules: rules}
}

// Game to play.
type Game struct {
	rules []Rule
}

// Play the game.
func (g *Game) Play(n uint) []string {
	r := make([]string, n)

	for i := 1; i <= int(n); i++ {
		for _, rl := range g.rules {
			if s, ok := rl(i); ok {
				r[i-1] = s

				break
			}
		}
	}

	return r
}

func none(n int) (string, bool) {
	return strconv.Itoa(n), true
}
