package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string]string)

	// Assign
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)        // map[Bob:bob@gmail.com Sharon:Sh....]
	fmt.Println(emails["Bob"]) // bob@gmail.com
	fmt.Println(len(emails))   // 2

	// Delete from map
	delete(emails, "Bob")

	// Declare map and assign key-vals
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emails2)

}
