"""
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
"""

import json


def is_palindrome(s:str) -> bool:
    # loop through from the left and right, comparing the string
    # and the left and right, but, skip over non alphanumeric 
    # characters.
    left = 0
    right = len(s) - 1

    while left < right:
        # so, before we can compare s at left and right, 
        # we need to make sure left and right aren't positioned at 
        # characters that are not alphanumeric. If they are, we need to
        # move them (but only until they don't cross)
        while not s[left].isalnum() and left < right:
            left += 1

        while not s[right].isalnum() and left < right:
            right -= 1

        # now we know we're at valid characters and/or the pointers haven't crossed.
        if s[left].lower() != s[right].lower():
            return False
        
        # increment left, decrement right
        left += 1
        right -= 1
    
    return True


if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            s = test['s']
            expected = test['expected']
            result = is_palindrome(s=s)
            passed = expected == result
            print(f"s: {s}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")