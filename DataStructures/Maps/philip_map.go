package maps

import (
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
	storage [][]KeyValuePair[T, V]

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
	pm.storage = make([][]KeyValuePair[T, V], size)
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

	for _, storedKvp := range pm.storage[index] {
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
		var numberOfBuckets = len(pm.storage)
		for idx := 0; idx < numberOfBuckets; idx += 1 {
			//var currentBucket = &pm.storage_2[idx]
			var bucketLength = len(pm.storage[idx]) // (*currentBucket))
			for bucketIdx := 0; bucketIdx < bucketLength; bucketIdx += 1 {
				var storedKvp = pm.storage[idx][bucketIdx] // &(*currentBucket)[bucketIdx]
				out <- storedKvp
			}
		}
		// for index := range pm.storage {
		// 	for e := pm.storage[index].Front(); e != nil; e = e.Next() {
		// 		var storedKvp = e.Value.(KeyValuePair[T, V])
		// 		out <- storedKvp
		// 	}
		// }

		close(out)
	}()

	return out
}

func (pm *PhilipMap[T, V]) Delete(key T) {
	var index = calculateHashIndex(key, len(pm.storage))
	for idx := 0; idx < len(pm.storage[index]); idx += 1 {
		if pm.storage[index][idx].Key == key {
			// if we find it, we have to manually do this. We don't care about the ordering.
			// kind of annoying there isn't an easier way to do this...
			pm.storage[index][idx] = pm.storage[index][len(pm.storage[index])-1]
			pm.storage[index] = pm.storage[index][:len(pm.storage[index])-1]
			pm.length -= 1
			return
		}
	}
}

/*
Adds or updates (if the key exists) the specified key value pair.
*/
func (pm *PhilipMap[T, V]) Put(key T, value V) {
	var index = calculateHashIndex(key, len(pm.storage))

	// we need to insert the key value pair, or, if it exists, update it.
	var kvp = KeyValuePair[T, V]{Key: key, Value: value}

	// Go makes a copy when you range over...
	var currentBucketLength = len(pm.storage[index])
	for idx := 0; idx < currentBucketLength; idx += 1 {
		if pm.storage[index][idx].Key == key {
			// update the value at the appropriate index.
			pm.storage[index][idx].Value = value
			return
		}
	}

	// append to the slice.
	pm.storage[index] = append(pm.storage[index], kvp)
	//fmt.Printf("Added (k,v) = (%v, %v) to new slot at index %v.\nBucket: %v\n", key, value, index, pm.storage_2)

	// // Go through and see if it exists...if it does we'll update.
	// for e := pm.storage[index].Front(); e != nil; e = e.Next() {
	// 	var storedKvp = e.Value.(KeyValuePair[T, V])
	// 	if storedKvp.Key == key {
	// 		e.Value = kvp
	// 		return
	// 	}
	// }

	// otherwise we need to add it in.
	//pm.storage[index].PushBack(kvp)
	pm.length += 1

}
