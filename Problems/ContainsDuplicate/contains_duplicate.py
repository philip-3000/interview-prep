import json
from typing import List

def contains_duplicate(input:List[int])->bool:
    duplicates = set()
    for i in input:
        if i in duplicates:
            return True

        duplicates.add(i)

    return False

if __name__ == "__main__":
    with open("test_cases.json", mode='r', encoding="utf-8-sig") as file:
        text = file.read() 
        test_cases = json.loads(text)
        for test in test_cases['tests']:
            input = test['input']
            expected = test['expected']
            result = contains_duplicate(input=input)
            passed = expected == result
            print(f"Input:{input}\nResult: {result}\nExpected: {expected}\nPassed: {passed}\n\n")