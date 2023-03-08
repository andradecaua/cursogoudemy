package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	arry2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(arry2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 1231, 41}
	fmt.Println(slice)

	slice = append(slice, 1030)
	fmt.Println(slice)

	// Arrays Internos

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
