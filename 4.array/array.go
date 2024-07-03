package main

import "fmt"

func Sum(nums []int) int {
	var sum int

	for _, number := range nums {
		sum += number
	}

	return sum
}

func Sum2(nums ...int) int {
	var sum int

	for _, number := range nums {
		sum += number
	}

	fmt.Println(nums)

	return sum
}
