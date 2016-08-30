package main

import "fmt"


func slice1(){
    var c = make([]int, 3)
    b := make([]string, 3)
    fmt.Println(c)
    fmt.Println(b)

    // used just like an array or list in python.

    b = append(b, "hooli")
    b[0] = "Welcome"
    b[1] = "to"
    b[2] = "the"
    fmt.Println(b)

    fmt.Println(b[1:5])

    t := []int{1, 2, 3}

    fmt.Println(t)

    // Two-Dimensonal slices

    td := make([][]int, 3)

    for i:=0; i<3; i++{
    	il := i+1
    	td[i] = make([]int, il)
    	for j:=0; j< il; j++{
    		td[i][j] = i + j
    	}
    }

    fmt.Println(td)


}


func maps1() {
	// Similar to dictinary in python.

	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println(m)

    // Uses double quotes, not single quotes.
	b := m["one"]
	fmt.Println(b)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}


func range1() {
	nums := []int{1,2,3}
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println("sum:", sum)

	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(dict)

    var s int
	for key, value := range dict{
		fmt.Println(key, value)
		s += value
	}
	fmt.Println(s)

	for i, c := range "google"{
		fmt.Println(i, c)
	}
}


func multiple_return() (int, string){
	return 2, "two"
}


func variadic_func(nums ...int) {
	fmt.Println(nums)

	for _, i := range nums{
		fmt.Println(i)
	}

}


func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}


func fact(n int) int{
	if n == 0{
		return 1
	}
	return n*fact(n-1)
}


func main() {
	//slice1()
	//maps1()
	//range1()
	//a, b := multiple_return()
	//fmt.Println(a, b)
	//variadic_func(1,2,3,4)
    
    //ni := intSeq()
    //fmt.Println(ni())
    //fmt.Println(ni())
    //fmt.Println(ni())

    fmt.Println(fact(20))
}





