package module01

func NumInList(list []int, num int) bool {

	for _, value := range list {
		if value == num {
			return true
		}
	}

	return false
}
