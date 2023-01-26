# Problem Description
From [leetcode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/):

> You are given an array prices where prices[i] is the price of a given stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.  Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.


Example:
prices = [7,1,5,3,6,4]
Output: 5
Explanation: If you buy on day 2 at 1 and sell on day 5 at 6, that's 6-1 = 5 in profit. That's the maximum profit you can make given the input data.

# Hints
Try to figure out a brute force solution first. This helped me kind of visualize what was going on. 

After that, ask yourself:
- can we do this in one pass?
- when we do encounter a negative difference, what does this tell us?

# Approach
So, let's try to think about what we need to do here.  We want to figure out if there are any intervals where the sell price is higher than the buy price. When we find such an interval, we need to calculate the difference, which will be a positive number (i.e. profit not loss).  One way to do this would be to check all the pairs in the array:

![Best Time to Buy Sell Stock - Brute Force Approach](https://drive.google.com/uc?export=view&id=1Ib7j2TEd5H9VemG64L84JAZkRHR3Dahw)

We should be able to translate this to code with 2 nested loops:
```python
def buy_sell_stock(input:List[int])->int:
    max_profit = 0
    for buy in range(len(input)):
        for sell in range(buy + 1, len(input)):
            # we're looking for a positive difference, which signals profit.
            difference = input[sell] - input[buy]
            if difference > 0:
                # and we want to maximimize this number
                max_profit = max(max_profit, difference)
    
    return max_profit
```
Since this is 2 nested loops, this would have O(n^2) run time; we didn't use any extra space. One thing to note from the brute force solution, as well as the illustration, is that we don't care about negative profits (i.e. the losses).  Another thing to note is that the two loops affectively gave us two pointers, but, we just didn't move them around in a way that was a bit more clever than using them to check all the pairs. 

Let's try this again, but, take those two observations into consideration:

![Best Time to Buy Sell Stock - Two Pointers](https://drive.google.com/uc?export=view&id=11TPrPqB2CTis3PH0OPTaRZYjNtyZ6-gJ)

We can see that we need to leave the buy pointer alone as long as we're seeing positive differences (i.e. a profit) over an interval/window; as soon as we see a loss over some interval, we can slide buy up to sell, and then move sell to the right and try again.  The loss signifies to us a new low or minimum going forward (since a negative difference means buy is less than sell). That allows us to keep our buy pointer 'sticky' - it only moves when we encounter a new lower price than what we bought it for (i.e. a new minumum from the current buy to sell points).  By moving the buy pointer only when we see a loss, and continuing to move the sell pointer to the right, we are essentially:
- finding the lowest low
- maximimizing the difference between the lowest low with the highest sell point we encounter

Here's another illustration:
![Best Time to Buy Sell Stock - Pointer Movement](https://drive.google.com/uc?export=view&id=1kJIVdjbTtfxV5DHcbh1YRY87Y30bB5-I)

We can see in the first figure, that the buy pointer gets 'stuck' on the second day's value of 1, which is also the smallest going forward. Even though the sell price fluctuates, we we can record the maximum difference as we move through.  So, two key pieces:

```python
difference = prices[sell] - prices[buy]
if difference > 0:
    # here we see a profit (i.e. prices[sell] > prices[buy]) record the max difference
else:
    # a loss, i.e. prices[sell] < prices[buy], move buy up to sell.
    buy = sell
```

We can now do this in one pass, which reduces our run time to linear time (O(n)), and without any extra space. 

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 best_time_to_buy_sell_stock.py
```

For Go:
```shell
go run best_time_to_buy_sell_stock.go
```



