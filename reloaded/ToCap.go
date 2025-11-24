package reloaded

func ToCap(res string) string {
	r := ""
	for i := 0; i < len(res); i++ {
		if i == 0 {
			if res[0] >= 'a' && res[0] <= 'z' {
				r = string(res[0] - 32)
			} else {
				r = string(res[0])
			}
		} else {
			if res[i] >= 'A' && res[i] <= 'Z' {
				r = r + string(res[i]+32)
			} else {
				r = r + string(res[i])
			}
		}
	}
	return r
}
