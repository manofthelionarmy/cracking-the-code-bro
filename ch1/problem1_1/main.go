package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	// hasUniqueCharacters := HashUniqueCharacters(input)
	// hasUniqueCharacters := DirectAddress(input)
	hasUniqueCharacters := BitVectorHasUniqueCharacters(input)
	if hasUniqueCharacters {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
