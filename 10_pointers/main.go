package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)     //5 0xc0000140b0
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // *int

	// use * to read val from address
	fmt.Println(*b) // 5
	fmt.Println(*&a)

	// change val with pointer
	*b = 10
	fmt.Println(a) // 10
}
