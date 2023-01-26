# Problem Description
This problem has appeared on various online assessments (see [here](https://leetcode.com/discuss/interview-question/372434)).  

> Given an int array nums and an int target, find how many unique pairs in the array such that their sum is equal to target. Return the number of pairs.

Example: 
Input: nums = [1, 1, 2, 45, 46, 46]
target = 47
Output: 2

Explanation:
1 + 46 = 47
2 + 45 = 47

Input: nums = [2, 2, 2, 2, 2]
target = 4
Output = 1

Explanation: 
While we can find more than 1 pair of numbers that add up to 4, it's always the same 2:
2 + 2 = 4
This is the only unique solution.

# Hints
The key part of this question is about finding the **unique** solutions. Things to consider:

- how do we find the two numbers that add up to target? We may have [seen](../TwoSum/README.md) this before.
- Is there a way we could look up if we've already seen a solution (and preferably look it up quickly)?

# Approach
If you haven't seen the standard [two sum](../TwoSum/README.md) problem yet, definitely check that out first.  The idea here is the same but with a small twist. Note that the problem description asks for the number of **unique** solutions.  Suppose we had the following array of numbers and target value:

```python
input = [2,2,2,2]
target = 4
```

As before with [two sum](../TwoSum/README.md), we can check all pairs:

![Two Sum - Find All Pairs](https://drive.google.com/uc?export=view&id=185GDWG4hpCsQZYPvcfkS7GDVTXFrbn-R)

but notice this time we found repeated solutions, when the problem asks for only the unique ones. So, what can we do?  Well, like the previous [two sum](../TwoSum/README.md), in the final solution we incorporated a hash map to do quick look ups.  What if we could look up our solutions?

![Two Sum - Look Up Pairs](https://drive.google.com/uc?export=view&id=16Mes5tXo8gH506evh-kyKbdLnV-pxE-Y)

If we store our solutions in a datastructure such as a map or set, (python conveniently has a set), we'd be able to look up if the solution already exists. There's only one catch.  Consider the following solutions:

```python
1 + 45 = 46
45 + 1 = 46
```

Which ones do we take? Either one is ok, but, how we store them in the hash map matters.  We could establish a convention of always storing the smaller number first:
```python
(1,45)
```
That way, when we hash the tuple into our set/map structure, we won't consider the pair (45,1).  We can now modify our original brute force TwoSum solution as follows:

```python
from typing import List
def two_sum_take_1(nums:List[int], target:int)->List[int]:
    unique_solutions = set()
    # outter loop from the beginning to the end of the array
    for idx_1 in range(len(nums)):
        # run the inner loop from the first index + 1 to the end of the array
       for idx_2 in range(idx_1 + 1, len(nums)):
           x = nums[idx_1]
           y = nums[idx_2]
           sum = x + y
           if sum == target:
               # form the tuple - but always put the smaller value first.
               # this way we will avoid counting (1,2) and (2,1) twice.
               pair = (x,y) if x < y else (y,x)
               unique_solutions.add(pair)
    
    return len(seen)
```

And, just as before, we can improve upon the nested loops. Just like in the standard [two sum](../TwoSum/README.md), we were able to incorporate a map to look up the 'compliment' value.  This time, we don't even need the index of the compliment value - we just need to know if the compliment value has been seen or exists. 

With that, we can now adapt our optimized two sum solution by incorporating the changes we made to our brute force solution:
- [python](two_sum_unique_pairs.py)
- [go](two_sum_unique_pairs.go)

Again, like for the regular two sum problem, our optimized solutions traded space for time complexity.  Now, the algorithm runs in O(n) and uses O(n) space.

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in.

From the current directory, for Python:
```shell
python two_sum_unique_pairs.py
```
and for Go:
```shell
go run two_sum_unique_pairs.go
```



