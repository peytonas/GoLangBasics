package main

import "fmt"

func main() {
	// var name string = "Peyton"
	// var age int32 = 30
	// const isCool = true
	// var size = 8.5
	//NOTE shorthand variable declaration:
	name, age, isCool, email, size := "Peyton", 30, true, "peyton.sonnefeld@gmail.com", 8.5

	fmt.Println(name, age, isCool, email, size)
	fmt.Printf("%T\n", size)
}
