"""
Given an array of integers, nums, and an integer, target, find the number of unique pairs in the array
such that their sum is equal to the targetr target. Return the number of pairs.

Example:

Input: nums = [1,2,3,4], target = 6
output: 2

Explanation: 
2 + 4 = 6
"""

from typing import List
import json

def two_sum_unique_pairs(nums:List[int], target:int)->int:
    
    # keep track of the unique solutions.
    unique_solutions = set()

    # collect/hash the unique numbers as iterate through.
    compliments = set()

    for val in nums:
        # we are looking for a compliment + val = target => compliment = target - val
        compliment = target - val
        if compliment in compliments:
            # great, we found it. Now construct the solution pair
            # and insert it into our solutions set
            pair = (val, compliment) if val < compliment else (compliment, val)
            unique_solutions.add(pair)
        
        # always add the number we just visited, val, into the set of compliments
        compliments.add(val)
    
    # you can return the unique solutions themselves rather than the count.
    return len(unique_solutions)



if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            target = test['target']
            expected = test['expected']
            result = two_sum_unique_pairs(nums=input, target=target)
            passed = expected == result
            print(f"Input:{input}\nTarget: {target}\nUnique Pairs: {result}\nExpected: {expected}\nPassed: {passed}\n\n")
    

