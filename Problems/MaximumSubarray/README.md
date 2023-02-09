# Problem Statement
From [leetcode](https://leetcode.com/problems/maximum-subarray/):
> Given an integer array nums, find the subarray with the largest sum, and return its sum.

Examples:
```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```

```
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
```

```
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
```

# Hints
This problem took a bit of time for me to digest. As usual, try coming up with a brute force solution first.  Then break out the pencil and paper. Do you notice anything about what happens when we keep adding negative numbers? When should we keep them, and when should we not?

# Approach
So, one thing to note is that the problem asked for the *sum*, and not the actual subarray. This will definitely help us a bit in coming up with the most optimal solution. If we just want the sum, we can start considering each pairing.  Let's draw this out with an example:

![Maximum Subarray - Example 1](https://drive.google.com/uc?export=view&id=1wgkFGSFxFhJK6ho1xHcstOFe1_GEn-Ad)

In the above example, we considered each pair, and we can see our maximum sum was 3. Since we considered each pair, this means we probably need two nested loops, and a running sum.  Let's try to write out the algorithm in quadratic time first:

```python
def maximum_subarray_sum(input):
    # consider the first element as being the largest (single) subarray sum.
    max_sum = input[0]
    for i in range(len(input)):
        running_sum = 0
        for j in range(i, len(input)):
            # accumulate the sum of the elements from i to j (inclusive)
            running_sum += input[j]
            
            # and then check to see if this is the biggest one so far.
            if running_sum > max_sum:
                max_sum = running_sum
    
    return max_sum
```
Not too bad/easy enough to understand, but, like most of these problems, can we do better? Well, let's return to our example we wrote out before. Let's consider the subarray where our first loop starts at i = 0, and the inner loop runs from i = 0 to the end of the array:

![Maximum Subarray - Example 2](https://drive.google.com/uc?export=view&id=15Upqv2j6DSLp9TvK0NT8dJQesadmjBtr)

We can see that when we add in the element -4, our sum dips below 0. When we add in the subsequent positive number, we actually made that subsequent number smaller than if it would have just been considered a single subarray on it's own.  In this case, [2]. Let's see another example with all negative numbers:

![Maximum Subarray - Example 2](https://drive.google.com/uc?export=view&id=1kvxKOBfVsO9cKXwwTl_yEbpRQKX2oulz)

We can see here that the more negative numbers we add to the sum, the worse (smaller) it gets.  So...do we discard the negative numbers? Well, not quite:

![Maximum Subarray - Example 3](https://drive.google.com/uc?export=view&id=166PspAEa42eGXe68yGj-e8c2iHNemRdW)

We can see in the above example that if we would have stopped at -1 and picked up at 1000, our subarray sum would have been smaller. That is, the -1 was ok here because our sum was still positive, which means if we add it to another positive number, it's an even larger positive number. 

I think this means that the only time it doesn't make sense to keep our running sum is if it dips below 0! When it does, it only takes away from considering the next element in the array on it's own. With that, we can come up with a linear time solution where we keep a running sum, and reset the running sum to the current element if it has fallen below 0:

```python
if running_sum < 0:
    # reset it to current element
    running_sum = i
```

Otherwise, if the sum is still positive, or zero (in this case, adding 0 doesn't change the running sum) it's worth keeping and adding it in to the current element:

```python
...
else:
    running_sum += i
```
While we go through the array, after we've done the above, we'll also hae to compare the running sum to the 'max so far' so we can keep track of the maximum. 

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 maximum_subarray_sum.py
```

For Go:
```shell
go run maximum_subarray_sum.go
```