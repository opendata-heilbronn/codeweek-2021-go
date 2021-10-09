package main

import "fmt"

func main() {
	i := 42
	fmt.Printf("i=%v\n", i)
	iPtr := &i
	fmt.Printf("i=%v, iPtr=%v\n", i, *iPtr)
	*iPtr = 15
	fmt.Printf("i=%v, iPtr=%v\n", i, *iPtr)
}
