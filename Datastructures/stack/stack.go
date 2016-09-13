/*
Implementation of stack in go language using slices.
*/

package main

import (
        "fmt"
        "errors"
       )


type stack struct{
    list []int
}


func (s stack) is_empty() bool{
    return len(s.list) == 0
}


func (s stack) peek() (int, error){
    if !s.is_empty(){
        return s.list[len(s.list)-1], nil
    }
    return s.size(), errors.New("stack is empty")
}


func (s stack)size() int{
    return len(s.list)
}


func (s *stack)push(value int){
    s.list = append(s.list, value)
}


func (s *stack)push(value string){
    s.list = append(s.list, value)
}


func (s *stack)pop() []int{
    return s.list[:len(s.list)-1]
}


func (s stack)print_stack() {
    fmt.Println(s.list)
}


func main() {
    s := stack{}

    v := fmt.Println 
    v(s)
    v(s.size())
    v(s.is_empty())
    v(s.peek())
    s.push(3)
    s.push(4)
    s.push(5)
    s.print_stack()
}
