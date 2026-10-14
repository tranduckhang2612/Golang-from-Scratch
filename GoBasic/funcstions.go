package gobasic

func Add(a, b int) int {
	return a + b
}

func ApplyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func CreateMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}