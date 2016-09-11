'''
Selection sort implementation in python.

Selection sort is not a good algorithm for large data sets as it's
complexity is O(n^2) in both average and worst case.

use this command to run:

`python selection.py`
'''


def selection_sort(array):
    n = len(array)
    
    for index, i in enumerate(array):
        min_ = i
        temp = i

        for j in array[index+1:]:
            if temp < j:
                temp = j

        print temp
        min_ = temp
        if min_ != i:
            array[min_], array[i] = array[i], array[min_]
    
    return array


if __name__ == "__main__":
    array = [2, 4, 6, 1, 8, 3]
    print selection_sort(array)
