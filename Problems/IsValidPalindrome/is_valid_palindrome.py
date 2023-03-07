"""
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
"""


# will need to add readme...for now I am trying the solution first.

def is_palindrome_v1(s:str) -> bool:

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

    

    # # ok now that we have the sanitized input, we need to check if the reverse of sanitized input is just sanitized input.
    # # we can also just do this with 2 pointers. one at the left, and one at the right
    # left = 0
    # right = len(sanitized_input) - 1

    # while left < right:
    #     if sanitized_input[left] != sanitized_input[right]:
    #         return False
        
    #     # increment left, decrement right
    #     left += 1
    #     right -= 1

    # # if we get here it was a palindrome.
    # return True

def is_palindrome(s:str) -> bool:
    # we can make a micro optimization here.  it won't affect the overall run time really, but,
    # instead of pre sanitizing into a list, we can do ths in one pass with two pointers.
    # the key point is we just have to seek past the non alphanumerical characters.
    # we can do this with two inner loops or continue statements.
    left = 0
    right = len(s) - 1

    while left < right:
        # so, before we can compare s at left and right, we need to make sure left and right aren't positioned at crap.
        while not s[left].isalnum() and left < right:
            left += 1

        while not s[right].isalnum() and left < right:
            right -= 1

        # now we know we're at valid characters and/or the pointers haven't crossed.
        if s[left].lower() != s[right].lower():
            return False
        
        # increment left, decrement right
        left += 1
        right -= 1
    
    return True


if __name__ == "__main__":
    input = "A man, a plan, a canal: Panama"
    result = is_palindrome_v1(s=input)
    print(f"Input: {input}\nIs Palindrome: {result}\n")

    # this is vacuously a palindrome
    input = "   "
    result = is_palindrome_v1(s=input)
    print(f"Input: {input}\nIs Palindrome: {result}\n")

    input =  "race a car"
    result = is_palindrome_v1(s=input)
    print(f"Input: {input}\nIs Palindrome: {result}\n")

    input = "A man, a plan, a canal: Panama"
    result = is_palindrome(s=input)
    print(f"Input: {input}\nIs Palindrome (v2): {result}\n")

    # this is vacuously a palindrome
    input = "   "
    result = is_palindrome(s=input)
    print(f"Input: {input}\nIs Palindrome (v2): {result}\n")

    input =  "race a car"
    result = is_palindrome(s=input)
    print(f"Input: {input}\nIs Palindrome (v2): {result}\n")