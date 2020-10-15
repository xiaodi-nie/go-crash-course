package main

import "fmt" //print stuff

func main() {
	ids := []int{33, 6, 23, 7, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID : %d\n", i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//range with map
	emails := map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharon": "sharon@gmail.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	//just use key
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
