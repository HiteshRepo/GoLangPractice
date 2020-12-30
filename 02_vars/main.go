package main

import "fmt"

func main() {

	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 unint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	// var name string = "Hitesh"
	// var name = "Hitesh"
	name := "Hitesh" // short-hand var declaration - cannot be used outside the function
	var age int = 26
	var ageNew int32 = 26
	var isCool = true
	const isNotCool = false
	// isNotCool = true - error compile time
	size := 1.3 // float64 by default
	email, phone := "hitesh@def.com", "9503955631"

	fmt.Println(name, age, isCool, isNotCool, email, phone)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", ageNew)
	fmt.Printf("%T\n", size)

}
