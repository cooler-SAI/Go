package misc

func Sum(a int, b int) int {
	return a + b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsEven(n int) bool {
	return n%2 == 0
}
