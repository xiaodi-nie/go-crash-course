package main

import "fmt"

func main() {
	x, y := 10, 10

	if x <= y {
		fmt.Printf("%d is less then or equal to %d\n", x, y)
	} else {
		fmt.Println("%d is less than %d\n", y, x)
	}

	color := "green"

	if color == "red" {
		fmt.Println("color is red\n")
	} else if color == "blue" {
		fmt.Println("color is blue\n")
	} else {
		fmt.Println("color is not blue or red\n")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}

}
