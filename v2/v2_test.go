package v2_test

import (
	"testing"

	v2 "github.com/alexfalkowski/fizzbuzz/v2"
	"github.com/stretchr/testify/require"
)

func TestFizzBuzz(t *testing.T) {
	type test struct {
		n uint
		e []string
	}

	tests := []test{
		{0, []string{}},
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "Fizz Buzz"}},
	}

	for _, te := range tests {
		g := v2.NewFizzBuzz()
		require.Equal(t, te.e, g.Play(te.n))
	}
}
