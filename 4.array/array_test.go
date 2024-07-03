package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2}

		got := Sum(numbers)
		want := 3

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size with variadic function", func(t *testing.T) {
		numbers := []int{1, 2}

		got := Sum2(numbers...)
		want := 3

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
