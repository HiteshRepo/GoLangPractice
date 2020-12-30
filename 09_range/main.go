package main

import "fmt"

func main() {

	ids := []int{33, 76, 55, 23}

	// loop through ids
	for i, id := range ids { // for _, id := range ids - use _ instead of i if index is not to be used as it will through compile time error
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Range - Map
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	for k, v := range emails2 {
		fmt.Printf("%s : %s\n", k, v)
	}

	for k := range emails2 {
		fmt.Println("Name: " + k)
	}

}
