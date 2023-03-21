"""
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] 
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation: 
    nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
    nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
    nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
    The distinct triplets are [-1,0,1] and [-1,-1,2].
    Notice that the order of the output and the order of the triplets does not matter.
"""


from typing import List
import json

def three_sum(input:List[int])->List[List[int]]:
    input.sort()
    solutions = []
    for idx, val in enumerate(input):
        
        # if the value is greater than 0, since we already sorted the input, there's
        # no way we can find 3 numbers that sum to 0 since they're all positive. So we can 
        # break at that point.
        if val > 0:
            break

        # keep skipping over duplicate values
        if idx > 0 and input[idx-1] == val:
            continue

        # now perform the two sum 2 scan we've done before.
        left = idx + 1
        right = len(input) - 1
        while left < right:
            three_sum = val + input[left] + input[right]
            if three_sum > 0:
                # hmm, too big make, make it smaller.
                right -= 1
            elif three_sum < 0:
                # hmm too small, make it bigger.
                left += 1
            else:
                # got one!
                solutions.append([val, input[left], input[right]])

                # similar to the outer loop, we need to seek 
                # past duplicate values on the left hand side.
                left += 1
                while input[left - 1] == input[left] and left < right: 
                    left += 1

    return solutions

if __name__ == "__main__": 
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            expected = test['expected']
            result = three_sum(input=input)

            # little bit more work to make sure the answers are correct.
            expected_triplets = set([tuple(sorted(n)) for n in expected])
            actual_results = set([tuple(sorted(n)) for n in result])

            passed = expected_triplets.intersection(actual_results) == expected_triplets
            print(f"Input: {input}\nOutput: {result}\nExpected: {expected}\nPassed: {passed}\n")

