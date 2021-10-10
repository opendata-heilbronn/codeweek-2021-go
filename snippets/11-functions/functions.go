package main

import "fmt"

func main() {
	fmt.Printf("%d + %d = %d\n", 1, 2, addIntegers(1, 2))
	fmt.Printf("%d + %d = %d\n", 3, 5, addIntegers(3, 5))
	fmt.Printf("%d + %d = %d\n", 2, -2, addIntegers(2, -2))
}

func addIntegers(a, b int) int {
	return a + b
}
