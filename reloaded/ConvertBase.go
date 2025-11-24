package reloaded

import "strconv"

func ConvertBase(res1, base string) string {
	if base == "(bin)" {
		temp, _ := strconv.ParseInt(res1, 2, 64)
		res1 = strconv.FormatInt(temp, 10)
	} else if base == "(hex)" {
		temp, _ := strconv.ParseInt(res1, 16, 64)
		res1 = strconv.FormatInt(temp, 10)
	}
	return res1
}
