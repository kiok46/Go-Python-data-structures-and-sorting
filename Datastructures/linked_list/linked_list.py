'''
Implementation of linked_list in python.

`python linked_list.py`
'''

class Node(object):

    def __init__(self, data=None, next_node=None):
        self.data = data
        self.next_node = next_node

    def get_data(self):
        return self.data

    def get_next(self):
        return self.next_node

    def set_next(self, new_node):
        self.next_node = new_node


class LinkedList():

    def __init__(self, head=None):
        self.head = head
        self.count = 0

    def insert(self, data):
        new_node = Node(data)
        new_node.set_next(self.head)
        self.head = new_node
        self.count += 1

    def delete(self, data, every=False):
        current = self.head
        previous = None


        while current:
            if current.get_data() == data:
                if every is False:
                    previous = current
                    break
                
            else:
                previous = current
                current = current.get_next()

        if current is None:
            raise ValueError('Data not available in the list.')

        if previous is None:
            self.head = current.get_next()
        else:
            previous.set_next(current.get_next())
            self.count -= 1


    def search(self, data, every=False):
        '''
        if every is set as True then it returns all the data nodes which
        contains the same data value.
        '''
        current = self.head
        if every is True: 
            container = []
        while current:
            if current.get_data() == data:
                if every is True:
                    container.append(current)
                else:
                    return current
            current = current.get_next()
        if (current is None) and (not container):
            raise ValueError("Data not available in the list.")
        return container

    def size(self):
        return self.count

    def print_list(self):
        current = self.head
        while current:
            print current.get_data()
            current = current.get_next()
    

if __name__ == "__main__":

    l = LinkedList()
    l.insert(3)
    l.insert(6)
    l.insert(2)
    l.insert(1)
    l.insert(2)
    l.print_list()
    
    print l.search(2)
    print l.search(2, True)

    l.delete(1, True)
    l.print_list()
