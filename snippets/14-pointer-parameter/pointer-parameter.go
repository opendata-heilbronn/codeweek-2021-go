package main

import "fmt"

func main() {
	myVariable := 1337
	fmt.Printf("myVariable = %d\n", myVariable)

	modifyValue(myVariable)
	fmt.Printf("myVariable = %d\n", myVariable)

	modifyPointer(&myVariable)
	fmt.Printf("myVariable = %d\n", myVariable)
}

func modifyValue(i int) {
	i = 42
}

func modifyPointer(i *int) {
	*i = 42
}
