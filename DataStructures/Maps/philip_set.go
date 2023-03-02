package maps

// A set like structure for determining if members/items exist in a set.
// Note that the items in the set must be 'comparable'.
type PhilipSet[T comparable] map[T]bool

// Returns a new instance of a PhilipSet[T]
func NewPhilipSet[T comparable]() PhilipSet[T] {
	ps := make(PhilipSet[T])
	return ps
}

/* Let's add some functions for this new set type.
   What are some functions that might be valuable for us for a set object?
    - set membership (i.e. is t in s)
	- adding values to a set.
	- removing values from a set
	- intersection of 2 sets
	- union of 2 sets

*/

/*
Returns the intersection of set1 and set2. Recall that the intersection of
set1 and set2 is the set of all items that belong in both set1 and set2.
*/
func (set1 PhilipSet[T]) Intersect(set2 PhilipSet[T]) PhilipSet[T] {
	// create a new set to return
	var intersection = NewPhilipSet[T]()

	// go through each element in set1 and check to see if it's in set2.
	for key := range set1 {
		if set2.Has(key) {
			// since it's in set2 as well, then it's in the intersection.
			intersection.Add(key)
		}
	}
	return intersection
}

/*
Returns the union of set1 and set2. Recall that the union of
set1 and set2 is the set of all items that belong in either set1 or set2.
*/
func (set1 PhilipSet[T]) Union(set2 PhilipSet[T]) PhilipSet[T] {
	// create a new set to return
	var union = NewPhilipSet[T]()

	// add in elements of set1
	for key := range set1 {
		union.Add(key)
	}

	// as well as elements of set2.
	for key := range set2 {
		union.Add(key)
	}

	return union
}

// Adds values into the set.
func (ps PhilipSet[T]) Add(values ...T) {
	for _, val := range values {
		ps[val] = true
	}
}

// Deletes values from the set.
func (ps PhilipSet[T]) Delete(values ...T) {
	for _, val := range values {
		// use built in delete function. in our case
		// the value itself is the key.
		delete(ps, val)
	}
}

// Checks for containment of the value in the set.
func (ps PhilipSet[T]) Has(value T) bool {
	_, ok := ps[value]
	return ok
}

// func main() {
// 	// Trying to make a set with generics as it might be more easility understood.
// 	// by philip...
// 	fmt.Println("Testing out PhilipSet!")

// 	var ps = NewPhilipSet[int]()
// 	fmt.Printf("length = %v\n", len(ps))

// 	fmt.Printf("Contains 42: %v\n", ps.Has(42))
// 	fmt.Println("Adding 42...")
// 	ps.Add(42, 43)
// 	fmt.Printf("Contains 42: %v\n", ps.Has(42))
// 	fmt.Printf("Contains 43: %v\n", ps.Has(43))

// 	fmt.Println("Deleting 42...")
// 	ps.Delete(42)
// 	fmt.Printf("Contains 42: %v\n", ps.Has(42))

// 	var ps2 = NewPhilipSet[int]()
// 	ps2.Add(42)
// 	ps2.Add(43)
// 	ps2.Add(45)
// 	ps2.Add(46)
// 	fmt.Println("ps2: ", ps2)

// 	var intersection = ps.Intersect(ps2)
// 	fmt.Println("Intersection: ", intersection)

// 	var union = ps.Union(ps2)
// 	fmt.Println("Union: ", union)

// }
