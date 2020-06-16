package main

import "fmt"

func main() {
	x := 10
	y := 10

	//If/Else
	if x <= y {
		//NOTE %d = base10 int's
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "yellow"
	//Else/If
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	//Switch
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
