package main

import (
    "fmt"
    "strconv"
    "time"
)
const con = "You can't change me!!!"
var glob = "I am a global function!!!"

func add(a int, b int) int{
    return a + b
}


func add2(a, b, c int) int{
    return a + b + c
}


func void_function(){
    fmt.Println("This is a void function.")
}


func global_func() {
    glob = "I changed the clobal function."
    //con = "ds"
}


func for_loop() {
    var i = 1 // Notice that there are 2 ways to define a variable.
    fmt.Println("---- i ----")

    for i <= 3{
        fmt.Println(i)
        i += 1
    }
    
    fmt.Println("---- j ----")

    for j:=1; j<=4; j++{
        fmt.Println(j)
    }
    
    for {
        fmt.Println("loop")
        break
    }
    
}

func if_else() {
    if 7%2 == 0{
        fmt.Println("7 is even.")
    } else{
        fmt.Println("7 is odd.")
    }

    if num := 9; num < 0{
        fmt.Println("num is negative.")
    } else if num < 10{
        fmt.Println("num has 1 digit.")
    } else {
        fmt.Println("num has multiple digits.")
    }
}


func switch1(i int){

    switch i{
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("Don't know, " + strconv.Itoa(i) + " maybe?")  
    }
}


func switch2() {
    switch time.Now().Weekday() {
        case time.Saturday, time.Sunday:
            fmt.Println("It's weekend.")
        default:
            fmt.Println("Weekday it is.")
    }
}


func array1() {
    var a[5]int
    fmt.Println(a)
    fmt.Println(len(a))

    fmt.Println(a[4])

    b := [5]int{1,2,3,4}
    fmt.Println(b)

}


func main(){
    //fmt.Println(add(3, 5))
    //fmt.Println(add2(3, 4, 5))
    //void_function()
    //fmt.Println(glob)
    //global_func()
    //fmt.Println(glob)
    //for_loop()
    //if_else()
    //switch2()
    array1()


}
