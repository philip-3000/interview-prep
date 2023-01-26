/* 

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
	Expected bool
	Comment string
}
type TestCases struct {
	Tests []TestCase
}


func containsDuplicate(input []int) bool {
	var duplicates = map[int]int{}

	for _, value := range input { 
		duplicates[value] += 1

		if duplicates[value] > 1 {
			return true
		}
	}

	return false
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
		var result = containsDuplicate(input)
		
		var passed = result == expected
		fmt.Printf("Input: %v\nResult: %v\nExpected: %v\nPassed: %v\n\n", input, result, expected, passed)
	}

}