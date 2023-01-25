/*
Given an array of integers, nums, and an integer, target, return the sum of two integers in the array such that they are 
closest to target.

You may assume the following: 
- that each input would have exactly one solution
- you may not use the same element twice.  
- 


Examples: 

Input: nums = [2,7,11,15], target = 9
Output: [0,1]

Explanation: 
nums[0] + nums[1] = 9
*/

package main

import (
	"fmt"
	"sort"
	"math"
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

type TwoSumSolution struct {
    minimumDifference, sum int
}


// not sure why go doesn't have a built in function for integer types...
func absVal(val int) int {
	if val < 0 {
		return -val
	}

	return val
  }

func twoSumClosest(nums []int, target int) int {
    // first thing we need to do is sort the array
	sort.Ints(nums)

	// initialize the solution
	var solution = TwoSumSolution{minimumDifference:math.MaxInt32, sum:-1}

	// and our left and right pointers
	left := 0
	right := len(nums) - 1
	
	// temp variables to use for calculating the difference and the sum at left and right
	var diff int
	var tempSum int
	
	// we need to keep searching as long as the left and right pointers don't cross.
	for left < right {
		// calculate the sum at left and right.
		tempSum = nums[left] + nums[right]
		
		// and now the distance between target and the temporary sum.
		diff = absVal(target - tempSum)

		// if that distance is now smaller than the one we have so far.
		if diff < solution.minimumDifference {
			// update it and the sum calculated at left and right.
			solution.minimumDifference = diff
			solution.sum = tempSum
		}

		// now determine which direction to move in. If we are smaller than the target..
		if tempSum < target {
			// try to make the sum bigger by moving the left hand pointer rightward so we choose a bigger value.
			left += 1
		} else {
			// otherwise, try moving the right hand pointer to the left so we make the sum smaller.
			right -= 1
		}
	}

	return solution.sum
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
		var closestSum = twoSumClosest(input, target)
		var passed = (closestSum == expected)
		fmt.Printf("Input: %v\nTarget: %d\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, target, closestSum, expected, passed)
	}
}

