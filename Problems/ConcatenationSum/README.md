# Problem Description
Given an array of positive integers, a, your task is to calculate the sum of every possible a[i] ∘ a[j], where a[i] ∘ a[j] is the concatenation of the string representations of a[i] and a[j], respectively.

Example:
a = [1,2,3]
concatenation sum = 198
Explanation:
11 + 12 + 13 + 21 + 22 + 23 + 31 + 32 + 33 = 198 

# Hints
Try the intuitive approach first. As you can probably guess, this involves nested loops.

After that, get the pencil and paper out and try re-arranging the sums on paper. Do some basic math first. A pattern will emerge that allows you to remove the nested loops.

# Approach
As with many of the problems done so far, most of the time we can come up with a fairly simple (but slower) solution involving nested loops.  So, from the problem description, it asks us to find all possible concatenations of the elements in the array, and add them all up.  So, let's see what that looks like for the example when we write this out:

![Concat Sum - Example](https://drive.google.com/uc?export=view&id=1oPENxzKmtqI_tQP1F5voqTRDKpl9VKdo)

I wrote out the indices as I think we can see this lends itself to a straight forward nested loop solution:

```python
def concatenation_sum(a:List[int])->int:
    total_sum = 0
    for i in a:
        for j in a:
            concatenated_values = str(i) + str(j)
            total_sum += int(concatenated_values)
    
    return total_sum
```

So, this works, but, like with most of these problems, this solution will be slow and/or will not be fast enough to pass a judge such as leetcode, hackerrank, codesignal, etc. My first train of thought was to see if there was some way I could look values up or some sort of memoization, however, I didn't get very far with that. 

This is where it's necessary to pull the pencil and paper out. What I did was try to rearrange things a bit and play with the numbers until I saw a pattern emerge.  Let's try this with less input:

![Concat Sum - Motivation 1](https://drive.google.com/uc?export=view&id=1Bh4AGIByoLHN6mrsyxXpcxQlLaqbz89g)

and let's go over the numbered steps one by one:
1. first we need to start out by just writing out the concatenation sum
1. notice that we can start pulling out the numbers from our array. For example, 22 is just 20 + 2, 210 is 200 + 10, etc, and 2,10 were numbers in our array.
1. let's group together the values from our array, a, that we extracted. We can see that we have a 2+10, twice.
1. we can reearrange the sums a bit where we take the first number, 20, from the first grouping, and add it with the first number from the second grouping, 100. Then repeat this process with the second set of pairings.
1. After the previous step, it's easier to see how we can factor out a power of 10.

Note how we're left with a few key pieces here:

![Concat Sum - Motivation 1](https://drive.google.com/uc?export=view&id=1SSxzbX9PnIslxslZGr7L9l1vSqkxfi4x)

Where exactly do these numbers from from? Well, let's look at the quantity (b) first. Note that it's twice the sum of just the elements of the array itself. Where does the '2' come from? Well, that happens to be the length of the array from our example. You can see how in the calcuations, we pulled out 2+10, and there was exactly two pairings of this sum. 

The quantity (a) is a bit trickier. We see in there we have the sum of the elements of our array, but then there are powers of 10 multiplying each of those sums. If we look at the first element of our input array, 2, we see it has 1 digit. If we look at the second element, 10, we notice it has 2 digits. Ten to the number of digits corresponds to those powers of 10! 

To show that this isn't just a one off trick, let's look at another example:

![Concat Sum - Motivation 2](https://drive.google.com/uc?export=view&id=1YsfdbOJBD91haOzrSOFbU339HROHRjsV)

You can see we're left with the same parts, (a) and (b). This time, each of our digits had a length of 1, so we just see a factor of 10 multiplying the sum of the elements. 

Now that we have a formula, we can essentially remove the nested loops.  The first quantity we have to calculate is just the sum of the elements (and in which case, we can also calculate part (b)):
```python
# this is equivalent to looping through and adding them all up.
sum_of_all_elements = sum(a)

# let's call our part (b) smaller_sum
smaller_sum = len(a) * sum_of_all_elements 
```
Now we have to run a loop over the elements of the array to calculate part (a), which I'll call 'power sums'
```python
power_sums = 0
for i in a:
    power_of_ten = 10**(len(str(i)))
    power_sums += ( power_of_ten * sum_of_all_elements )
```

And finally, we just add the two together:
```python
total_sum = power_sums + smaller_sum
```

Phew! So, now, we have an algorithm that runs in O(n) time instead of quadratic time. We also didn't use up any more space. 


# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 concat_sum.py
```

For Go:
```shell
go run concat_sum.go
```