package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Write a function that takes an integer n and prints a pyramid with n levels using the # character.
	// For example, if n is 3, the output should be:
	//   '  #  '
	//   ' ### '
	//   '#####'
	// If n is 4, the output should be:
	//   '   #   '
	//   '  ###  '
	//   ' ##### '
	//   '#######'
	// The base of the pyramid should be (2 * n - 1) characters wide.
	// You can assume that n is a positive integer.

	var n int = 9

	for i := range n {
		fmt.Println(Spaces(n, i) + Pyramid(strconv.Itoa(i+1), i) + Spaces(n, i))
	}
}

func Spaces(max int, quantity int) string {
	var maxLength int = len(strconv.Itoa(max))
	return strings.Repeat("•", ((max - quantity - 1) * maxLength))
}

func Pyramid(character string, quantity int) string {
	return strings.Repeat(character, quantity*2+1)
}
