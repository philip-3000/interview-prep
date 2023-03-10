# Problem Statement
From [leetcode](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/):
> Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length. 
> 
> Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
>
> The tests are generated such that there is exactly one solution. You may not use the same element twice.
> 
> Your solution must use only constant extra space.
>
>Example 1:
>
>Input: numbers = [2,7,11,15], target = 9
>Output: [1,2]
>Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
>Example 2:
>
>Input: numbers = [2,3,4], target = 6
>Output: [1,3]
>Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
>Example 3:
>
>Input: numbers = [-1,0], target = -1
>Output: [1,2]
>Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

# Hint
This is like the standard two sum problem, however, they tell us that the input array is sorted. How can we use this to our advantage? Can the brute force algorithm give us some insight?

# Approach
Like the regular two sum problem, we can try a a brute force method of checking all pairs. For some of the intuition behind this, check out the approach in the regular [two sum](../TwoSum) problem.  Once you understand that, let's look at a simple brute force solution:

```python
def two_sum_two_brute_force(input:List[int], target:int) -> List[int]:
    for i in range(len(input)):
        for j in range(i+1, len(input)):
            two_sum = input[i] + input[j]
            if two_sum == target:
                return [i + 1, j + 1]

    # input says a solution is guarunteed...we shouldn't get here.
    return []
```
There are two differences in this problem compared to the standard two sum:
- it asks for the answer in terms of a 1 indexed array, not 0 indexed.
- the input is sorted.

Since the input is sorted, we probably need to use that to our advantage somehow.  Let's start drawing a picture to see if we can gain any insight here:

![Two Sum 2 - Brute Force Approach](https://drive.google.com/uc?export=view&id=11b4rL5KcFF5_850u-MDCTB4U19-cP2Q_)

In the table, we increment *i* and *j* independently, sum the values at the indices, and check to see if the sum is equal to the target value.  Notice that as we are iterating through the loops, since the input is sorted in ascending order, the values get bigger as we iterate 'to the right'.  In fact, at some point, it looks like they get larger than the target itself. At that point, we don't even need to consider values that sum beyond the target in our second loop.  So, let's go ahead and make that micro optimization (it's still nested loops). I'd also like to rename our indices:

```python
def two_sum_two_brute_force(input:List[int], target:int) -> List[int]:
    for left in range(len(input)):
        for right in range(left+1, len(input)):
            two_sum = input[left] + input[right]
            if two_sum > target:
                break

            if two_sum == target:
                return [left + 1, right + 1]
```

Now our code will break that inner loop when the sum gets too big - there's no point in continuing on in that direction to make the numbers bigger (since the sum gets bigger).  This now gives us one little hint: we have some level of control in the sum.  In our case here, we controlled the sum from getting too big by limiting how far we moved to the right.  With our two loops, however, we don't have too much control in making the sum smaller; that is, we've only stopped it when we've exceeded the sum. How could we make the sum smaller? Consider if our left and right were positioned as such:

![Two Sum 2 - Where to Go?](https://drive.google.com/uc?export=view&id=1_1ek9O5zYckba6bhbKF7pYusDT4FrAzN)

Well, one thing to ask is if have our two indices, how would we move both of them? Our two for loops weren't that smart in that we could only stop one of the loops from exceeding the sum, but we couldn't make the sum smaller. Looking at the above example, our *left* and *right* indices are at a point where the sum is 8.  We could:
- move right to the right to make the sum bigger
- move left to the right to make the sum bigger

Similarly, what if the sum was too big? We could:
- move right to the left to make the sum smaller
- move left to the right to make the sum smaller.

That's a lot of choices though.  What if we only moved one index in one particular direction, and another index in the opposite direction? 

![Two Sum 2 - Left Goes Right, Right Goes Left](https://drive.google.com/uc?export=view&id=1-QSjxKGpTshXF1ZD7YoyLGjiUWNzb-pz)

Now, we can see that if we only move *left* to the right in order to make the sum larger, and *right* to the left to make the sum smaller, we have an easy decision over how to move our indices to control the sum.  

**Important:** By marching the indices inward, we're basically ruling out values on the left or the right of the array based on whether they make the sum too small or too big, respectively.  

In the above drawing, I started the indices from a random point in the array and then made a decision, but, we still have one more question to answer: where to start out with left and right? Well, if we want to rule out numbers too small, we probably want to start *left* at the first index. That way, as we move to the right, we're ruling out the smaller numbers. Similarly, if we want to rule out numbers too big, we should probably start *right* at the very end of the array:

![Two Sum 2 - I think we got it!](https://drive.google.com/uc?export=view&id=1jbdYH61_R6TEVf7Hl27OCRzZOvwOl7E0)

Now we have complete control over our approach! In fact, this approach is called a 'two pointer' approach, and I believe is the first problem involving this particular algorithm in this repo so far. Two pointers can sometimes be used when we have an iterable data structure like an array.  In fact, if you're familiar with binary searching, it also involves a left and right pointer (when searching over an array).  There's no single formula for a two pointer approach.  Sometimes the pointers move independently, as in our case, and sometimes they might move with each other (for example, a slow and fast pointer for linked list traversal). From what I've seen so far, there's a few characteristics:
- two or more pointers involved. They can move dependently or independently.
- a decision on deciding which pointers to move and how
- processing the array in some way when the pointers are moved.

In our case, what we did was:
- use two pointers, a *left* and a *right*. We move them independently.
- if we decide to make the sum bigger, we'd move *left* to the right, and to make the sum smaller, we'd move *right* to the left. 
- our 'processing' would be to take the sum of the values at *left* and *right* and decide if it equals the target. If not, make the decision above.

At this point, we should be able to code up our solution:

```python
def two_sum_two(input:List[int], target:int)->List[int]:
    left = 0
    right = len(input) - 1

    while left < right:
        two_sum = input[left] + input[right]
        if two_sum == target:
            # great! we're done
            break
        elif two_sum > target:
            # too big
            right -= 1
        else:
            # too small.
            left += 1

    # it mentions 1-indexed.  
    return [left + 1, right + 1]
```

The prompt indicates a solution exists, but, we don't want our *left* and *right* indices to cross; we should stop when they meet.

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 two_sum_two.py
```

For Go:
```shell
go run two_sum_two.go
``` 

