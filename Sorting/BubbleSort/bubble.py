'''
Bubble sort implementation in python.

Bubble sort is not a good algorithm for large data sets as it's
complexity is O(n^2) in both average and worst case.

use this command to run:

`python bubble.py`
'''


def bubble_sort(array):
    '''
    Expects 1 parameter
        - list containing integers.
    '''
    count_new = 0
    while True:
        count_old = count_new
        for i in range(len(array) - 1):
            if array[i] > array[i+1]:
                count_new += 1
                array[i], array[i+1] = array[i+1], array[i]
        if count_old == count_new:
            return array


if __name__ == "__main__":
    array = [2, 4, 6, 1, 8, 3]
    print bubble_sort(array)
