"""
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, 
find two numbers such that they add up to a specific target number. 
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length. 

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.
 
Your solution must use only constant extra space.
"""

from typing import List
import json

def two_sum_two(input:List[int], target:int)->List[int]:
    # the prompt indicates that the input is sorted. We can make use of this fact so we can also
    # do what it says about not using extra space.

    left = 0
    right = len(input) - 1

    while left < right:
        # ok, so, let's see where we are.  we want the number at left and number at right, when added together
        # to equal the target.
        two_sum = input[left] + input[right]

        if two_sum == target:
            # great! we're done
            break
    
        elif two_sum > target:
            # too big...we need to make the sum smaller.  in order to do that, let's make
            # one of the numbers smaller
            right -= 1
        else:
            # too small.  we need to use bigger numbers
            left += 1


    # it mentions 1-indexed.  
    return [left + 1, right + 1]


if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            target = test['target']
            expected = test['expected']
            result = two_sum_two(input=input, target=target)
            passed = expected.sort() == result.sort()
            print(f"Input:{input}\nTarget: {target}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")