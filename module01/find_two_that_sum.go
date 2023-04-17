package module01

func FindTwoThatSum(number []int, sum int) (int, int) {

	for i, val := range number {
		for j, val2 := range number {
			if i == j {
				continue
			}
			if val+val2 == sum {
				return i, j
			}
		}
	}

	return -1, -1
}
