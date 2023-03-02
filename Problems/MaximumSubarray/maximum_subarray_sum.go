/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestCase struct {
	Input    []int
	Expected int
	Comment  string
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

func maxSubarraySum(numbers []int) int {
	var maxSum = numbers[0]
	var runningSum = 0

	for _, number := range numbers {

		if runningSum < 0 {
			runningSum = number
		} else {
			runningSum += number
		}

		maxSum = max(runningSum, maxSum)
	}

	return maxSum
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
		var result = maxSubarraySum(input)

		// a bit more work here for Go
		var passed = result == expected
		fmt.Printf("Input: %v\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, result, expected, passed)
	}

}
