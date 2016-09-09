/*
Bubble sort implementation in go.

Bubble sort is not a good algorithm for large data sets as it's
complexity is O(n^2) in both average and worst case.

use this command to run:

`go run bubble.go`
*/

package main
import "fmt"


func bubble_sort(array []int) []int{
    var count_new, count_old int = 0, 0
    
    for {
        count_old = count_new
        for i :=0; i<len(array)-1; i++{
            if array[i] > array[i+1]{
                count_new += 1
                array[i], array[i+1] = array[i+1], array[i]
            }
        }
        if count_old == count_new{
            return array
        }
    }
}


func main() {
    array := []int{2, 4, 1, 6, 8}
    fmt.Println(bubble_sort(array))
}
