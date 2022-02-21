package main

import (
	"fmt"
)

func main() {
	// a := []int{1, 2, 3, 4, 5, 6, 7}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	a := make([]int, 3)
	fmt.Println(a)
	fmt.Printf("Len: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)

}
