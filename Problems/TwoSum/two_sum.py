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

    # (2 + 7) => [0, 1] 
    input = [2,7,11,15]
    target = 9
    indices = two_sum(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nIndices: {indices}\n")


    # (3 + 5) => [3, 5] 
    input = [0,1,2,3,4,5]
    target = 8
    indices = two_sum(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nIndices: {indices}\n")


    # (-2 + 6) => [2, 3] 
    input = [0,1,-2,6,4,5]
    target = 4
    indices = two_sum(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nIndices: {indices}\n")
  