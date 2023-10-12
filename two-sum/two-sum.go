package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 6))
}

func twoSum(nums []int, target int) []int {

	result := []int{}

	for i := 0; i < len(nums); i++ {
		for b := i + 1; b < len(nums); b++ {
			if (nums[i] + nums[b]) == target {
				result = append(result, i, b)
				return result
			}
		}
	}

	return result
}
