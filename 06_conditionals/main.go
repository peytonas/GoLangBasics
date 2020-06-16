package main

import "fmt"

func main() {
	x := 10
	y := 10

	//NOTE If/Else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "yellow"
	//NOTE Else/If
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	//NOTE Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	case "green":
		fmt.Println("Color is green")
	default:
		fmt.Println("Color is not blue, red, or green")
	}
}
