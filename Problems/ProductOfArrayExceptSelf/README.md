# Problem Statement
From [leetcode](https://leetcode.com/problems/product-of-array-except-self/):
> Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]. The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer. You must write an algorithm that runs in O(n) time and without using the division operation.

Examples:
```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```
```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
```

# Hints
I found this problem to be particularly tricky. It took me various steps and stages to understand how this works.  Here's what I had to do:
- first thing I had to try was to write down the brute force solution, even though it says 'you must write an algorithm that runs in linear time'.  
- if you can write down the brute force solution, you should then try noticing the indices.  Maybe print them out.
- try writing the solution out by hand - notice how the product at index i is the product of all the numbers to the left of it, and all the numbers to the right of it.

# Approach
Ok, so, first thing we should do is try to hand calculate this.  Let's start with a simple input array:
```
[1,2,3,4]
```
and perform the calculation by hand:
![Product of Array Except Self - Hand Calculation](https://drive.google.com/uc?export=view&id=1QRzmSp9vYLKU9qyGWq_fhF4jh3n_L86g)
So, basically at each index i, we skipped over the value at the i'th index, and multipled all the other values together.  With that said, that sounds simple enough to write down the brute force solution:
```python
def product_except_self(nums:List[int]):
    output = []
    for i in range(len(nums)):
        product = 1
        for j in range(len(nums)):
            if i != j:
                product *= nums[j]
        output.append(product)
    return output
```
That's pretty much what was verbally described above.  We know this is not a fast enough solution, so, we need to start thinking about a different way.  One thing noted in the hints that it might be good to visualize the indices a bit.  We can start simple with some print statements:
```python
def product_except_self(nums:List[int]):
    output = []
    for i in range(len(nums)):
        product = 1
        print(f"product at index {i}: ", end=" ")
        for j in range(len(nums)):
            if i != j:
                print(f"{j}", end=" ")
                product *= nums[j]
        output.append(product)
        print()
    return output
```
Now, let's run this on similar input:
```python
print(product_except_self([1,2,3,4,5]))
```
We should get the following output:
```shell
product at index 0:  1 2 3 4 
product at index 1:  0 2 3 4 
product at index 2:  0 1 3 4 
product at index 3:  0 1 2 4 
product at index 4:  0 1 2 3 
[120, 60, 40, 30, 24]
```
Notice that for each index, 'i', we don't see the index 'i' in the list of indices on the right hand side output.  Let's see if we can try re-writing the output a different way. I want to collect the numbers from the input array to the left of the current index, 'i', and then to the right of the current index 'i':
```python
def product_except_self(nums:List[int]):
    output = []
    for i in range(len(nums)):
        product = 1
        left = []
        right = []
        print(f"product at index {i}: ", end=" ")
        for j in range(len(nums)):
            if j < i:
                left.append(nums[j])
            if j > i:
                right.append(nums[j])
            if i != j:
                product *= nums[j]
        print(f"Left: {left} Right: {right}",  end="")
        output.append(product)
        print()
    return output
```
If we run this, we can see the numbers from the input array in our two separate arrays:
```
product at index 0:  Left: [] Right: [2, 3, 4, 5]
product at index 1:  Left: [1] Right: [3, 4, 5]
product at index 2:  Left: [1, 2] Right: [4, 5]
product at index 3:  Left: [1, 2, 3] Right: [5]
product at index 4:  Left: [1, 2, 3, 4] Right: []
[120, 60, 40, 30, 24]
```
We can take this one step further, and have it print out product of the numbers in left and right, respectively, and then muliptly those two products together:
```python
def product_except_self(nums:List[int]):
    output = []
    for i in range(len(nums)):
        product = 1
        left = []
        right = []
        print(f"Index {i}: ", end=" ")
        for j in range(len(nums)):
            if j < i:
                left.append(nums[j])
            if j > i:
                right.append(nums[j])
            if i != j:
                product *= nums[j]
        print(f"\n\tLeft: {left}\n\tRight: {right}\n\tProduct(left) * Product(right): {math.prod(left) * math.prod(right)}")
        output.append(product)
        print()
    return output
```
and then run this program:
```
Index 0:  
        Left: []
        Right: [2, 3, 4, 5]
        Product(left) * Product(right): 120

Index 1:  
        Left: [1]
        Right: [3, 4, 5]
        Product(left) * Product(right): 60

Index 2:  
        Left: [1, 2]
        Right: [4, 5]
        Product(left) * Product(right): 40

Index 3:  
        Left: [1, 2, 3]
        Right: [5]
        Product(left) * Product(right): 30

Index 4:  
        Left: [1, 2, 3, 4]
        Right: []
        Product(left) * Product(right): 24

[120, 60, 40, 30, 24]
```
We notice something kind of interesting happened.  We can see that at each index, left and right consist of all the numbers in the array, except at the actual index 'i'. And when we multiply them together, we get the same numbers as we do in our final output! We're on to something here. So, like in other problems we've seen, can we exchange space for efficiency here?  That is, is it possible to precalculate those left and right arrays first?  In particular, could we define a 'left' array where at each index, 'i', the values in left are the product of all the numbers in the array nums to the left of index 'i'?  Well, it turns out we can:

![Product of Array Except Self - Left Prefix Sum](https://drive.google.com/uc?export=view&id=15HnkmYUpD8Sj8OkT2Y6z-Y-nab5day2W)

We can do a similar calculation/method for finding a 'right' array, where each element at index 'i' in 'right' is the product of all the numbers in the input array, nums, that are 'to the right' of index i (i.e. greater than i):

![Product of Array Except Self - Right Prefix Sum](https://drive.google.com/uc?export=view&id=10v-QHxDlah9bJ8tq1DB_XwY2MnM9i7jx)

Now, from our code snippet above, we saw that at each index 'i' of our output, we multiplied the product of all the numbers in our 'left' array with the product of all the numbers in our 'right' array.  This is exactly the same as multiplying a value in left with a value in right at a particular index i:

![Product of Array Except Self - Right times Left](https://drive.google.com/uc?export=view&id=1XCGj1Cjh5EvY3B76fIiRroxpirc1-REn)

We've finally arrived at a recipe for doing this without the division operator or nested loops.  We essentially need to:
- use a loop to calculate the left prefix products array
- use another loop to calculate the right prefix products array
- and then a final loop to multiple each element of left with each element of right

Our solution has 3 loops, each running in linear time.  We still consider this a linear time algorith, O(n).  For space, we ended up using 2 temporary arrays, but, still linear space complexity.  

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 product_except_self.py
```

For Go:
```shell
go run product_except_self.go
```





