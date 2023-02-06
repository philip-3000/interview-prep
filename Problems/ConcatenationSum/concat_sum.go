/* 
Given an array of positive integers, a, your task is to calculate the sum of every possible a[i] ∘ a[j], where a[i] ∘ a[j] is the concatenation of the string representations of a[i] and a[j], respectively.

Example:
a = [1,2,3]
concatenation sum = 198
Explanation:
11 + 12 + 13 + 21 + 22 + 23 + 31 + 32 + 33 = 198 
*/
package main

import (
	"fmt"
	"strconv"
	"encoding/json"
    "io/ioutil"
	"log"
)

type TestCase struct {
	Input []uint32
	Expected uint64
}
type TestCases struct {
	Tests []TestCase
}

func sum(array []uint32) uint64 {
	var total uint64 = 0
	for _, val := range array {
		total += uint64(val)
	}
	return total
}

// Helper function to calculate powers of 10.
// Built in math function expected floats; seemed easier
// to just write what I needed.
func powerofTen(exponent uint32) uint64 {
	if exponent == 0 {
		return 1
	}

	var result uint64 = 1
	var counter uint32 = 1
	for counter <= exponent {
		result *= 10
		counter += 1
	}

	return result
}

func concatSum(a []uint32) uint64 {
	var totalSum uint64 = 0

	// first we need to find the sum of all the elements in the array
	var sumOfAllElements = sum(a)

	// now, the lower/smaller sum is just the length of the array, a, muliplied
	// by the length of the array
	var smallerSum = sumOfAllElements * uint64(len(a))
	
	// next, we need to find the power of 10 sums. 
	var powerSums uint64 = 0
	for _, val := range a {
		// need to calculate the power of 10, which is 10 raised to the length of
		// the value in the array, a.
		var exponent = len(strconv.FormatUint(uint64(val), 10))
		var power_of_ten = powerofTen(uint32(exponent))

		// and then multiple that power of ten by the sum of all elements, and aggregate
		// it.
		powerSums += (power_of_ten * sumOfAllElements)

	}

	totalSum = smallerSum + powerSums
	return totalSum
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
		var result = concatSum(input)
		
		// a bit more work here for Go
		var passed = result == expected
		fmt.Printf("Input: %v\nResult: %d\nExpected: %d\nPassed: %v\n\n", input, result, expected, passed)
	}

}