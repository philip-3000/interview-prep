from typing import List
import json
import math
"""
Given an array of integers and a target, find 2 numbers such that 
their sum is closest to the target.  

example:
input = [34,23,1,24,75,33,54,8], target = 60

closest we can get here is 58, which is 34 + 24.

example:
input = [3,2, -5, -3], target = 4
closest we can get is 3 + 2 = 5.

"""

def twoSumClosest(input:List[int], target):
    
    # first we must sort the input in ascending order.
    input.sort()
  
    # initialize left ptr to left side of sorted array.
    # right pointer all the way to the right side of sorted array.
    left = 0
    right = len(input) - 1

    # initialize our closest sum to biggest value and our closest sum to -1
    closest = (math.inf, -1)
    
    # keep looking till the pointers cross
    while left < right:
        temp_sum = input[left] + input[right]
        diff = abs(temp_sum - target)
        if diff < closest[0]:
            # if so, store diff, and the value at left and right
            closest = (diff, temp_sum)
        
        # now, we want to get as close as possible, so need to figure out which pointer to adjust.
        # if our sum was smaller than the target,
        if temp_sum < target:
            # try to make it bigger by advancing the left hand
            # pointer to the right
            left += 1
        else:
            # hmm too big, try to make it smaller.
            right -= 1
    
    # return the closest sum
    return closest[1]

if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            target = test['target']
            expected = test['expected']
            result = twoSumClosest(input=input, target=target)
            passed = expected == result
            print(f"Input:{input}\nTarget: {target}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")
