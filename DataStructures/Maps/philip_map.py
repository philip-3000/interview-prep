class PhilipMap(object):
    """
    A dictionary/map like object for educational purposes. 
    """
    def __init__(self, initial_size=1000):
        self.__size__ = initial_size
        self.__bucket__ = [[] for _ in range(initial_size)]
        
        # initialize length to 0
        self.__length__ = 0

    def __getitem__(self, key)->any:
        """
        Returns the value/item for the specified key. Raises a KeyError if the key does not exist in the collection.
        Example: 
            value = m[key]
        """
        # calculate our index/slot to retreive the key/value.
        hash_index = self.__calculate_hash_index__(key)

        # There could be collisions, so, use a loop.
        # if there's no collisions, we'll just use up 1 iteration.
        for (k, v) in self.__bucket__[hash_index]:
            if k == key:
                return v
        
        # if we get here...the key doesn't exist. 
        raise KeyError(f"The key '{key}' doesn't exist in this instance of '{self.__class__.__name__}'")
        
    def __len__(self):
        """
        Returns the number of items currently in the map.
        Example:
            len(m) # returns 5 if there are 5 items.
        """
        return self.__length__
    
    def __delitem__(self, key):
        """
        Removes an entry from the map by key, if it exists. Otherwise, a KeyError is raised.
        Example: 
            del m[42] # removes key value pair specified by key 42.
        """
        # calculate our index/slot so we can look up the kvp in our bucket.
        hash_index = self.__calculate_hash_index__(key)

        # There could be collisions, so, use a loop.
        # if there's no collisions, we'll just use up 1 iteration.
        for idx, (k, v) in enumerate(self.__bucket__[hash_index]):
            if k == key:
                # remove it. hmmm...note that i think this causes an array reshuffling. Wonder if using a linked list for collisions would actually 
                # be better here. We have to search trough them linearly anyway...but removing it is less time consuming.
                self.__bucket__[hash_index].pop(idx)
                self.__length__ -= 1
                return
        
        # if we get here...the key doesn't exist. 
        raise KeyError(f"The key '{key}' doesn't exist in this instance of '{self.__class__.__name__}'")


    def __calculate_hash_index__(self, key)->int:
        """
        Returns a zero based index into the map's bucket for the given key.
        """
        return hash(key) % self.__size__

    def __setitem__(self, key, value):
        """
        Magic method for adding or updating a key-value into the map.
        
        Example:
            map[key] = value # adds or updates the value specified by key.
        """
        # calculate our index/slot to store the key/value.
        # our hash index is our slot into our bucket
        hash_index = self.__calculate_hash_index__(key)
    
        # store key/value as tuple.
        kvp = (key, value)
    
        # First check to see if the key value pair already exists, in which case we want to update it's value.
        # if there's nothing in that particular slot, we won't actually enter this loop.
        for idx, (k,v) in enumerate(self.__bucket__[hash_index]):
            if k == key:
                # ok found the key, we need to udate it i think here.
                self.__bucket__[hash_index][idx] = kvp
                return

        # which means if we make it here...either this is a first time insert, OR a collision.
        # in either case, we treat it the same by appending the kvp into the particular slot for the key.
        self.__bucket__[hash_index].append(kvp)
        self.__length__ += 1
        return       

    def clear(self):
        """
        Removes all entries in the map.
        """
        self.__bucket__ = [[] for _ in range(self.__size__)]
        
        # reset length to 0
        self.__length__ = 0
        return

    def __contains__(self,key):
        """
        Returns true if the specified key is in the collection, otherwise false.
        """
        # calculate our index/slot to lookup the key/value.
        # our hash_index is our slot into our bucket
        hash_index = self.__calculate_hash_index__(key)

        # There could be collisions, so, use a loop.
        # if there's no collisions, we'll just use up 1 iteration.
        for (k, v) in self.__bucket__[hash_index]:
            if k == key:
                return True
        
        # if we get here, we didn't find it.
        return False

    def __iter__(self):
        """
        Yelds the keys present in the map.
        """
        for (k,_) in self.__yield_items__():
            # just yield the key
            yield k

    def __yield_items__(self):
        """
        Helper method to yield key-value pairs.
        """
        # for each slot in our bucket...
        for slot in self.__bucket__:
            # if it's empty..
            if not slot:
                # nothing to yield so try the next one.
                continue
            # otherwise, we need to go through each key value pair. 
            # if there are collisions, there will be more than 1 of these. 
            for (k,v) in slot:
                # just yeild the key, k.
                yield (k,v)

    def items(self):
        """
        Yields keys and values from the collection.
        """    
        return self.__yield_items__()


    def __str__(self):
        """
        Returns the contents of the collection as a string, mirroring the output from __str__() of the standard Python dictionary.
        """
        key_value_pairs = []
        for slot in self.__bucket__:
            for (k,v) in slot:
                # put '' around strings to make it look nicer
                output = ""
                if isinstance(k, str):
                    output += f"'{k}' : "
                else:
                    output += f"{k} : "
                
                if isinstance(v, str):
                    output += f"'{v}'"
                else:
                    output += f"{v}"
                
                key_value_pairs.append(output)
        return "{" + ", ".join(key_value_pairs) + "}"






