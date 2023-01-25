/*
Given an array of integers, nums, and an integer, target, find the number of unique pairs in the array
such that their sum is equal to the target. Return the number of pairs.

Example:

Input: nums = [1,2,3,4], target = 6
output: 2

Explanation: 
2 + 4 = 6

*/

package main

import (
	"fmt"
	"encoding/json"
    "io/ioutil"
	"log"
)

type TestCase struct {
	Input []int
	Target int
	Expected int
	Comment string
}
type TestCases struct {
	Tests []TestCase
}

// A set like structure of integers.
type IntSet map[int]struct{}

// Represents a tuple of 2 elements
type Pair struct {
    x, y int
}

// Returns the number of unique pairs of integers in 
// nums that add up to target.
func twoSumUniquePairs(nums []int, target int) int {

	// let's keep track of the numbers we've seen already.  We'll need to look these up.
	var compliments = IntSet{}

	// and the unique solutions
	seen := make(map[Pair]struct{})
	
	// go through each number in the list...
	for _, value := range nums {
		// for each value, we need to calculate compliment + value = target => compliment = target - value.
		// we then need to see if we've seen this value before
		var compliment = target - value

		if _, ok := compliments[compliment]; ok {
			// great, we found the other number! 

			// always want to store the smaller number first
			var pair = Pair{compliment, value}
			if value < compliment {
				pair.x = value
				pair.y = compliment
			} 
			
			// add the pair into our seen solutions.
			seen[pair] = struct{}{}
		}
        
		// always need to add the number, value, that we just visited.
		compliments[value] = struct{}{}

    }

	return len(seen)
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
		var target = test.Target
		var expected = test.Expected
		var closestSum = twoSumUniquePairs(input, target)
		var passed = (closestSum == expected)
		fmt.Printf("Input: %v\nTarget: %d\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, target, closestSum, expected, passed)
	}
}