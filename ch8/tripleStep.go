package ch8

func tripleStepHelper(n int, mem map[int]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if _, ok := mem[n]; !ok {
		mem[n] = TripleStep(n-1) + TripleStep(n-2) + TripleStep(n-3)
	}
	return mem[n]
}

func TripleStep(n int) int {
	mem := make(map[int]int)
	return tripleStepHelper(n, mem)
}
