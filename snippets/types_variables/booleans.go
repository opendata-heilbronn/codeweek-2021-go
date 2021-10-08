package main

import "fmt"

func bools() {
	//START OMIT
	var a bool

	b := 1 == 1 // --> true
	c := 1 != 1 // --> false
	//END OMIT
	fmt.Println(a, b, c)
}
