/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer. 
You must write an algorithm that runs in O(n) time and without using the division operation.

Examples:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
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
	Expected []int
}
type TestCases struct {
	Tests []TestCase
}

func testEq(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }
	
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func productExceptSelf(nums []int) []int {
	var leftPrefixes = make([]int,len(nums))
    var rightPrefixes = make([]int,len(nums))
	var output = make([]int,len(nums))

	// build the left prefix products array.
	// we move from left to right
	leftPrefixes[0] = 1
	for i := 1; i < len(nums); i++ {
		var leftPrefixProduct = nums[i-1] * leftPrefixes[i-1]
		leftPrefixes[i] = leftPrefixProduct
	}

	// build the right prefix products array. 
	// we move from right to left.
	rightPrefixes[len(nums)-1] = 1
	for j := len(nums) -2; j >= 0; j-- {
		var rightPrefixProduct = nums[j+1] * rightPrefixes[j+1]
		rightPrefixes[j] = rightPrefixProduct
	}

	// the output array is constructed by multiplying the value at the k'th index
	// from right and left.  this is multiplying the product of all the numbers to the left of the k'th
	// value times all the numbers to the right of the k'th value.
	for k := range output {
		var productExceptAtk = leftPrefixes[k] * rightPrefixes[k]
		output[k] = productExceptAtk
	}
		
	return output
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
		var output = productExceptSelf(input)
		
		// a bit more work here for Go
		var passed = testEq(output, expected)
		fmt.Printf("Input: %v\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, output, expected, passed)
	}
}