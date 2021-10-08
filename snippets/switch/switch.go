package main

import "fmt"

func main() {
	//START OMIT
	i := 1

	switch i{
	case 1:
		fmt.Println("i ist 1")
		// kein break nötig
	case 2:
		fmt.Println("i ist 2")
	}

	fmt.Println("Ende")
	//END OMIT
}
