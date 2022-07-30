package tailRecursion

func FibRec(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return FibRec(n-1) + FibRec(n-2)
}

func FibTailRec(n int) int {
	result, _, _ := fibGo(n, 0, 1)
	return result
}

func fibGo(n, a, b int) (int, int, int) {
	if n == 0 {
		return a, b, a + b
	}

	if n == 1 {
		return b, b, a + b
	}

	return fibGo(n-1, b, a+b)
}
