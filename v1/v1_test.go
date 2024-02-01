package v1_test

import (
	"testing"

	v1 "github.com/alexfalkowski/fizzbuzz/v1"
	"github.com/stretchr/testify/require"
)

func TestV1(t *testing.T) {
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
		r := v1.FizzBuzz(te.n)
		require.Equal(t, te.e, r)
	}
}
