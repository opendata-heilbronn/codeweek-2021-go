package main

import "fmt"

func main() {
	example := map[string]int{}

	example["apfel"] = len("apfel")
	example["banane"] = len("banane")
	example["clementine"] = len("clementine")
	example["datteln"] = len("datteln")

	fmt.Println(example)

	length := example["apfel"]
	fmt.Printf("Apfel hat %d Buchstaben\n", length)

	length, present := example["birne"]
	if !present {
		fmt.Println("Birne hat keine Länge hinterlegt")
	} else {
		fmt.Printf("Birne hat %d Buchstaben\n", length)
	}
}
