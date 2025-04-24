package main

import (
	"fmt"
)

func main() {
	sliceAkaArray := []string{"uno", "dos", "tres"}

	fmt.Printf("%v (type %T)\n", sliceAkaArray, sliceAkaArray)

	fmt.Println(len(sliceAkaArray))
	fmt.Println(sliceAkaArray[1:])

	fmt.Println("================")

	for i := range sliceAkaArray {
		fmt.Println(i)
	}

	fmt.Println("================")

	for i, name := range sliceAkaArray {
		fmt.Println(i)
		fmt.Println(name)
	}

	fmt.Println("================")

	for _, name := range sliceAkaArray {
		fmt.Println(name)
	}

	sliceAkaArray = append(sliceAkaArray, "memo")
	fmt.Println(sliceAkaArray)

	fmt.Println("================")
	fmt.Println("================")

	distances := map[string]float64{
		"home":         125,
		"work":         18,
		"unemployment": 0.1,
	}

	fmt.Println(len(distances))
	fmt.Println(distances["home"])
	fmt.Println(distances["homesss"]) // returns 0 (default value)

	value, found := distances["unemployment"]
	if !found {
		fmt.Println("unemployment not found")
	} else {
		fmt.Println(value)
	}

	distances["homan"] = 1240.11
	fmt.Println(distances["homan"])

	delete(distances, "work")
	fmt.Println(distances)

	fmt.Println("================")

	for key := range distances {
		fmt.Println(key)
	}

	fmt.Println("================")

	for key, value := range distances {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
