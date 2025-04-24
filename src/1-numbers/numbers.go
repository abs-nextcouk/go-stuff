package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 1
	y = 2

	fmt.Printf("x=%v, typeof %T\n", x, x)
	fmt.Printf("y=%v, typeof %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("mean=%v, typeof %T\n", mean, mean)

}
