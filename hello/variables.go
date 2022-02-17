package main

import "fmt"

func main() {

	// var i int
	// i = 42
	// var j int = 27
	// k := 9

	var j float32 = 27
	// fmt.Println(i, j, k)
	fmt.Printf("%v, %T", j, j) // %v means value and %T means type

	var (
		actorName string = "Elizabeth",
		companion string = "Testy",
		doctorName int = 3,
		season int 11
	)

	// NOTE: Walrus operator will take either int or float64 if you want more control use float32 directly
	// Shadowing: Just like python, inner variable takes precendence

	// There are levels of scoping

	// 1. Package scoped, Naming convention is small case
	// 2. Global scoped, Which are also exported from the package, Naming convetion is Large case
	// 3. Block scoped, Just like python

	// Naming convention rules
	// 1.  lower case first letter for package scope
	// 2.  upper case first letter to export
	// 3.  no private


	// Pascal or camelCase and capitalize acronyms
	// Longer names longer life
	// strconv package for strings
	



}
