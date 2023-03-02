/* 
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

Note that the inputs are restricted to lower case english letters. 
*/
package main

import (
	"fmt"
	"encoding/json"
    "io/ioutil"
	"log"
)

type TestCase struct {
	S string
	T string
	Expected bool
}
type TestCases struct {
	Tests []TestCase
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var frequency_s = map[byte]int{}
	var frequency_t = map[byte]int{}

	for i := 0; i < len(s); i++ {
		var s_i = s[i]
		var t_i = t[i]

		frequency_s[s_i] += 1
		frequency_t[t_i] += 1
	}

	// no go through the frequency maps.
	for characterInS, frequencyInS  := range frequency_s {

		// first check for containment
		frequencyInT, ok := frequency_t[characterInS]
		if !ok {
			return false
		}

		// now check frequency 
		if frequencyInT != frequencyInS {
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
		var s = test.S
		var t = test.T
		var expected = test.Expected
		var result = isAnagram(s,t)
		
		// a bit more work here for Go
		var passed = result == expected
		fmt.Printf("s: %v\nt: %v\nResult: %v\nExpected: %v\nPassed: %v\n\n", s, t, result, expected, passed)
	}

}