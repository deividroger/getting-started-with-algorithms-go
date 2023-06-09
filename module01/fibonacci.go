package module01

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	nMinus2 := 0
	nMinus1 := 1

	var cur int
	for i := 2; i <= n; i++ {
		cur = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 = cur
	}

	return cur
}

// func Fibonacci(n int) int {

// 	if n <= 1 {
// 		return n
// 	}

// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }
