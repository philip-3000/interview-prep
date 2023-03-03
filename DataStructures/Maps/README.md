# What is a Hash Table/Map?
A hash map is a data structure that allows us to map one piece of data to another, or, more specifically, keys to values. You might hear them referred to as an associative array or dictionary. Now, you might ask, 'Why not just use an array?'.  Well, let's try an example use case. Suppose you had an array of objects that had a persons name and phone number:
```python
index = [
    {
        'name'   : 'bob',
        'number' : '123-456-0192'
    },
    {
        'name'   : 'alice',
        'number' : '142-453-0196'
    },
    {
        'name'   : 'philip',
        'number' : '983-423-0971'
    }
]
```
If we want to look up Alice's phone number, well, that's not too hard with just a few entries:
```python
for item in index:
    if item['name'] == 'alice':
        return item['number']
``` 
But, suppose there were millions of entries in our phonebook.  Since finding the phone number would require a linear search through our array, this would require a bit more processing time.  Well, this is exactly what a hash table can help us with!  The hash map allows us to look up a value by a particular key. In our case, what if we could just look up a phone number by someone's name? We can restructure our example Python snippet to use a dictionary instead:
```python
index = {
    'bob'    : '123-456-0192',
    'alice'  : '142-453-0196',
    'philip' : '983-423-0971'
}
```
Now, if we wanted to find Alice's phone number, we can 'key' into our data structure with her name:
```python
alices_number = index['alice']
```
In fact, if you've seen some of the practice problems in this repository, you might recall that often we can reduce time complexity by trading it for some space, and one of the things we've definitely used for this is a hash map (see [Two Sum](../Problems/TwoSum) for an example).

But, how does any of this work? Well, that's what I wanted to find out! 

# A brief note on hashing
Before we can make our own simple hash map, we need to talk about hashing. Simply put, a hash function is a function that can transform data of arbitrary size, say a string, into a piece of data of a fixed size.  So for example:
```
h('some string') => 5
```
One thing to note is that a hash function is one way. That is, it's impossible to recover the input from the output.  Ideally, we want this mapping to be be unique; that is, for each input value, we produce a distinct output value.  However, this is not always the case, and we call this a collision.  Let's consider the following hash function that takes integers as input, and produces integers for output:
```python
def simple_hash(input:int)->int:
    return input % 10
```
Our function works great provided we only feed it numbers 0-9.  But, what happens if we wanted to compute the hash value for 11? Well, the output of simple_hash(11) would be 1, which is the same as simple_hash(1), simple_hash(21), etc.  We'll have to figure out how to deal with collisions when we make our own map/dictionary.

And we probably need a better hash function than that. Some hashing functions are really good and have an extremely low chance of a collision, such as [SHA-2](https://en.wikipedia.org/wiki/SHA-2). 

# What's hashing got to do with building a dictionary?
Good question!  One thing we didn't talk about is how we are going to store these key value pairs. A natural way to do this would be to use an array. It acts as our bucket if you will, and as long as we have an index, we have quick look ups. If we're using an array, we need to find a way to map our key to a key value. More specifically, we need to find a way to map our key to an index into our array/bucket.  This is where the hashing function comes in. So, let's say we've got 3 buckets, or, an array of size 3, and we are just interested in hashing 3 key value pairs in:
![No Collisions](https://drive.google.com/uc?export=view&id=1drdCy-X5zrgZWLEhySaoA_UADKZG1Q0q)

We can see here our hash function is well behaved enough, and we have enough storage, so that we can cleanly map each key to a unique bucket/index (i.e. no collisions).  But, let's say we want to start adding more phone numbers in:
![Collisions](https://drive.google.com/uc?export=view&id=139EMzcLhtn8lds6iZujzED3H7TxhA7RO)

We can see here that now "joe" needs to share a bucket with "philip".  This is because we now have more keys/values in our data structure then we have unique slots for.  This is the [pidgeonhole principle](https://math.hmc.edu/funfacts/pigeonhole-principle/#:~:text=The%20pigeonhole%20principle%20is%20one,with%20at%20least%202%20pigeons.) in action.

So, how do we deal with this? Well, we need to make sure that our buckets can hold more than 1 item to account for possible collisions. How do we do this? We still need to map a key to an array index for quick look ups, but, once we have our index into our bucket, now the internal structure of the bucket may differ.  We could, for example, use nested arrays:
```python
index = [
    [ ('bob', '123-456-0192')],
    [ ('alice', '142-453-0196')],
    [ ('philip', '983-423-0971'), ('joe', '847-123-4567')]
]
```
So, now 'joe' and 'philip' can share a bucket/list. Another thing we could do is use an array of linked list nodes:
![Buckets of Linked Lists](https://drive.google.com/uc?export=view&id=1ugQVjzmLbDxk9LU3rE3nFkoibaoadwxt)

For both of these structures, this is actually called ['seperate chaining'](https://www.geeksforgeeks.org/separate-chaining-collision-handling-technique-in-hashing/?ref=lbp).
If we have collisions, now we will have to start scanning through from where our intial key index landed us. This is where things can potentially slow down. There are other methods of dealing with this, in particular, [open addressing](https://www.geeksforgeeks.org/open-addressing-collision-handling-technique-in-hashing/). In open addressing, we ensure we have enough room in our list/bucket such that we're always storing entries in the table itself. For example, if we have a collision, that means we will need to find another available slot to store the key/value in. This also means that re-sizing the structure is something to take in consideration; as our bucket fills up, we need to get a bigger bucket. There are also various recipes for how to actually look up a key value pair if there are collisions on keys: linear scanning, quadratic scanning, etc.  From what I can gather, open addressing is more sophisticated and will yield better performance. Seperate chaining is simpler, but, will come at a cost in performance. 

# Building our own map/dictionary
With that said, let's try building our own, non-production ready dictionary for educational purposes only.  I'll use seperate chaining since it's simpler, and it'll be my own personal dictionary. Some operations it should probably support:
- inserting a key-value pair
- look up a value from a given key, returining it if it exists.
- iteration over keys
- iteration over keys and values
- removing a key-value pair from our collection
- containment (i.e. does a given key exist in our collection)
- be able to print itself out in a string representation (e.g. when you write print(your_dictionary))

That's not an exhaustive list, but, probably enough to get us going with something functional. Let's do this in Python first.

## In Python
Let's try this in Python first as I have a feeling it'll be a little simpler.  If you're familiar with Python, then you've almost certainly used a dictionary before. So, we know it can do all of the operations I listed above:
```python
# insert 
phone_book['philip'] = '983-423-0971'

# get if exists - otherwise it raises a KeyError
philips_number = phone_book['philip']

# iteration over keys
for name in phone_book:
    print(name)

# iteration over keys and values
for (k,v) in phone_book.items():
    print(f"({k}, {v})")

# removal of key-value pair
del phone_book['philip']

# containment
if 'philip' in phone_book:
    # call philip!

#...etc
```

So, I'd like to make my own map be Pythonic enough such that a Python programmer would easily understand how to use it. You might be wondering how can we make our own map utilize the "in" operator, or the brackets for keying into our collection. That is, how do we actually define those? If you're a C# or C++ programmer, you've probably come across operator overloading. Well, for Python, we need to implement [dunder methods](https://mathspp.com/blog/pydonts/dunder-methods#:~:text=What%20are%20dunder%20methods%3F,__%20or%20__add__%20.) (double underscore).  Some people refer to these as 'magic methods' since they seem to be magically invoked (i.e. not invoked by the progammer).  That is because they are impliclty called by Python for certain operations.  That is, when we write:
```python
if 'philip' in phone_book:
```
what actually gets invoked is the dunder method "\__contains__".  [Here's](https://mathspp.com/blog/pydonts/dunder-methods#list-of-dunder-methods-and-their-interactions) a nice list of the dunder methods and what they're utilized for.  With that in mind, we've got various dunder methods to implement. 

Let's start with the dunder method we need to intiailize our collection, "\__init__":
```python
class PhilipMap(object):
    """
    A dictionary/map like object for educational purposes. 
    """
    def __init__(self, initial_size=1000):
        # __size__ indicates the capacity of our bucket.
        self.__size__ = initial_size

        # __bucket__ is where we'll store our keys and values
        self.__bucket__ = [[] for _ in range(initial_size)]
        
        # initialize __length__, which indicates how many items are in our bucket, to 0
        self.__length__ = 0
```
In this, we see that we have a list of lists for our bucket, of a predefined size since we're going to use separate chaining. I made the size configurable as a constructor parameter so that you can actually see the chaining/linear scanning in action when collisions start occurring. You'll see them happen more frequently as you set initial_size to be something smaller, however, I think the chance of a collision is still decent even with setting it to 1000 (see the [birthday paradox](https://betterexplained.com/articles/understanding-the-birthday-paradox/#:~:text=The%20birthday%20paradox%20is%20strange,assumptions%2C%20by%20the%20way)).

Ok, so, we have a bucket, and an initializer.  In order to be able to put stuff in our bucket, well, we need to be able to convert/hash keys into indices into our bucket. There's many different hash functions, but, for our non-production dictionary, we can use the built in hash function in Python:
```python
def __calculate_hash_index__(self, key)->int:
        """
        Returns a zero based index into the map's bucket for the given key.
        """
        return hash(key) % self.__size__
```
The Python [hash()](https://www.geeksforgeeks.org/python-hash-method/) function works differently depending on the type of data that's being hashed. One thing to note is this only works on immutable types. You can try this out for yourself in the terminal and see what comes back. For our hash function, we need to modulo the output of the Python hash() function by the capacity of our bucket so we can get a zero based index into the bucket. So, now that we have a 'good enough' hash function, we can implement another dunder method to actually add stuff into our collection. For this, we'll need to implement "\__setitem__".  Here's what it'll have to do:
- given a key and value, calculate an index for the key
- if the slot at the calculated index has entries, we need to find the appropriate key and update the key-value pair.
- if it's empty, we need to add it in as a new pair and increment the length.

The following implementation of "\__setitem__" will accomplish this:
```python
    def __setitem__(self, key, value):
        """
        Magic method for adding or updating a key-value into the map.
        
        Example:
            map[key] = value # adds or updates the value specified by key.
        """
        # calculate our index/slot to store the key-value as a tuple
        hash_index = self.__calculate_hash_index__(key)
        kvp = (key, value)
    
        # First check to see if the key-value pair already exists, in which case we want to update it's value.
        # if there's nothing in that particular slot, we won't actually enter this loop. If there is, we'll start linearly
        # scanning for the correct key-value pair to update.
        for idx, (k,v) in enumerate(self.__bucket__[hash_index]):
            if k == key:
                # ok found the key, we need to udate it i think here.
                self.__bucket__[hash_index][idx] = kvp
                return

        # if we make it here...either this is a first time insert, OR a collision.
        # in either case, we treat it the same by appending the kvp into the particular slot for the key.
        self.__bucket__[hash_index].append(kvp)
        self.__length__ += 1
        return   
```
Again, we're not using open addressing, resizing our bucket, etc as that's a bit more complicated. At this point, I think we know enough to write the other magic methods to support our standard dictionary uses cases. You can see my full implementation [here](philip_map.py).  You can run the basic suite of unit tests as well:
```shell
python3 tests.py 
..........
----------------------------------------------------------------------
Ran 10 tests in 0.002s

OK
```

## In Go
This is a bit trickier. My first struggle was just trying to figure out how to hash any '[comparable](https://go.dev/blog/comparable)' key type. Go does this for us with the built in map type, however, they don't expose the actual hashing part to us.  It [looks](https://github.com/golang/go/issues/21195) like I'm not the only one that has tried to search for a way to do this.  I found some [workarounds](https://www.dolthub.com/blog/2022-12-19-maphash/), including a library called [hashstructure](https://github.com/mitchellh/hashstructure). I decided I'd try hashstructure and see how things go.

Let's define our basic map type. I'd like to create a structure that can behave somewhat like a map, and has a generic interface to support keys that are comparable, and any value.  We'll need the map itself, as well as a key value pair structure to store a key and value:

```go
// defines a key value pair to store in our map.
type KeyValuePair[T comparable, V any] struct {
	Key   T
	Value V
}

type PhilipMap[T comparable, V any] struct {
	// underlying storage/buckets.
	storage [][]KeyValuePair[T, V]

	// indicates how many items are currently in the map.
	length int
}
```

I utilized the same 'seperate chaining' process as I did for the Python based map, so, the storage field is a list of lists containing key value pairs. While not nearly as sophisticated as the built in map type in Go, it was a bit easier for me to work with to at least get some working code up and running. 

On to the hashing part. We can do a similar thing as we did for the Python based map, in that I delegated the hashing of the key to another component (in this case, hashstructure, for Python, the built in hash function).  Again, we need to calculate a hash value, and transform that to an index into our storage bucket:

```go
func calculateHashIndex[T comparable](key T, size int) int {
	// delegate the actual hash generation to something else...
	var hashValue, _ = hash.Hash(key, nil)
	var six_four_size = uint64(size)

	// and then convert that to an index into our underlying storage/bucket.
	var index = hashValue % six_four_size
	return int(index)
}
```

Go is a bit pickier about types than Python is, so, I had to cast values appropriately.  

In Python, I was able to implement the dunder methods so that you could get standard indexing/iteration/etc with my map, but, in Go, I don't think that's possible. There's no way as far as I can tell to implement some sort of operator overloading like in C++, C#, etc, or, implement dunder methods like in Python. So, we need our own contract!  We need some basic things in our map:
- look up a value by key
- add/update a key value pair
- iterate over items
- remove a key/value by key
- check for containment
- get the number of items in our collection

We can translate these to Go methods such as:
- Get(key T) (bool, V)
- Put(key T, value V)
- Items() <-chan KeyValuePair[T, V]
- Delete(key T)
- Length() int

Some of these are more interesting than others, but, I guess let's start with inserting/updating items first as we won't be able to delete or get anything if there are no items in the collection.

For adding/updating a key value pair in our map, the process is pretty similar to what we did in Python:

- calculate an index into our bucket
- if we find an item in the bucket at the calculated index whose key matches the input key, we update it
- otherwise, we increase the size of our bucket at the particular index by adding in the new key value pair

We can see this below:

```go
func (pm *PhilipMap[T, V]) Put(key T, value V) {
	var index = calculateHashIndex(key, len(pm.storage))

	// we need to insert the key value pair, or, if it exists, update it.
	var currentBucketLength = len(pm.storage[index])
	for idx := 0; idx < currentBucketLength; idx += 1 {
		if pm.storage[index][idx].Key == key {
			// update the value at the appropriate index.
			pm.storage[index][idx].Value = value
			return
		}
	}

	// append to the slice in this case.
	var kvp = KeyValuePair[T, V]{Key: key, Value: value}
	pm.storage[index] = append(pm.storage[index], kvp)

	// bump the length.
	pm.length += 1
}
```

I'm new to Go, but, I believe "Put()" is a struct 'method'.  I say 'method' because I think in Go, our "Put()" function is a function that 'receives' a pointer to a *PhilipMap* struct instance.  I elected to do this because I want to mutate the contents of the struct from within the function; I don't want the struct copied.  Go automatically handles the fact that this is a reference rather than a value; we didn't have to use the & operator.  

Another thing to note that is that the bucket, *storage*, contains slices.  In the above snippet, we use build in "append()" function to dynamically add the new key value pair to our slice. 

The "Get()" and "Delete()" instance methods are very similar, so, I'll skip further explanation of them. In the next section, we'll discuss iteration.

### Iterating over keys and values
**Disclaimer:** this next part is probably overkill, but, I wanted to experiment. In my first crack at this, I discovered that I was leaking goroutines under certain conditions.  In particular, if I broke out of my loop early with an unbuffered channel (or, with a buffered one that was too small), the leak would occur. In any case, I went down this path as it was a good way to start learning about channels, goroutines, goroutine leaks, etc. I'm sure it's not nearly as performant if as simply copying all the keys and values into a slice, however, again, a *PhilipMap* is not for production use; it's intended for my own educational purposes (and possibly others if they ever stumble upon this).

What is a bit different than the Python *PhilipMap* is the way we can provide an iterater through the key value pairs in the collection. Python, and other languages like C#, make use of [generators and *yield*](https://realpython.com/introduction-to-python-generators/) to efficiently allow for iteration over collections/data sets. Go didn't seem to have a built in *yield* construct that would allow for utilization with 'ranging' over a collection. I wanted to support the *range* key word as that is the idiomatic way a programmer would iterate over keys and values in a standard map. So, I decided to experiment with [channels](https://gobyexample.com/channels).  They reminded me a bit of IObservable<T> if you've ever tried [Reactive Extensions in C#](https://github.com/dotnet/reactive) (or [RxJS](https://rxjs.dev/)).  

In Go, you can 'range' over slices, channels, arrays and maps.  I wanted to experiment with channels as it's pretty cool that a construct like this is baked into the language itself:

```go
func (pm *PhilipMap[T, V]) Items() <-chan KeyValuePair[T, V] {
	// create a buffered channel the size of the total number of keys/values
	out := make(chan KeyValuePair[T, V], pm.Length())
	go func() {
		// go through each bucket
		var numberOfBuckets = len(pm.storage)
		for idx := 0; idx < numberOfBuckets; idx += 1 {
			// and publish the key value pairs, if any, in each bucket to the channel.
			var bucketLength = len(pm.storage[idx])
			for bucketIdx := 0; bucketIdx < bucketLength; bucketIdx += 1 {
				var storedKvp = pm.storage[idx][bucketIdx]
				out <- storedKvp
			}
		}
		
		close(out)
	}()

	return out
}
```

The "Items()" function "sends" values to our send only conduit/channel, *out*, in a [goroutine](https://go.dev/tour/concurrency/1). Our channel is buffered up to the current length of the map - meaning the sender can immediately start pushing values out to that channel. In our case, the receiver is range:

```go
for kvp := range pm.Items() {
	// do something with kvp
}
```

The buffering aids in preventing a leak that could occur when a 'sending' goroutine is blocked on a channel that either hasn't been 'read enough' (e.g. it has a buffer size of say 2, and it's filled up. The sending routine will be blocked) or is unbuffered to begin with.  See [here](https://stackoverflow.com/questions/63235829/what-happens-when-you-break-the-for-statement-with-a-range-channel) and [here](https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html) for some more information on that. Uber has a nice [library](https://github.com/uber-go/goleak) for detecting go routine leaks.  If you take the buffering away, you can see the affect it has upon early exit from a range loop:

```shell
go test
--- FAIL: TestItemsIteration (0.45s)
    philipmap_test.go:52: found unexpected goroutines:
        [Goroutine 22 in state chan send, with github.com/philip-3000/Practice/DataStructures/Maps.(*PhilipMap[...]).Items.func1 on top of the stack:
        goroutine 22 [chan send]:
        github.com/philip-3000/Practice/DataStructures/Maps.(*PhilipMap[...]).Items.func1()
                /home/cabox/workspace/interview-prep/DataStructures/Maps/philip_map.go:88 +0xa5
        created by github.com/philip-3000/Practice/DataStructures/Maps.(*PhilipMap[...]).Items
                /home/cabox/workspace/interview-prep/DataStructures/Maps/philip_map.go:80 +0xaa
        ]
FAIL
exit status 1
FAIL    github.com/philip-3000/Practice/DataStructures/Maps     0.448s
```

There are apparently ways to 'cancel' (this sounds similar to C#'s [Task cancellation](https://learn.microsoft.com/en-us/dotnet/standard/parallel-programming/task-cancellation)) or signal completion, however, I felt like I wasn't quite there yet, so, I went with a buffer big enough to hold our items and let the runtime clean up.

You can run the simple suite of unit tests by simply asking go on the shell:

```shell
go test
PASS
ok      github.com/philip-3000/Practice/DataStructures/Maps     0.003s
```

# Closing Thoughts
I'd like to perhaps modify this one day to experiment with open addressing/resizing. Going back to the basics with this exercise is interesting as you learn some things perhaps you forgot, or, just new things in general, especially if you are trying this out in a new language like I did with Go.  You also get a sense of how much thought and care goes into these standard library structures/functions. 
 
Some new tricks I picked up with Go:
- learning how to utilize Generics in Go
- learning how to implement methods that receive a struct
- learning about channels, goroutines and goroutine leaks
- learning how to write unit tests in Go

If you actually made it this far, thanks for reading! 







