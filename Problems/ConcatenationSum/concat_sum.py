"""
Given an array of positive integers, a, your task is to calculate the sum of every possible a[i] ∘ a[j], where a[i] ∘ a[j] is the concatenation of the string representations of a[i] and a[j], respectively.

Example:
a = [1,2,3]
concatenation sum = 198
Explanation:
11 + 12 + 13 + 21 + 22 + 23 + 31 + 32 + 33 = 198 
"""
from typing import List
import json

def concatenation_sum(a:List[int])->int:
    total_sum = 0 

    # first, let's collect the sum of all the elements
    sum_of_all_elements = sum(a)

    # our 'smaller sum' is just the sum of all elements multiplied by the
    # length of the array
    smaller_sum = len(a) * sum_of_all_elements

    # now we need to go through and calculate the powers of 10 times the sum of 
    # all elements in a, and accumulate these
    power_sums = 0
    for i in a:
        # calculate 10 to the power of the number of digits in i.
        power_of_ten = 10**(len(str(i)))

        # now multiply that power of 10 by the sum of all elements, and accumulate these 
        # values
        power_sums += (power_of_ten * sum_of_all_elements)


    # add the two together now
    total_sum = smaller_sum + power_sums

    return total_sum


if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            expected = test['expected']
            result = concatenation_sum(a=input)
            passed = expected == result
            print(f"Input:{input}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")






