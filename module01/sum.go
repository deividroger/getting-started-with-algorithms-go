package module01

func Sum(list []int) int {

	if len(list) == 0 {
		return 0
	}
	return list[0] + Sum(list[1:])
}

// func Sum(list []int) int {

// 	var total = 0
// 	for _, value := range list {
// 		total += value
// 	}
// 	return total
// }
