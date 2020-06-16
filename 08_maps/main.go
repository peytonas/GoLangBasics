package main

import "fmt"

func main() {
	//NOTE Declare map
	emails := make(map[string]string)

	//NOTE  Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Jane"] = "jane@gmail.com"
	emails["Joe"] = "joe@gmail.com"

	//NOTE Declare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Jane": "jane@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	//NOTE Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
