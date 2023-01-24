# Problem Description
From [leetcode](https://leetcode.com/problems/two-sum/description/):

> Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

# Hints
This problem can be solved in various ways. Consider:
- start with the 'slowest' way first. This is the most intuitive.
- is there a way to speed it up? Given an integer x, can you rewrite x + y = target in order to look up the other value y?
- what can we use to do this look up quickly?

# Intuition
This is a pretty common problem that has a several of variations (e.g. 3 Sum, Two Sum Sorted, etc), so, it's a good place to start. Intuitively, we want to move through the array to find two numbers that add up to the target value. Let's try this for the following parameters:

```python
input = [1,2,3,4,5]
target = 8
```

So, what we can do is run two nested loops. Visually, we can see this as we check each pair, where the outter loop starts at the left most index, and the inner loop runs from the left index plus one, and keeps going all the way to the end of the array:

![Two Sum - Brute Force Approach](https://drive.google.com/uc?export=view&id=14zgJ89UjjuyaBpaI9wJ6jTxiQW1SGnCt)

When we find the answer, we can return the indices immediately; the problem constraint says there's just one answer.  That leads us to the following brute force solution:

```python
from typing import List
def two_sum_take_1(nums:List[int], target:int)->List[int]:
    # outter loop from the beginning to the end of the array
    for idx_1 in range(len(nums)):
        # run the inner loop from the first index + 1 to the end of the array
       for idx_2 in range(idx_1 + 1, len(nums)):
           sum = nums[idx_1] + nums[idx_2]
           if sum == target:
               return [idx_1, idx_2]
```
The above algorithm has two nested loops running through the array, so, it would run in O(n^2) runtime. Can we improve on this? What if we could look up the other number? 

![Two Sum - Intuition](https://drive.google.com/uc?export=view&id=1GF-q9-k9ll7BH405APyhri7IVxeWsKZy)

We can look up the other number as we go through the array by storing the numbers in a map.  We can map the value in the array (i.e. the value will be the key) to the index (i.e. the index will be the value).  Then we're able to look up the other index/value! 

I originally thought you needed to do this in two passes, i.e. first populate the mapping with one loop, and then iterate through the array in a second loop.  However, it turns out, we can just do this in one shot.  If we draw this out, we can see how it works:

![Two Sum - Lookups](https://drive.google.com/uc?export=view&id=1aRuaio71HLMVMPg5Z6P7pQbamzGtu8B0)

Notice how at each index, we're attempting to look up the compliment value.  We add in the values that we encounter in each iteration of the loop, after we check to see if we can find the compliment value. With that, we can pretty much arrive at the solutions.