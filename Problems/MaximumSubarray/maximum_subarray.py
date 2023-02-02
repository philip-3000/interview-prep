a = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
import math

def max_subarray_nested(input):
    max_sum = -math.inf
    for i in range(len(input)):
        current_sum = 0

        # by starting j at i, you consider the i'th element itself
        for j in range(i, len(input)):
            current_sum += input[j]
            #print(f"(j, input[j], current_sum) = ({j}, {input[j]}, {current_sum})")
            
            # expand this out
            if current_sum > max_sum:
                print(f"current_sum ({current_sum}) > max_sum ({max_sum})")
            else:
                print(f"current_sum ({current_sum}) < max_sum ({max_sum})")
            max_sum = max(max_sum, current_sum)
        
        print()
    #print(max_sum)
    return max_sum


all_positive = [1,2,3]
result = max_subarray_nested(all_positive)
print(f"Input: {all_positive}\nresult: {result}\n")

all_negative = [-2,-1,-3]
result = max_subarray_nested(all_negative)
print(f"Input: {all_negative}\nresult: {result}\n")

mixed = [1, -2, 2, 5]
result = max_subarray_nested(mixed)
print(f"Input: {mixed}\nresult: {result}\n")
# ptr1 = 0
# ptr2 = 0

# current_sum = 0
# max_sum = 0

# for idx, i in enumerate(a):
#     if current_sum < 0:
#         print(f"resetting current_sum = {current_sum} at index {idx}")
#         current_sum = 0
#         print("resetting current sum")
    
#     current_sum += i
#     max_sum = max(max_sum, current_sum)

# print(max_sum)






