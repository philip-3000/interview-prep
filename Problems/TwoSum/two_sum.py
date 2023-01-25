"""
Given an array of integers nums and an integer target, return indices of the two numbers such that they 
add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Examples: 

Input: nums = [2,7,11,15], target = 9
Output: [0,1]

Explanation: 
nums[0] + nums[1] = 9
"""

from typing import List
import json

def two_sum(nums:List[int], target:int)->List[int]:
    # we can use a map, where the key is the value of the element in the array, which we'll map to the
    # index of the value from the array.
    compliments = {}

    # go through the array 1 element at a time...
    for idx, val in enumerate(nums):
        # we're looking for two numbers, such that val + compliment = target
        # => compliment = target - val.  So, let's see if compliment is in our map.
        compliment = target - val
        if compliment in compliments:
            # problem states there's one answer; return it.
            other_index = compliments[compliment]
            return [other_index, idx]
        
        # always add the value that we just visited in
        compliments[val] = idx

    
if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            target = test['target']
            expected = test['expected']
            result = two_sum(nums=input, target=target)
            passed = expected.sort() == result.sort()
            print(f"Input:{input}\nTarget: {target}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")