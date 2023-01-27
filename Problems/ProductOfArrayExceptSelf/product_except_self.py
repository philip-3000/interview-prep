"""
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Examples:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
"""

from typing import List
import json


def product_except_self(nums:List[int]):
    output = [1] * len(nums)
    left_prefix_products = [1] * len(nums)
    right_prefix_products = [1] * len(nums)
    
    # build our left prefix products array
    for i in range(1, len(nums)):
        left_prefix_product = nums[i-1] * left_prefix_products[i-1]
        left_prefix_products[i] = left_prefix_product

    # build our right_prefix_products prefix products array. Slightly trickier with the loop indices
    for j in range(len(nums) - 2, -1, -1):
        right_prefix_product = nums[j+1] * right_prefix_products[j+1]
        right_prefix_products[j] = right_prefix_product

    # now we just have to iterate through the length of nums
    # and multiply the i'th index of left times the i'th index of right
    # to get the i'th value of our output, 
    # which is all the elements to the left of i multiplied together
    # with all the elements to the right of i, excluding the i'th value.
    for k in range(len(nums)):
        output[k] = left_prefix_products[k] * right_prefix_products[k]
        
    return output

if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            expected = test['expected']
            result = product_except_self(nums=input)
            passed = expected == result
            print(f"Input:{input}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")


        