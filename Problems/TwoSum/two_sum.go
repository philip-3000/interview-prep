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
	Comment string
}
type TestCases struct {
	Tests []TestCase
}


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
		var indices = twoSum(input, target)
		
		// a bit more work here for Go
		var passed = testEq(indices, expected)
		fmt.Printf("Input: %v\nTarget: %d\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, target, indices, expected, passed)
	}
}

