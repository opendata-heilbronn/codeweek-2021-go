package slices

import "fmt"

func slice(){
	//START OMIT
	//nil slice
	var s []int

	//leeres slice
	t := make([]float64, 0)

	//einelementiges slice: [Text]
	u := []string{"Text"}

	//t: [Text, Nachricht]
	u = append(u, "Nachricht")

	v := u[0] //v ist "Text"

	w := u[1] // w ist "Nachricht"
	//END OMIT

	fmt.Println(s,t,u,v,w)
}
