package main

import (
	"fmt"
)

/*
Two Sum Problem

Write a function that takes a slice of integers and a target integer,
and returns the indices of the two numbers that add up to the target.

Constraints:
- Do not use the same element twice
- Exactly one solution is guaranteed
- You can return the result in any order

Example:
nums := []int{2, 7, 11, 15}
target := 9
Output: [0, 1] because nums[0] + nums[1] == 9
*/

// TwoSum finds two indices in nums whose values add up to target.
// It returns the indices as a slice of two integers.
// Assumes exactly one solution exists.
func TwoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if val+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("two sum")

}
