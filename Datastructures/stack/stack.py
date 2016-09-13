'''
Implementation of stack using list.

Operations are O(1)
'''

class Stack():

    def __init__(self):
        self.items = []

    def is_empty(self):
        return self._is_empty()

    def push(self, item):
        self._push(item)

    def pop(self):
        return self._pop()

    def peek(self):
        return self._peek()

    def size(self):
        return self._size()

    def print_stack(self):
        print self.items

    # privates

    def _is_empty(self):
        return self.items == []

    def _push(self, item):
        self.items.append(item)

    def _pop(self):
        if not self.is_empty():
            return self.items.pop()
        raise Exception('Stack is already empty')

    def _peek(self):
        return self.items[len(self.items)-1]

    def _size(self):
        return len(self.items)

if __name__ == "__main__":
    stack = Stack()
    stack.push('kuldeep')
    stack.print_stack()
    stack.push(3)
    stack.print_stack()
    stack.pop()
    stack.print_stack()
    stack.push('test')
    stack.peek()
    stack.print_stack()
