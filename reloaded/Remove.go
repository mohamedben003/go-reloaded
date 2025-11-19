package reloaded

func Remove(res []string, i int) []string {
	return append(res[:i], res[i+1:]...)
}
