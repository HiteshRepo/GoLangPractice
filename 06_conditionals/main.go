package main

import "fmt"

func main() {
	x, y := 5, 10

	if x < y { // && or ||
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Printf("color is red")
	} else if color == "blue" {
		fmt.Printf("color is blue")
	} else {
		fmt.Printf("color is neither red nor blue")
	}

	// Switch
	switch color {
	case "red":
		fmt.Printf("color is red")
	case "blue":
		fmt.Printf("color is blue")
	default:
		fmt.Printf("color is not blue or red")
	}

}
