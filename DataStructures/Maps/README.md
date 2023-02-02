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
