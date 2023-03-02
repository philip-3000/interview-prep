"""
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

Note that the inputs are restricted to lower case english letters. 
"""
import json
from collections import defaultdict

def is_anagram(s:str, t:str)->bool:
    # We don't have to worry about punctuation, spaces, etc. 
    # Given this constraint, if they aren't the same length, they can't be anagrams. 
    if len(s) != len(t):
        return False

    # initialize our counters. 
    # note - in python, we could use the Counter class. I am not using it here since
    # there's no built in analogue in Go.
    frequency_s = defaultdict(int)
    frequency_t = defaultdict(int)

    # since s and t have the same length, we can use 1 loop.
    for i in range(len(s)):
        # for each index i, take the character at i in s and t...
        c_1 = s[i]
        c_2 = t[i]

        # and build our frequency map with them.
        frequency_s[c_1] += 1
        frequency_t[c_2] += 1

    
    # now have to go through and make sure the character frequencies match
    for c in frequency_s:
        # each character of s must be in t
        if c not in frequency_t:
            return False
        
        # and have the same frequency/count.
        if frequency_s[c] != frequency_t[c]:
            return False


    return True


if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            s = test['s']
            t = test['t']
            expected = test['expected']
            result = is_anagram(s=s, t=t)
            passed = expected == result
            print(f"s: {s}\nt: {t}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")
    