package module01

func Reverse(word string) string {

	var res string

	for _, r := range word {
		res = string(r) + res

	}
	return res
}

// func Reverse(word string) string {
// 	var sb strings.Builder

// 	for i := len(word) - 1; i >= 0; i-- {
// 		sb.WriteByte(word[i])
// 	}
// 	return sb.String()
// }
