package reloaded

func an(s []string) []string {
	for i := 0; i < len(s)-1; i++ {

		vowel := s[i+1][0] == 'a' || s[i+1][0] == 'e' || s[i+1][0] == 'i' || s[i+1][0] == 'o' || s[i+1][0] == 'u' || s[i+1][0] == 'h' ||
			s[i+1][0] == 'A' || s[i+1][0] == 'E' || s[i+1][0] == 'I' || s[i+1][0] == 'O' || s[i+1][0] == 'U' || s[i+1][0] == 'H'

		if s[i] == "a" && vowel {
			s[i] = "an"
		} else if s[i] == "A" && vowel {
			s[i] = "An"
		}
	}
	return s
}
