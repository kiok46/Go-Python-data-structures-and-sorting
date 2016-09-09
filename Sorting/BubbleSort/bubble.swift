/*
Bubble sort implementation in swift.

Bubble sort is not a good algorithm for large data sets as it's
complexity is O(n^2) in both average and worst case.

use this command to run: 

`swift -sdk $(xcrun --show-sdk-path --sdk macosx) bubble.swift`
*/

func bubble_sort(array: [Int]) -> [Int]{
    var array = array
    var count_new: Int = 0
    var count_old: Int = 0
    //var i: Int
    let length: Int = array.count-2

    while true{
        count_old = count_new
        for i in 0...length{
            if array[i] > array[i+1]{
                count_new += 1
                swap(&array[i], &array[i+1])
            }
        }
        if count_old == count_new{
            return array
        }
    }
}


var array = [2, 1, 5, 8, 4]
print (bubble_sort(array))