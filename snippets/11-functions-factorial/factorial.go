package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d: %d\n", i, factorial(int64(i)))
	}
}

func factorial(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
