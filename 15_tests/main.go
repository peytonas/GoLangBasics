package main

import "fmt"

type Room interface {
	room() string
}

type BL struct {
	name, description string
}

func (b BL) room() string {
	return b.name
}

func getRoom(r Room) string {
	return r.room()
}

// var playing := true

func main() {
	bl := BL{name: "BadLands", description: "Desert"}

	fmt.Printf("Room: " + getRoom(bl))
	fmt.Println("\nRoom: " + bl.description)
}
