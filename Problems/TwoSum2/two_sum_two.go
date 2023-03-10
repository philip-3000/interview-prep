/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, 
find two numbers such that they add up to a specific target number. 
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length. 

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.
 
Your solution must use only constant extra space.
*/

package main

import (
	"fmt"
	"sort"
	"encoding/json"
    "io/ioutil"
	"log"
)

type TestCase struct {
	Input []int
	Target int
	Expected []int
}
type TestCases struct {
	Tests []TestCase
}


func twoSumTwo(input []int, target int) []int {
    // start the left on the left hand side, and right on the right hand side.
	var left = 0
	var right = len(input) - 1

	// keep going until we find the solution/the pointers cross.
	for left < right {
		var two_sum = input[left] + input[right]

		// if we found the target, stop.
		if two_sum == target {
			break
		}

		// otherewise, move right to the left if it's too big.
		if two_sum > target {
			right -= 1
		} else {
			// and left to the right if it's too small.
			left += 1
		}
	}

	// return the answer as a 1 indexed array.
	var pair = []int{ left + 1, right + 1 }
	return pair
}

func testEq(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }
	sort.Ints(a)
	sort.Ints(b)

	
    for i := range a {
        if a[i] != b[i] {
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
		var target = test.Target
		var expected = test.Expected
		var indices = twoSumTwo(input, target)
		
		// a bit more work here for Go
		var passed = testEq(indices, expected)
		fmt.Printf("Input: %v\nTarget: %d\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, target, indices, expected, passed)
	}
}

