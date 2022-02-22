package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California":  123,
		"California1": 123,
		"California2": 123,
		"California3": 123,
	}

	m := map[[3]int]string{}

	fmt.Println(statePopulations)
	fmt.Println(m)
}
