package reloaded

func ToUpper(res string) string {
	if res[0] >= 'a' && res[0] <= 'z' {
		val := res[0] - 32
		return string(val) + res[1:]
	}
	return res
}
