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

	// since the input consists of just lower case ascii letters, we know that if the lengths are different,
	// they cannot be anagrams. 
	if len(s) != len(t) {
		return false
	}

	// start building the frequency maps for s and t
	// note that when we index into the string in Go, we'll get the ascii byte value.  
	// i.e. "hello"[1] => 101 for the ascii value of the letter e.
	var frequencyOfCharsInS = map[byte]int{}
	var frequencyOfCharsInT = map[byte]int{}
	for i := 0; i < len(s); i++ {
		var s_i = s[i]
		var t_i = t[i]

		frequencyOfCharsInS[s_i] += 1
		frequencyOfCharsInT[t_i] += 1
	}

	// now go through the frequency maps.
	for characterInS, frequencyInS  := range frequencyOfCharsInS {

		// first check for containment
		frequencyInT, ok := frequencyOfCharsInT[characterInS]
		if !ok {
			return false
		}

		// now check frequency 
		if frequencyInT != frequencyInS {
			return false
		}
	}

	// if all the checks pass, we're good to 'go'...
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
		var passed = result == expected
		fmt.Printf("s: %v\nt: %v\nResult: %v\nExpected: %v\nPassed: %v\n\n", s, t, result, expected, passed)
	}

}