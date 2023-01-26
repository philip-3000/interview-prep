# Problem Description
From [leetcode](https://leetcode.com/problems/contains-duplicate/):
> Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Examples:
```
input: [1,2,3,4]
output: false
explanation: all the values in the array are unique
```
```
input: [1,2,3,4,3]
output: true
explanation: The value 3 appears twice
```

# Hints
Can we do this with brute force? How about in one pass through the input? Maybe try it both ways!

# Approach
This problem lends itself to a pretty straight forward brute force solution.  We basically have to just consider pairs:

![Contains Duplicate](https://drive.google.com/uc?export=view&id=1R4b2auIWbLqUPxgqEyEWJ6kTMwLsjeZ6)

We can go ahead and write down the simple solution:
```python
def contains_duplicate(input:List[int])->bool:
    for i in range(len(input)):
        for j in range(i + 1, len(input)):
            if input[i] == input[j]:
                return True

    return False
```
Since we have nested loops, this solution would have a run time complexity of O(n^2); we didn't use any extra space. But, is there a way we can make this a bit faster?  Well, in other problems like [two sum](../TwoSum), we ended up increasing our utilization of space in order to gain some run time back.  Maybe we can try a similar approach here. What if we kept a map to keep track of the numbers we've seen so far.  If we encounter the same number again, that means we encountered a duplicate!  You can see this below in the illustration:
![Contains Duplicate - Use Map](https://drive.google.com/uc?export=view&id=1qYUiA2b60tLHF6pOAcx0mj1oC6ucWIil)
Since we only scanned through the loop once from left to right, our run time decreased to linear time, O(n^2), and our space complexity jumped to O(n) since we're storing the elements from the array in our map.

There's probably a few ways you can implement this. For example, in Python, there's a convenient set structure:
```python
for i in input:
    if i in duplicates:
        return True
    
    # otherwise add it in 
    duplicates.add(i)

# if we reached here, no duplicates
```
In Go, there isn't a natural set structure, so, it might be easier to just map the value to the count like we did in the illustration. 

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 contains_duplicate.py
```

For Go:
```shell
go run contains_duplicate.go 
```

