# Problem Description
This problem has appeared in various online [assessments](https://leetcode.com/discuss/interview-question/241808/Google-Two-sum-closest), and is also on [leetcode premium](https://leetcode.com/problems/two-sum-less-than-k/).  There could be several variations of it (e.g. the leetcode premium problem requires that the sum is less than the target, you could be asked for indices, two values, etc).  The overall gist is as follows: 

> Given an array of integers, input, and an integer value, target, find two values in input that sum closest to the target value and return their sum.

Examples: 

```python
Input = [1, 2, 3, 4, 5]
target = 10
Output: 9
```
We can see that the numbers 4 and 5, at indices 3 and 4, respectively, sum to 9, which is as close to 10 as we can get.

Again, there are several variations of this problem. Some ask for the indices, some ask for the two elements, some just ask for the total sum of the two elements closest.  

# Hints
If you haven't already, try the standard [two sum](../TwoSum/README.md) problem first.  This problem is slightly different in that it does not ask you to find the two numbers that sum exactly to the target value, but, rather which 2 get you the closest to it.  Things to consider:
- can we hash map our way to finding the compliment value?
- can we inch closer to our answer instead?

# Intuition
One thing to note is that the problem description doesn't really have any constraints. For example, we could ask:
- can the 'closest sum' be bigger than the target value? Or does it have to be strictly less than the target?
- if so, are we guarunteed it exists?
- can the target value be negative?

Let's pretend we're the interviewer, and let's make some assumptions:
- just try to find the sum closest to the target. it can be above it or below it.
- yes, you can assume the given input always has a solution.
- let's make it more interseting and assume the target value can be any integer (i.e. positive, negative, or even 0)

Great, with that in mind, let's see if we can figure out how to approach this problem. We can try starting where we did with the standard two sum value, which was with comparing all pairs. This time, we not only need to compare the sum of the pairs, but we need to ask how close that sum is to the target value:

![Two Sum - Lookups](https://drive.google.com/uc?export=view&id=1O5PulYSSFYBHxbjp0qPoKK8b8ZQCbcjM)

For the above example, we saw that the two numbers that summed closest to 10 were 4 and 5.  Two things to note from the graphic:
- I only updated the smallest difference when we actually find one less than the current smallest difference. 
- I calculated the distance between the sum and the target with absolute value, even though the example I used doesn't really factor that in. 

With that in mind, we can write down the modified two sum brute force solution:
```python
import math
def twoSumClosest_nested_loops(input:List[int], target)->int:
    closest = (math.inf, -1)
    for i in range(len(input)):
        for j in range(i + 1, len(input)):
            # calculate the sum of the pairs.
            sum = input[i] + input[j]

            # then check the difference. note when calculating the distance between the two numbers, we better use absolute value.
            diff = abs(sum - target)
            if diff < closest[0]:
                # Update closest sum and smallest difference if the difference is smaller than the smallest difference we've seen
                # so far.
                closest = (diff, sum)

    # then return the closest sum.
    return closest[1]
```

That's pretty straight forward - we only made a couple of modifications to our original two sum brute force solution. The downside of course is that this runs in O(n^2) complexity, but doesn't use any extra space.  Is there any way we can speed this up?  

Well, unlike with the other two sum problems, we can't "hash map" our way to looking up an answer. The problem is we're not trying to find an exact value like before, but, rather we need to consider all the pairs, or, enough pairs, as we go through the input array. 

What if we tried considering two pairs at a time?  In our previous example, we had the input array
```python
[1,2,3,4,5]
```
One thing to note is that it's already sorted.  What if we were to use a left and a right pointer? This gives us a pair to consider, namely (input[left], input[right]), and this way, we can control how big our sum gets by moving the left and the right pointers in a particular direction rather looping through all possible pairs. That is, if we want to try making the sum closer, we would consider a smaller number if the sum was bigger than the target; that is, we'd move our right pointer to the left.  If we need to make our sum bigger as it's smaller than the target, we could consider a bigger number; that is, we'd move our left pointer to the right.

Now, our solution is able to 

# Run the Solutions
Feel free to try more test cases by adding them to [test_cases.json](test_cases.json).  

For Python:
```shell
python3 two_sum_closest.py
```

In Go:
```shell
go run two_sum_closest.go
```




