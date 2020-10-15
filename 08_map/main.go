package main

import "fmt"

func main() {
	//define map
	//emails := make(map[string]string) //map[key]value
	//
	////assign kv
	//emails["Bob"] = "bob@gmail.com"
	//emails["Sharon"] = "sharon@gmail.com"
	//emails["Mike"] = "mike@gmail.com"

	//declare map and add kv
	emails := map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharon": "sharon@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	//delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
