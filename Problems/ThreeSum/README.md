# Problem Description
From leetcode:
> Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
>
> Notice that the solution set must not contain duplicate triplets.
>
>Example 1:
>
> Input: nums = [-1,0,1,2,-1,-4]
> Output: [[-1,-1,2],[-1,0,1]]
> Explanation: 
>  nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
>  nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
>  nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
>  The distinct triplets are [-1,0,1] and [-1,-1,2].
>  Notice that the order of the output and the order of the triplets does not matter.
>
> Example 2:
>
> Input: nums = [0,1,1]
> Output: []
> Explanation: The only possible triplet does not sum up to 0.
>
> Example 3:
>
> Input: nums = [0,0,0]
> Output: [[0,0,0]]
> Explanation: The only possible triplet sums up to 0.

# Hint
If you haven't already, I highly recommend checking out [Two Sum](../TwoSum) and [Two Sum 2](../TwoSum2) (especially this one). Two Sum 2 will get you pretty close to a solution - the only thing missing after that is dealing with generating duplicate values.

# Approach
Like in the Two Sum problems, we can start with trying to come up with a brute force solution.  The brute force solution for Two Sum involved finding two numbers that add up to the target value using two loops.  We're now asked to find all triplets that sum up to 0.  Let's try this with one of the examples given in the problem description:

![Three Sum - By Hand](https://drive.google.com/uc?export=view&id=1U2lPu84iRvrH6-ucgu9pMNVPYoyPWT1q)


This problem adds one more twist: it asks for unique triplets. So, from above, we can't use (0,1,-1) because we already found (-1,0,1), which, contains the same numbers.  If we were to do this with 3 loops, we have to also figure out how to keep our triplets unique.  Let's go ahead and just try adding a third loop around our standard two sum brute force solution:

```python
def three_sum_brute_force(input:List[int])->List[List[int]]:
    solutions = []
    for i in range(len(input)):
        for j in range(i+1, len(input)):
            for k in range(j + 1, len(input)):
                three_sum = input[i] + input[j] + input[k]
                if three_sum == 0:
                    solutions.append([input[i], input[j], input[k]])

    
    return solutions
```

Now, let's run it for the input:
```python
input = [-1,0,1,2,-1,-4]
answer = three_sum_brute_force(input=input)
```
It will produce the following result:

```
[[-1, 0, 1], [-1, 2, -1], [0, 1, -1]]
```

Again, the prompt tells us we cannot have duplicates, so, we can see in our output we have a problem as [-1,0,1] and [0,1,-1] are the same answers because they contain the same numbers. 

Ok, so, is there any way we could even salvage the above? I think we could do a few things. For one, what if we could filter out duplicates - could we find a way to filter out the triplets with a hash like data structure? Well, one thing we have to do is you cannot hash lists in Python, so, we'd have to store the results as tuples. Let's try that:

```python
def three_sum_brute_force(input:List[int])->List[List[int]]:
    solutions = set()
    for i in range(len(input)):
        for j in range(i+1, len(input)):
            for k in range(j + 1, len(input)):
                three_sum = input[i] + input[j] + input[k]
                if three_sum == 0:
                    solutions.add((input[i], input[j], input[k]))

    
    return solutions
```

Hmmm...this produces the same output:

```shell
{(0, 1, -1), (-1, 0, 1), (-1, 2, -1)}
```

That is because the tuples (0,1,-1) and (-1,0,1) are technically not equal to each other. How could we get them to be the same? We could sort them first! 

```python
def three_sum_brute_force(input:List[int])->List[List[int]]:
    candidates = set()
    for i in range(len(input)):
        for j in range(i+1, len(input)):
            for k in range(j + 1, len(input)):
                three_sum = input[i] + input[j] + input[k]
                if three_sum == 0:
                    candidates.add(tuple(sorted((input[i], input[j], input[k]))))

    # need a list of list back
    solutions = []
    for candidate in candidates:
        solutions.append([candidate[0], candidate[1], candidate[2]])

    return solutions
```

Now this produces the output:

```shell
[[-1, 0, 1], [-1, -1, 2]]
```

which is what we expect.  I also massaged the set back into a list of lists as that is what the function is expected to return.  Ok, so, this works, but, it's slow for sure! Those three nested loops are a problem. In the hint, I suggested that the Two Sum 2 problem could get us partly the way there; can we extend our Two Sum 2 solution?

The key two Two Sum 2 was that the input was sorted. So, what if we were to sort the input first? Could we then perform Two Sum 2 for each element in the input? Our recipe for this would look like the following:

 - sort the array first
 - for each element n the array:
    - perform two sum 2 on the rest of the array, where the target value is 0.
    - when we find a solution, add it to our results
 - return the results. 
 
Ok, so, let's put an outer loop to index through the array, and then lift our Two Sum 2 solution, and adapt it to sum three numbers instead of just two:

```python
def three_sum_take1(input:List[int])->List[List[int]]:
    input.sort()
    solutions = []
    for idx, val in enumerate(input):
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

                # stop and go to the next one like we did in Two Sum 2?
                break

    return solutions
```

Let's try some input:

```python
input = [-1,0,1,2,-1,4]
```

When I run the above code, it produces the following output:

```shell
[[-1, -1, 2], [-1, 0, 1]]
```

Great.  Looks OK so far.  Let's try another:

```python
input = [-2,1,1,0,2]
```

which produces the output:

```shell
[[-2, 0, 2]]
```

Hmm...that doesn't seem right.  If we sort the input:

```python
[-2, 0, 1, 1, 2]
```

We can see that we should get two triplets:

```python
[-2, 0, 2]
[-2, 1, 1]
```

So, it looks like we have a bug or we're not accounting for something in the input data. Let's draw a picture and step through our current algorithm:

![Three Sum - Take 1](https://drive.google.com/uc?export=view&id=1h-eg-LDQibv_4y_bzFGNHbdcaY_kaZqO)


We can see that when we start at the first index, our value is -2, so, we start searching to the left of that value.  Once we find one, we break.  However, there's still another solution present in our subarray to the left of the -2! Unlike some of the other Two Sum problems, this one asks for *all unique* solutions. We need a way to be able to find additional solutions starting from our current value in the outer loop.

Let's try to re-examine the above input data. What if we 'kept going' after we found a solution? That is, keep trying to 'two sum 2' our way to another solution with the given outer index/value:

![Three Sum - Take 2](https://drive.google.com/uc?export=view&id=1KuQCiW-Sagb0BUwPLi99yYLLv9i2GOLA)

We can see if we had kept marching towards the left, we find more solutions. So let's try to adapt our 'take 1' version to 'take 2' where we march our left pointer to the right and keep looking:

```python
def three_sum_take2(input:List[int])->List[List[int]]:
    input.sort()
    solutions = []
    for idx, val in enumerate(input):
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

                # move left towards the other end of the array.
                # we'll consider more pairs or stop when the pointers cross.
                left += 1

    return solutions
```

When we run the 'take 2' version, we do indeed find the second triplet! So, are we there yet? Well, let's try to see if we can come up with another input scenario that's problematic.  Let's consider the following:

```python3
input = [-2,-2,-2,1,1,0,2]
```
and run version 'take 2' again.  We get the following output:

```shell
[[-2, 0, 2], [-2, 1, 1], [-2, 0, 2], [-2, 1, 1], [-2,0,2]]
```

Now we have too many.  Let's see if we can find out what happened here:

![Three Sum - Take 2 Problems](https://drive.google.com/uc?export=view&id=1SBWmfMiPqhcEwKu8BSuU9sYieH3R-ict)

It looks like since we have duplicate -2's, we keep considering the same solution.  That is, we'll start at -2 as our current value, and in the inner loop:
- we'll find one solution [-2, 0, 2].
- We move our left pointer, then our right pointer, and boom, we find another solution [-2, 1, 1].
- We'll then break our inner loop. On the next go around on our main outer loop, we'll move to the next -2.  

This means we'll find the same solutions again. We probably need to skip over duplicate values in our outer loop.  Since we sorted the array, that's pretty easy to do.  We just check that for a given index greater than 0, the value at the current index is not equal to the value at the previous index:

```python
def three_sum_take3(input:List[int])->List[List[int]]:
    input.sort()
    solutions = []
    for idx, val in enumerate(input):
        
        # keep skipping over duplicate values
        if idx > 0 and input[idx-1] == val:
            continue

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

                # move left towards the other end of the array.
                # we'll consider more pairs or stop when the pointers cross.
                left += 1

    return solutions
```

If we run our 'take 3' version, we will have gotten rid of our duplicate solutions and produce the correct output.  Will that do it? We keep finding problematic input, so, let's try to see if we can find another input array that can break our solution. What if we try an array of all 0's? Certainly this has a solution. Let's try this against our 'take 3' version.  Our input:

```python
input = [0, 0, 0, 0]
```

And our 'take 3' outputs:

```python
[ [0, 0, 0], [0, 0, 0] ] 
```

Hmm, so, it found solutions, but, again we have duplicates. We can see what happens visually:

![Three Sum - Take 3 Problems](https://drive.google.com/uc?export=view&id=1Xg3onRsiKzPOVekUo4Gc6L5bHaL6YIvG)

We can see our outer loop skips subsequent 0's, however, our inner loop does not. So now it's the inner loop that is producing duplicate values.  What if we try sliding the left pointer to the right in our inner loop until no more duplicates are found?

```python
def three_sum_take4(input:List[int])->List[List[int]]:
    input.sort()
    solutions = []
    for idx, val in enumerate(input):
        
        # keep skipping over duplicate values
        if idx > 0 and input[idx-1] == val:
            continue

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

                # seek past duplicate values on the left hand side.
                left += 1
                while left < right and input[left - 1] == input[left]:
                    left += 1

    return solutions
```

Now let's run our 'take 4' version.  Bingo! We now get just the one set of 0's:

```python
[[0, 0, 0]] 
```

I think we've finally arrived at a solution that produces the valid/correct output. There's just one more optimization we can make: since the numbers are sorted in ascending order, if we hit a number in our outer loop that is greater than 0, then there's no way adding that number plus two larger numbers could ever add up to 0.  We can just terminate the loop at that point.   

As for runtime, I think we cut our brute force solution down from O(n^3) to O(n^2) with this more optimized solution.  We still have one loop that potentially goes throught the entire array, and then inside of that, another loop that scans in a linear fashion as well. 

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 three_sum.py
```

For Go:
```shell
go run three_sum.go
```
