"""
You are given an array prices where prices[i] is the price of a given stock on the ith day. 
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example:
prices = [7,1,5,3,6,4]
Output: 5
Explanation: If you buy on day 2 at 1 and sell on day 5 at 6, that's 6-1 = 5 in profit. That's the maximum profit you can make given the input data.
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