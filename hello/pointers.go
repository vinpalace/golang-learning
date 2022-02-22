package main

import (
	"fmt"
)

func main() {
	// var a int = 42
	// var b *int = &a
	// fmt.Println(a, b)
	// a = 27
	// fmt.Println(a, *b)

	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42

	fmt.Println((*ms).foo)

}

type myStruct struct {
	foo int
}
