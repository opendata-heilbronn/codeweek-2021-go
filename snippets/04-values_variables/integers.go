package main

import "fmt"

func integers() {
	//START OMIT
	var a int8
	var b uint8
	var c int16
	var d uint16
	var e int32
	var f uint32
	var g int64
	var h uint64	// 0, 1, 2, 3, 4, ...
	var i int		// ..., -4, -3, -2, -1, 0, 1, 2, 3, 4, ...
	//END OMIT
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
