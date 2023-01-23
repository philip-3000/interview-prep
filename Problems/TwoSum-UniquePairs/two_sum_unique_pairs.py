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

    # empty array should return no solutions
    input = []
    target = 2
    unique_pairs = two_sum_unique_pairs(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nNumber Of Unique Pairs: {unique_pairs}\n")

    # Only 1 unique solution
    input = [2,2,2,2]
    target = 4
    unique_pairs = two_sum_unique_pairs(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nNumber Of Unique Pairs: {unique_pairs}\n")

    # 2 unique solutions: 
    # 1 + 46  
    # 2 + 45
    input = [1, 1, 2, 45, 46, 46]
    target = 47
    unique_pairs = two_sum_unique_pairs(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nNumber Of Unique Pairs: {unique_pairs}\n")
    

    # no solutions: 
    input = [1, 1, 2, 45, 46, 46]
    target = 42
    unique_pairs = two_sum_unique_pairs(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nNumber Of Unique Pairs: {unique_pairs}\n")
    

    # 1 solution: 
    # -2 + 44 
    input = [0, 4, 44, -2]
    target = 42
    unique_pairs = two_sum_unique_pairs(nums = input, target=target)
    print(f"Input: {input}\nTarget: {target}\nNumber Of Unique Pairs: {unique_pairs}\n")
    

