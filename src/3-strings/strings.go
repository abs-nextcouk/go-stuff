package main

import (
	"fmt"
)

func main() {
	hola := "Que tal estan los maquinas lo primero de todo"
	fmt.Println(hola)
	fmt.Println(len(hola))

	fmt.Printf("hola[0] = %v (typeof %T)\n", hola[0], hola[0])

	fmt.Println(hola[8:13])
	fmt.Println(hola[8:])
	fmt.Println(hola[:13])

	hola = `
	buah learning so so so so
	much my guy
	`
	fmt.Println(hola)

	evenEnded := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			n := i * j
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				evenEnded++
			}
		}
	}
	fmt.Printf("there are %d even ended numbers\n", evenEnded)
}
