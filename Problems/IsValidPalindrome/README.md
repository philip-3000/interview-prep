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
- read the input character by character
- remove any non alphanumeric characters as we iterate
- store the lower case alphanumeric characters in an array
- 