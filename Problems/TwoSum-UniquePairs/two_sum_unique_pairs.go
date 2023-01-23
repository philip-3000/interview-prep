package main

import "fmt"

// A set like structure of integers.
type IntSet map[int]struct{}

// Represents a tuple of 2 elements
type Pair struct {
    x, y int
}

// Returns the number of unique pairs of integers in 
// nums that add up to target.
func twoSumUniquePairs(nums []int, target int) int {

	// let's keep track of the numbers we've seen already.  We'll need to look these up.
	var compliments = IntSet{}

	// and the unique solutions
	seen := make(map[Pair]struct{})
	
	// go through each number in the list...
	for _, value := range nums {
		// for each value, we need to calculate compliment + value = target => compliment = target - value.
		// we then need to see if we've seen this value before
		var compliment = target - value

		if _, ok := compliments[compliment]; ok {
			// great, we found the other number! 

			// always want to store the smaller number first
			var pair = Pair{compliment, value}
			if value < compliment {
				pair.x = value
				pair.y = compliment
			} 
			
			// add the pair into our seen solutions.
			seen[pair] = struct{}{}
		}
        
		// always need to add the number, value, that we just visited.
		compliments[value] = struct{}{}

    }

	return len(seen)
} 

func main() {
	
	// 2 Solutions:
	// 		2 + 3
	// 		1 + 4
	arr := []int{1, 2, 3, 4, 5}
	target := 5
	uniquePairs := twoSumUniquePairs(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nUnique Pairs: %d\n\n", arr, target, uniquePairs)

	// 1 Solution:
	// 		2 + 2
	arr = []int{2, 2, 2, 2, 2}
	target = 4
	uniquePairs = twoSumUniquePairs(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nUnique Pairs: %d\n\n", arr, target, uniquePairs)

	// no solutions: 
    arr = []int{1, 1, 2, 45, 46, 46}
    target = 42
	uniquePairs = twoSumUniquePairs(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nUnique Pairs: %d\n\n", arr, target, uniquePairs)

	// 1 solution:
	//		-2 + 44
	arr = []int{0, 4, 44, -2}
    target = 42
	uniquePairs = twoSumUniquePairs(arr, target)
	fmt.Printf("Input: %v\nTarget: %d\nUnique Pairs: %d\n\n", arr, target, uniquePairs)

}