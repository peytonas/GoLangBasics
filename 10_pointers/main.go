package main

import "fmt"

func main() {

	//NOTE "&a" prints memory address, outputs as "0x________"; that number is where a's value is stored in memory
	a := 5
	b := &a

	fmt.Println(a, b)

	//NOTE a returns int as expected, b returns *int
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//NOTE Use * to read value at address
	fmt.Println(b)
	fmt.Println(*b)

	//NOTE Change value (at a memory address) with pointer
	*b = 10
	fmt.Println(a)
}
