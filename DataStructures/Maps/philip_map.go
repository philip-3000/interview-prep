package maps

import (
	"container/list"

	hash "github.com/mitchellh/hashstructure"
)

// defines a key value pair to store in our map.
type KeyValuePair[T comparable, V any] struct {
	Key   T
	Value V
}

/*
A map like object for educational purposes only. Currently uses seperate
chaining for collisions.
*/
type PhilipMap[T comparable, V any] struct {
	// underlying storage/buckets.
	storage []list.List

	// indicates how many items are currently in the map.
	length int
}

/*
Returns the number of items currently in the map.
*/
func (pm *PhilipMap[T, V]) Length() int {
	return pm.length
}

/*
Creates a new PhilipMap with a specified size. The size parameter
specifies the number of buckets alloted.
*/
func NewPhilipMap[T comparable, V any](size int) PhilipMap[T, V] {
	var pm PhilipMap[T, V]
	pm.storage = make([]list.List, size)
	return pm
}

// Calculates an index based on the key and size of the structure.
func calculateHashIndex[T comparable](key T, size int) int {
	// delegate the actual hash generation to something else...
	var hashValue, _ = hash.Hash(key, nil)
	var six_four_size = uint64(size)

	// and then convert that to an index into our underlying storage/bucket.
	var index = hashValue % six_four_size
	return int(index)
}

/*
Retreives a value for the specified key, if found.
*/
func (pm *PhilipMap[T, V]) Get(key T) (bool, V) {
	var index = calculateHashIndex(key, len(pm.storage))

	// Go through and see if it exists...if it does we'll update.
	for e := pm.storage[index].Front(); e != nil; e = e.Next() {
		var storedKvp = e.Value.(KeyValuePair[T, V])
		if storedKvp.Key == key {
			return true, storedKvp.Value
		}
	}
	var emptyResult V
	return false, emptyResult
}

/*
	 Allows for ranging over key value pairs via a channel.

	 Example:
	 for kvp := pm.Items() {
		fmt.Printf("(key, value) = ('%v', '%v')\n", kvp.Key, kvp.Value)
	 }
*/
func (pm *PhilipMap[T, V]) Items() <-chan KeyValuePair[T, V] {
	out := make(chan KeyValuePair[T, V])
	go func() {
		for index := range pm.storage {
			for e := pm.storage[index].Front(); e != nil; e = e.Next() {
				var storedKvp = e.Value.(KeyValuePair[T, V])
				out <- storedKvp
			}
		}

		close(out)
	}()

	return out
}

/*
Adds or updates (if the key exists) the specified key value pair.
*/
func (pm *PhilipMap[T, V]) Put(key T, value V) {
	var index = calculateHashIndex(key, len(pm.storage))

	// we need to insert the key value pair, or, if it exists, update it.
	var kvp = KeyValuePair[T, V]{Key: key, Value: value}

	// Go through and see if it exists...if it does we'll update.
	for e := pm.storage[index].Front(); e != nil; e = e.Next() {
		var storedKvp = e.Value.(KeyValuePair[T, V])
		if storedKvp.Key == key {
			e.Value = kvp
			return
		}
	}

	// otherwise we need to add it in.
	pm.storage[index].PushBack(kvp)
	pm.length += 1

}
