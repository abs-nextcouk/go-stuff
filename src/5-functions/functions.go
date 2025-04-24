package main

import (
	"fmt"
	"math"
)

func main() {
	val := prueba(100, 4, "hola")
	fmt.Println(val)

	val1, val2 := prueba2(1, 4)
	fmt.Println(val1)
	fmt.Println(val2)

	// parameters are pass-by-value
	values := []int{1, 2, 33, 4, 68}
	doubleAt(values, 2)
	fmt.Println(values)

	value := 4
	double(value)
	fmt.Println(value)

	// pointers!
	doublePtr(&value)
	fmt.Println(value)

	//errors
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("negative!!! this is an error lol (%f)", n)
	}
	return math.Sqrt(n), nil
}

func prueba(x int, y int, s string) int {
	if s[0] == 'a' {
		return x + y
	} else {
		return x - y
	}
}

func prueba2(x int, y int) (int, int) {
	return x + y, x - y
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(i int) {
	i *= 2
}

func doublePtr(i *int) {
	*i *= 2
}
