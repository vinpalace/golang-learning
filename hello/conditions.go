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

	// Initializer; Variable, Its only scoped to this bracket
	if pop, ok := statePopulations["California2"]; ok {
		fmt.Println(pop)
	}

}
