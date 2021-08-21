package main

func DirectAddress(input string) bool {
	charList := make([]int, 256) // num of bytes int * 256

	for i := range charList {
		charList[i] = 0
	}

	for _, r := range input {
		charList[int(r)] += 1
	}

	for i := range charList {
		if charList[i] > 1 {
			return false
		}
	}
	return true
}
