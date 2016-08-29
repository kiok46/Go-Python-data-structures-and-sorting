package main

import "fmt"


type person struct{
	name string
	age int
}

func test_struct() {
	fmt.Println(person{"Kuldeep", 21})

	s := person{"Kuldeep", 21}
	fmt.Println(s.name)

	p := s
	fmt.Println(p)
}

func pass_by_value(variable int){
	variable = 0
}

func pass_by_reference(variable *int) {
	*variable = 0
	fmt.Println(variable)
}

func pass_array(variable [4]int) {
    fmt.Println(variable)
}

func main(){
	/*
    var i int = 10

    fmt.Println(i)

    pass_by_value(i)

    fmt.Println(i)

    pass_by_reference(&i)

    fmt.Println(i)


    j := [4]int{1,2,3,4}

    pass_array(j)
    */
    test_struct()

}