package main

import "fmt"

func main() {
	q, r, err := divide(8, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d / %d = %d R %d", 8, 3, q, r)
}

func divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("divisor can't be zero")
	}

	remainder := a % b
	quotient := a / b
	return quotient, remainder, nil
}
