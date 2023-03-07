# Problem 
From [leetcode](https://leetcode.com/problems/valid-palindrome/):
> A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers. Given a string s, return true if it is a palindrome, or false otherwise.

Examples:

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.


Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.


Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

# Hint
The prompt says that there can be non alphanumerical characters in the input.  We might want to sanitize the input first.

# Approach
There are a few ways to approach this problem. The prompt is pretty straight forward in that it gives us a recipe to try.  We can:
- read the input character by character, ignoring any non alphanumeric characters as we iterate.
- store the lower case alphanumeric characters in an array
- reverse the sanitized input and compare it with itself

Let's try this with an example drawn out:

![Is Palindrome - Three Loops](https://drive.google.com/uc?export=view&id=1Aj50VqJq7qs7GWed9qAoCDulKD_AHMB5)

We can see that "Race. Car" is indeed a palindrome once we filter out the non alphanumerics and only compare lowercase characters. We should be able to code this up in Python pretty quick.  Let's call it our version 1:

```python
def is_palindrome_v1(s:str) -> bool:

    # let's strip out all the unecessary stuff, then put it into an array.
    sanitized_input = []
    for character in s:
        if character.isalnum():
            sanitized_input.append(character.lower())
    
    # now we have to reverse it. go through the sanitized input array in reverse
    # and append the values to our new reverse array.
    reverse = []
    for idx in range(len(sanitized_input)-1, -1, -1):
        reverse.append(sanitized_input[idx])
        
    # now one last loop to do a final comparison
    for idx in range(len(sanitized_input)):
        if sanitized_input[idx] != reverse[idx]:
            return False
    
    return True
```

I purposely bypassed a few niceties of Python (list comprehension and built in array reversal) since we won't have them in other languages such as Go. We can see that we basically have 3 loops:
- the first loop cleans up our input
- the second loop reverses 
- the third loop compares character by character

Even though it's 3 loops, this is still a linear algorithm.  We also used linear space as we used 2 arrays to store the sanitized input, and the reverse of the sanitized input, respectively.

So, is there anything we can do to make this better? Well, first let's think of if there's any other way to check the sanitized input array.  That is, can we check that the reverse of *sanitized_input* is equal to *sanitized_input* without creating another array? It turns out we can! We can use two pointers:

![Is Palindrome - Check Reverse In Place](https://drive.google.com/uc?export=view&id=1q1xkda0Z6VhUrl2W5N6fBoeSkrK5UQj1)


As we can see above, we can move one pointer, *left* that starts at the left hand side of the array, towards the right. We can then move another pointer, *right*, to the left hand size of the array.  Before we move the pointers, we check that the value on the left is equal to the value on the right. We can also see there's no reason to check behond where the pointers cross.  That is, if we make it to the middle, we're good! So, let's try to come up with a version 2 that improves upon version 1: 

```python3 
def is_palindrome_v2(s:str) -> bool:

    # let's strip out all the unecessary stuff, then put it into an array.
    sanitized_input = []
    for character in s:
        if character.isalnum():
            sanitized_input.append(character.lower())
        
    # we can check the sanitized input array in place with two pointers.
    left = 0
    right = len(sanitized_input) - 1

    # the left and right don't need to cross.
    while left < right:
        if sanitized_input[left] != sanitized_input[right]:
            return False

        # move left to the right and right to the left.
        left += 1
        right -= 1
    
    return True
```

Now, we've used 1 less array, and one less loop. The overall run time hasn't changed, and neither has the overall space complexity, BUT, we're inching closer to a more optimal solution! Is there a way we can kind of combine the logic of the two loops (i.e. the sanitization loop and the reversal check loop), and do this in place? How do we 'skip over' non alphanumerics? Well, we can just add an inner loop to seek past the garbage characters, on both the left and the right:

```python
def is_palindrome(s:str) -> bool:
    left = 0
    right = len(s) - 1

    while left < right:
        while not s[left].isalnum() and left < right:
            left += 1

        while not s[right].isalnum() and left < right:
            right -= 1

        if s[left].lower() != s[right].lower():
            return False
        
        left += 1
        right -= 1
    
    return True
```

The inner loop to move our left pointer to the right:

```python
while not s[left].isalnum() and left < right:
            left += 1
```

doesn't increase our time complexity since we are just inching the left pointer along, not iterating over the entre array again.  In fact, you could replace the inner loop wth just a check to see if it's not alphanumeric; if it's not, then increment left and continue:

```python
if not s[left].isalnum():
    left += 1
    continue
```
 
The same for the right pointer of course.  

You can see the above is kind of like the various loops we've written, but, meshed together. We've reduced our space complexity to constant space; we don't use any extra space. The overally runtime is still linear time, but, we did remove the extra loops.

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 is_valid_palindrome.py
```

For Go:
```shell
go run is_valid_palindrome.go
```
