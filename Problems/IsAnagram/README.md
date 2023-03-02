# Problem Description
From [leetcode](https://leetcode.com/problems/valid-anagram/):
> Given two strings s and t, return true if t is an anagram of s, and false otherwise. An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Example 2: 

Input: s = "rat", t = "car"
Output: false

# Hints
There's a few ways to solve this. A key thing to keep in mind is that if *t* is an anagram of *s*, *t* will have the same letters as *s*, and the same frequency of each letter.

# Approach
One way to start thinking about this problem is if we have two strings *s* and *t*, can we go through each character of *s* and cross it off in *t*, and be left with nothing crossed off in each string? Let's try this with two strings that are anagrams:

![Anagrams - Valid](https://drive.google.com/uc?export=view&id=1340kZlF4Ykym1apjJQp9HDpQxkemoEPK)

Notice how there are no letters left in either string; we can cross them all off.  Now, let's try an example that doesn't work:

![Anagrams - Invalid](https://drive.google.com/uc?export=view&id=1L9X5VMs2fDeleMK0_RA-B-54-qrh_3Wh)

Notice that for *camera* and *cinema*, we cannot cross off all the letters in each string.  We're left with characters in both strings that aren't contained in one or the other.

So, how could we approach this? Well, one task that we need to accomplish that comes to mind is: how could we quickly look up the characters in the strings? We've seen other problems in this repository that make use of a datastructure that lets us do this.  We need a hashtable (i.e. a map or dictionary)!  That would provide for quick look ups.  However, note that just being able to look up the characters is not enough:

![Anagrams - Different Count](https://drive.google.com/uc?export=view&id=1cVwRwi6fPvod1A30Y73qu-oN_EWloGrV)

In the above example, we can see that while each character in "cat" is also in "caat", we have one leftover "a".  Because of this, they cannot be anagrams of each other; we need to be able to look up each character to check for existence and also it's frequency/count.  Let's try this by mapping each character in *s* and *t* to the number of times it occurs in each string, respectively:

![Anagrams - Characters and Count](https://drive.google.com/uc?export=view&id=1FLEGWGZ-XTdzekUOxJZRxYIG02DgQ8Bu)

We can construct an algorithm to do the following:
- insert each character of *s* into a hashset, *frequency_s*. The characters map to the # of times it occurs, respectively.
- do the same for *t*
- for each character in *frequency_s* check that it exists in *frequency_t* and that the counts are the same. If these conditions fail, return false.

Note that the contraints of the problem indicate that the characters are just lower case English characters. We don't have to worry about spaces or or punctuation/do much in the way of input sanitization. Because of this, we can also add a quick check up front:

```python
if len(s) != len(t):
    return False
```

We can combine all of the above with the following pseudo code:

```python
# check length of s and t. return false if they are not equal.

# load up freqeuncy map for s, frequency_s

# load up frequency map for t, frequency_t

# for each character c in frequency_s:
#   if c is not in frequency_t, return false

#   if c does not have the same frequency/count in both maps, return false. 

# if we make it to here, we're good! Return True.
```

For the [Python solution](is_anagram.py), note that we can do this in only a few lines of code with the Counter class.  This construct doesn't exist in Go, so, I elected not to use it so that the two solutions are somewhat similar to each other. 

# Run the Solutions
The test cases are in [test_cases.json](test_cases.json). Feel free to add more test cases in. 

For Python:
```shell
python3 is_anagram.py
```

For Go:
```shell
go run is_anagram.go
```






