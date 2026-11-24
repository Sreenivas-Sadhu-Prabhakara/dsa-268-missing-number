package main

import "fmt"

// MissingNumber returns the one value of 0..n absent from the array of n
// distinct numbers. XOR of all indexes-and-values against 0..n cancels
// everything present.
func MissingNumber(nums []int) int {
	x := len(nums) // seeds the n that has no index to pair with
	for i, v := range nums {
		x ^= i ^ v
	}
	return x
}

func main() {
	fmt.Println(MissingNumber([]int{3, 0, 1}))
	fmt.Println(MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
