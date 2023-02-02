import unittest
from philip_map import PhilipMap

class TestPhilipMap(unittest.TestCase):

    def test_setitem(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        phone_book['bob'] = '123-456-0192'
        phone_book['philip'] = '983-423-0971'
        
        self.assertEqual(len(phone_book), 3)

    def test_getitem(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        phone_book['bob'] = '123-456-0192'
        phone_book['philip'] = '983-423-0971'
        self.assertEqual(len(phone_book), 3)

        self.assertEqual(phone_book['alice'], '142-453-0196')
        self.assertEqual(phone_book['bob'], '123-456-0192')
        self.assertEqual(phone_book['philip'], '983-423-0971')

    def test_delete(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        phone_book['bob'] = '123-456-0192'
        phone_book['philip'] = '983-423-0971'
        self.assertEqual(len(phone_book), 3)

        del phone_book['philip']
        self.assertEqual(len(phone_book), 2)

    def test_getitem_raises_key_error_for_nonexistent_key(self):  
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        self.assertEqual(len(phone_book), 1)
        with self.assertRaises(KeyError):
            my_number = phone_book['philip']

    def test_delete_raises_key_error_for_nonexistent_key(self):  
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        self.assertEqual(len(phone_book), 1)
        with self.assertRaises(KeyError):
            del phone_book['philip']

    def test_handles_collisions(self):
        # pidgeonhole principle guarantees us if we put more than 2 items in our collection, one of the slots will definitely have more than 1 item.
        phone_book = PhilipMap(initial_size=2)
        phone_book['alice'] = '142-453-0196'
        phone_book['bob'] = '123-456-0192'
        phone_book['philip'] = '983-423-0971'
        self.assertEqual(len(phone_book), 3)

        self.assertEqual(phone_book['alice'], '142-453-0196')
        self.assertEqual(phone_book['bob'], '123-456-0192')
        self.assertEqual(phone_book['philip'], '983-423-0971')

    def test_setitem_handles_updates(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        self.assertEqual(phone_book['alice'], '142-453-0196')
        self.assertEqual(len(phone_book), 1)
        
        # change value
        phone_book['alice'] = '983-423-0971' 
        self.assertEqual(phone_book['alice'], '983-423-0971')
        self.assertEqual(len(phone_book), 1)

    def test_contains(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        self.assertTrue('alice' in phone_book)
        self.assertFalse('philip' in phone_book)

    def test_iter_keys(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        phone_book['bob'] = '123-456-0192'
        phone_book['philip'] = '983-423-0971'

        # since our collection supports iteration...this should add all the keys in.
        s = set(phone_book)

        # and then we should be able to loop through each key in phonebook and verify it's in our set s.
        for name in phone_book:
            self.assertTrue(name in s)

    def test_items(self):
        phone_book = PhilipMap()
        phone_book['alice'] = '142-453-0196'
        phone_book['bob'] = '123-456-0192'
        phone_book['philip'] = '983-423-0971'

        # check the keys and the values.
        for (name,number) in phone_book.items():
            if name == 'alice':
                self.assertEqual(number, '142-453-0196')
            elif name == 'bob':
                self.assertEqual(number, '123-456-0192')
            elif name == 'philip':
                self.assertEqual(number, '983-423-0971')
            else:
                raise Exception(f"unknow key: {name}")
            


if __name__ == '__main__':
    unittest.main()