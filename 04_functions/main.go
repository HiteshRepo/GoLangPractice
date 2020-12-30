package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int { // (num1, num2 int)
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Hitesh"))
	fmt.Println(getSum(1, 2))
}
