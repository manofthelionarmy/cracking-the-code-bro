package main

func HashUniqueCharacters(input string) bool {
	hashTable := make(map[rune]bool) // how to optimize the memory? not using a data structure?
	for _, r := range input {
		if _, ok := hashTable[r]; !ok {
			hashTable[r] = true
		} else {
			return false
		}
	}
	return true
}
