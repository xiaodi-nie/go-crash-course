package main

import "fmt" //print stuff

func main() {
	a := 5
	b := &a //b is a pointer to a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // *int

	//use * to read val from address
	fmt.Println(*b)  //5
	fmt.Println(*&a) //5

	*b = 10
	fmt.Println("after assignment: ", a)
}
