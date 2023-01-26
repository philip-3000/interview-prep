/* 
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
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
	Expected int
	Comment string
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

func buySellStock(prices []int) int {
	var maxProfit = 0
	var buy = 0
	var sell = 1

	for sell < len(prices) {
		// calculate the difference at the sell and buy points.
		var delta = prices[sell] - prices[buy]

		// if we see a profit over that interval
		if delta > 0 {
			// sell price was higher than buy price.
			// calculate the max so far.
			maxProfit = max(maxProfit, delta)
		} else {
			// sell price was lower than buy. we saw a loss, so, we can move our buy up to sell. 
			// this is minimizing the buy price.
			buy = sell
		}

		// always move the sell pointer to the right
		sell += 1
	}

	return maxProfit
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
		var result = buySellStock(input)
		
		// a bit more work here for Go
		var passed = result == expected
		fmt.Printf("Input: %v\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, result, expected, passed)
	}

}