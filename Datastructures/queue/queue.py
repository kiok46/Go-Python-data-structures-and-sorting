'''
Implementation of queue using list in python.

Operations are O(1)

`python queue.py`
'''

class Queue():

    def __init__(self):
        self.items = []

    def is_empty(self):
        return self._is_empty()

    def enqueue(self, item):
        self._enqueue(item)

    def dequeue(self):
        return self._dequeue()

    def size(self):
        return self._size()

    def print_queue(self):
        print self.items

    # private methods.

    def _is_empty(self):
        return self.items == []

    def _enqueue(self, item):
        self.items.insert(0, item)

    def _dequeue(self):
        if not self.is_empty():
            self.items = self.items[1:]
            return self.items
        raise Exception('queue is already empty')

    def _size(self):
        return len(self.items)

if __name__ == "__main__":
    queue = Queue()
    queue.enqueue(2)
    print(queue.size())
    queue.enqueue(4) 
    queue.enqueue(5)
    queue.enqueue(9)
    queue.print_queue()
    queue.dequeue()
    queue.dequeue()
    queue.print_queue()
