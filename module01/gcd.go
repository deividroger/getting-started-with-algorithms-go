package module01

func GCD(a, b int) int {

	for b != 0 {

		a, b = b, a%b
	}

	return a

}

// func GDC(a, b int) int {

// 	for {
// 		if b == 0 {
// 			return a
// 		}

// 		a, b = b, a%b
// 	}

// }

// func GDC(a, b int) int {

// 	if b == 0 {
// 		return a
// 	}

// 	a, b = b, a%b
// 	return GDC(a, b)

// }
