package reloaded

func Punctuations(r []byte) []byte {

	// Handle single quotes
	check := false
	for i := 0; i < len(r); i++ {
		if r[i] == '\'' {
			if !check {
				//Opening '
				if i < len(r)-1 && i > 0 && isAlpha(r[i-1]) && isAlpha(r[i+1]) {
					continue
				}

				if i < len(r)-1 && r[i+1] == ' ' {
					r = []byte(string(r[:i+1]) + string(r[i+2:]))
				}
				if i > 0 && r[i-1] != ' ' {
					r = []byte(string(r[:i]) + " " + string(r[i:]))
					i++
				}
				check = true
			} else {
				// Closing '
				if i < len(r)-1 && i > 0 && isAlpha(r[i-1]) && isAlpha(r[i+1]) {
					continue
				}
				if i > 0 && r[i-1] == ' ' {
					r = []byte(string(r[:i-1]) + string(r[i:]))
					i--
				}
				if i < len(r)-1 && r[i+1] != ' ' {
					r = []byte(string(r[:i+1]) + " " + string(r[i+1:]))
				}
				check = false
			}
		}
	}

	for i := 0; i < len(r); i++ {
		if r[i] == '.' || r[i] == ',' || r[i] == '!' || r[i] == '?' || r[i] == ':' || r[i] == ';' {
			// Remove space before punctuation
			if i > 0 && r[i-1] == ' ' {
				r = append(r[:i-1], r[i:]...)
				i--
			}

			// Add space after punctuation if not followed by space or another punctuation
			if i+1 < len(r) {
				nextIsPunct := r[i+1] == '.' || r[i+1] == ',' || r[i+1] == '!' ||
					r[i+1] == '?' || r[i+1] == ':' || r[i+1] == ';' || r[i+1] == '\''

				if r[i+1] != ' ' && !nextIsPunct {
					r = append(r[:i+1], append([]byte{' '}, r[i+1:]...)...)
					i++
				}
			}
		}
	}

	return r
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
