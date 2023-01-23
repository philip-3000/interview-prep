/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they 
add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Examples: 

Input: nums = [2,7,11,15], target = 9
Output: [0,1]

Explanation: 
nums[0] + nums[1] = 9
*/

package main

import "fmt"


func twoSum(nums []int, target int) []int {
    var compliments = map[int]int{}

	for idx, value := range nums {
		// for each value, we need to calculate compliment + value = target => compliment = target - value.
		// we then need to see if we've seen this value before
		var compliment = target - value

		if other_idx, ok := compliments[compliment]; ok {
			// great, we found the other number! Now we can just return the
			// answer as an array with the current index and the other index of the compliment.
			var pair = []int{ other_idx, idx }
			return pair
		}
        
		// always need to add the number, value, that we just visited.
		compliments[value] = idx
	}

	return []int{}
}

func main() {
	fmt.Printf("hello")

	// -2 + 44 => [2, 3]
	arr := []int{0, 4, 44, -2}
	target := 42
	indices := twoSum(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nIndices: %v\n\n", arr, target, indices)

	// (2 + 7) => [0, 1] 
	arr = []int{2,7,11,15}
	target = 9
	indices = twoSum(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nIndices: %v\n\n", arr, target, indices)

	// (3 + 5) => [3, 5] 
	arr = []int{0,1,2,3,4,5}
	target = 8
	indices = twoSum(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nIndices: %v\n\n", arr, target, indices)

	// (-2 + 6) => [2, 3] 
	arr = []int{0,1,-2,6,4,5}
	target = 4
	indices = twoSum(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nIndices: %v\n\n", arr, target, indices)
}

