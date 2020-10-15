package main

import "fmt"



func main(){
	//MAIN TYPES
	//String
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	//Using var
	//var name string = "Xanthia"
	//type can be omitted
	//var name = "Xanthia"
	var age int32 = 24 //cast into int32
	//var must be used
	const isCool = true

	//Shorthand(cannot appear outside of function)
	name := "Xanthia"
	size := 1.3 //float64

	name, email := "Xanthia", "xiaodn@amazon.com"


	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", email) //type of variable
}
