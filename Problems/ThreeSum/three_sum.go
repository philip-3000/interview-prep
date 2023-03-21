/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] 
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
    nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
    nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
    nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
    The distinct triplets are [-1,0,1] and [-1,-1,2].
    Notice that the order of the output and the order of the triplets does not matter.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type TestCase struct {
	Input    []int
	Expected [][]int
}

type TestCases struct {
	Tests []TestCase
}

// helper function to find max value of two integers.
// (built in math.Max function only works on floats)
func max(val1 int, val2 int) int {
	if val1 < val2 {
		return val2
	}

	return val1
}

func threeSum(numbers []int) [][]int {
	// first thing we need to do is sort the input array
	sort.Ints(numbers)

	var outputs [][]int

	for idx, val := range(numbers) {

		if idx > 0 && val == numbers[idx - 1] {
			continue
		}

		if val > 0 {
			break
		}

		var left = idx + 1
		var right = len(numbers) - 1

		for left < right {
			var threeSum = val + numbers[left] + numbers[right]
			if threeSum > 0 {
				right -= 1
			} else if threeSum < 0 {
				left += 1
			} else {
				// bingo! 
				var newSolution = []int{val, numbers[left], numbers[right]}
				outputs = append(outputs, newSolution)

				// seek past any duplicates
				left += 1
				for left < right && numbers[left] == numbers[left -1] {
					left += 1
				}
			}

		}
	}

	return outputs
}

// Helper function to help compare the slices of slices to each other.
// I did this to remove any dependency on ordering of answers; instead, I convert
// them to maps and do a look up.
func compare(expected [][]int, actual [][]int) bool {
	if len(expected) != len(actual) {
		return false
	}

	// can use an array as a map key in Go.
	var expectedMap = map[[3]int]bool {}
	var actualMap = map[[3]int]bool {}

	for idx, expected_triplet := range expected {
		var actual_triplet = actual[idx]

		// must be a triplet.
		if len(actual_triplet) != 3 {
			return false
		}

		// sort the expected and actual slices.
		sort.Ints(expected_triplet)
		sort.Ints(actual_triplet)

		// now store them in a map to verify equality.
		expectedMap[[3]int { expected_triplet[0], expected_triplet[1], expected_triplet[2] }] = true
		actualMap[[3]int { actual_triplet[0], actual_triplet[1], actual_triplet[2] }] = true
	}

	for k, _ := range expectedMap {
		if ok, _ := actualMap[k]; !ok {
			return false
		}
	}
	return true
}


func main() {
	content, err := ioutil.ReadFile("./test_cases.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var testCases TestCases
	err = json.Unmarshal(content, &testCases)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)

	}


	for _, test := range testCases.Tests {
		var input = test.Input
		var expected = test.Expected
		var result = threeSum(input)

		var passed = compare(expected, result)
		fmt.Printf("Input: %v\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, result, expected, passed)

		
	}

}
