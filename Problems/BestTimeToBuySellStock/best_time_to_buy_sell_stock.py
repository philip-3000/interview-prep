"""
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
"""

from typing import List
import json


def buy_sell_stock(prices:List[int])->int:
    max_profit = 0

    buy = 0
    sell = 1
    
    while sell < len(prices):
        # calculate the difference between the sell and buy point.
        delta = prices[sell] - prices[buy]

        # if it's positive
        if delta > 0:
            # that means we made a profit.  
            max_profit = max(max_profit, delta)
        else:
            # a loss...we can discard negative deltas. move buy up to sell.
            buy = sell


        # in either case, we'll move our sell pointer to the right
        sell += 1
    
    return max_profit

if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            expected = test['expected']
            result = buy_sell_stock(prices=input)
            passed = expected == result
            print(f"Input:{input}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")