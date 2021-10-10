package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	for index, element := range slice {
		fmt.Printf("%d: %d, ", index, element)
	}
	fmt.Print("\n")

	m := map[string]string{
		"hello": "world",
		"go":    "workshop",
	}
	for key, value := range m {
		fmt.Printf("%s: %s, ", key, value)
	}
	fmt.Print("\n")

	for index, character := range "hello world" {
		fmt.Printf("%d: %c, ", index, character)
	}
	fmt.Print("\n")
}
