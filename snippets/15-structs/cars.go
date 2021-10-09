package main

import "fmt"

type Car struct {
	Doors        int
	Type         string
	Power        float64
	Color        string
	Registration string
}

func main() {
	myCorsa := Car{
		Doors:        5,
		Type:         "gasoline",
		Power:        80,
		Color:        "black",
		Registration: "HN-PH 404",
	}
	fmt.Printf("my corsa: %#v", myCorsa)
}
