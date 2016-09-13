/*
Implementation of queue in go language using slices.

`go run queue.go`
*/

package main

import (
        "fmt"
       )


type queue struct{
    list []int
}


func (q queue) is_empty() bool{
    return len(q.list) == 0
}


func (q queue) size() int{
    return len(q.list)
}


func (q *queue) enqueue(value int){
    q.list = append(q.list, value)
}


func (q *queue) dequeue() []int{
    q.list = q.list[1:]
    return q.list
}


func (q queue) print_queue() {
    fmt.Println(q.list)
}


func main() {
    q := queue{}

    v := fmt.Println 
    v(q)
    v(q.size())
    v(q.is_empty())
    q.enqueue(3)
    q.enqueue(4)
    q.enqueue(5)
    q.print_queue()
    q.dequeue()
    q.print_queue()
}
