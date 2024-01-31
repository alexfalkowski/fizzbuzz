package v1_test

import (
	"testing"

	v1 "github.com/alexfalkowski/fizzbuzz/v1"
	"github.com/stretchr/testify/require"
)

func TestEmpty(t *testing.T) {
	r := v1.FizzBuzz(0)
	require.Empty(t, r)
}

func TestFizz(t *testing.T) {
	r := v1.FizzBuzz(3)
	require.Equal(t, []string{"1", "2", "Fizz"}, r)
}

func TestBuzz(t *testing.T) {
	r := v1.FizzBuzz(5)
	require.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, r)
}

func TestFizzBuzz(t *testing.T) {
	e := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "Fizz Buzz"}
	r := v1.FizzBuzz(15)
	require.Equal(t, e, r)
}
