import json

def maximum_subarray_sum(input):
    max_sum = input[0]
    running_sum = 0
    for i in input:
        if running_sum < 0:
            # reset to current element 
            running_sum = i
        else:
            # otherwise, add in the current element
            running_sum += i
        
        # and then check to see if this is the biggest one so far.
        max_sum = max(running_sum, max_sum)
    
    return max_sum


if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            expected = test['expected']
            result = maximum_subarray_sum(input=input)
            passed = expected == result
            print(f"Input:{input}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")
