package maps

import (
	"container/list"
	"fmt"

	hash "github.com/mitchellh/hashstructure"
)

// Trying out the linked list in go. I had left a comment in
// python verson that perhaps an array of linked list would have been better than
// an array of arrays.

// hmmm...super lame...hashing in go is quite a bit ore difficult.  I see that
// when people make a generic hash table, they make you pass in the hashing function
// you want to use to actually make it work...I think i might have to do this
// for now.

// defines a key value pair to store in our map.
type KeyValuePair[T comparable, V any] struct {
	Key   T
	Value any
}

// A map for educatonial purposes only. Currently uses seperate
// chaining for collisions.
type PhilipMap[T comparable, V any] struct {
	//storage [][]KeyValuePair[T, V]

	// need array of linkedlists of key value pair
	storage []list.List

	length int
}

func (pm *PhilipMap[T, V]) Length() int {
	return pm.length
}

/*
Creates a new PhilipMap with a specified size.
*/
func NewPhilipMap[T comparable, V any](size int) PhilipMap[T, V] {
	var pm PhilipMap[T, V]
	pm.storage = make([]list.List, size)
	return pm
}

func calculateHashIndex[T comparable](key T, size int) int {
	var hashValue, _ = hash.Hash(key, nil)
	var six_four_size = uint64(size)
	//fmt.Printf("Hash Value: %v, six_four_size: %v, size: %v", hashValue, six_four_size, size)
	var index = hashValue % six_four_size
	return int(index)
}

/*
Retreives a value for the specified key, if found.
*/
func (pm *PhilipMap[T, V]) Get(key T) (bool, V) {
	var index = calculateHashIndex(key, len(pm.storage))

	//fmt.Printf("(key, calculated index) = (%v, %v)\n", key, index)

	// Go through and see if it exists...if it does we'll update.
	for e := pm.storage[index].Front(); e != nil; e = e.Next() {
		var storedKvp = e.Value.(KeyValuePair[T, V])
		if storedKvp.Key == key {
			//fmt.Printf("Retrievng (key,value) = (%v, %v)\n", key, storedKvp.Value)
			return true, storedKvp.Value.(V)
		}
	}
	var emptyResult V
	return false, emptyResult
}

/*
Adds or updates (if the key exists) the specified key value pair.
*/
func (pm *PhilipMap[T, V]) Put(key T, value V) {
	// hmmm...ok this might work...
	var index = calculateHashIndex(key, len(pm.storage))

	//fmt.Printf("(key, calculated index) = (%v, %v)\n", key, index)

	// we need to insert the key value pair, or, if it exists, update it.
	var kvp = KeyValuePair[T, V]{Key: key, Value: value}

	// Go through and see if it exists...if it does we'll update.
	for e := pm.storage[index].Front(); e != nil; e = e.Next() {
		var storedKvp = e.Value.(KeyValuePair[T, V])
		if storedKvp.Key == key {
			//fmt.Printf("Updating (key,value) = (%v, %v)\n", key, value)
			e.Value = kvp
			return
		}
	}

	// otherwise we need to add it in.
	//fmt.Printf("Adding (key,value) = (%v, %v)\n", key, value)
	pm.storage[index].PushBack(kvp)
	pm.length += 1

}

func main() {

	fmt.Println("hello")

	var pm = NewPhilipMap[int, int](10)

	fmt.Println("pm  = ", pm)

	fmt.Println("pm.s[0]  = ", pm.storage[0])
	fmt.Println("pm.s[0].Len()  = ", pm.storage[0].Len())

	// let's see if we can put something in here.
	var newKvp = KeyValuePair[int, int]{Key: 10, Value: 10}

	fmt.Println("Pushing on {10,10} at index 0")
	pm.storage[0].PushBack(newKvp)
	fmt.Println("pm.s[0]  = ", pm.storage[0])
	fmt.Println("pm.s[0].Len()  = ", pm.storage[0].Len())
}
