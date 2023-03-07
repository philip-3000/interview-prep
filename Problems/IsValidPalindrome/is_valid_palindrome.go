/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
)

type TestCase struct {
	S        string
	Expected bool
}

type TestCases struct {
	Tests []TestCase
}

// Helper function to check if a character is alphanumeric.
func isAlphanumeric(character byte) bool {
	var r = rune(character)
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	var left = 0
	var right = len(s) - 1

	// we want to check the character on the left against the character
	// on the right.
	for left < right {
		// as long as the left character is alphanumeric
		for left < right && !isAlphanumeric(s[left]) {
			left += 1
		}

		// and so is the right character.
		for left < right && !isAlphanumeric(s[right]) {
			right -= 1
		}

		// once we've skipped over any alphanumerics, compare them
		// in a lowercase fashion.
		var left_value = unicode.ToLower(rune(s[left]))
		var right_value = unicode.ToLower(rune(s[right]))
		if left_value != right_value {
			return false
		}

		// move the left to the right and the right to the left.
		left += 1
		right -= 1
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
		var expected = test.Expected
		var result = isPalindrome(s)
		var passed = result == expected
		fmt.Printf("s: %v\nResult: %v\nExpected: %v\nPassed: %v\n\n", s, result, expected, passed)
	}

}
