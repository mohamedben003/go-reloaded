package reloaded

func Punctuations(r []byte) []byte {
	// First, remove multiple spaces
	for i := 1; i < len(r); i++ {
		if r[i-1] == ' ' && r[i] == ' ' {
			r = append(r[:i-1], r[i:]...)
			i--
		}
	}

	// Handle single quotes
	check := false
	for i := 0; i < len(r); i++ {
		if r[i] == '\'' {
			if !check {
				// Opening quote - remove space after it
				if i+1 < len(r) && r[i+1] == ' ' {
					r = append(r[:i+1], r[i+2:]...)
				}
				check = true
			} else {
				// Closing quote - remove space before it
				if i > 0 && r[i-1] == ' ' {
					r = append(r[:i-1], r[i:]...)
					i--
				}
				check = false
			}
		}
	}

	// Handle punctuation marks (., ,, !, ?, :, ;)
	for i := 1; i < len(r); i++ {

		if r[i] == '.' || r[i] == ',' || r[i] == '!' || r[i] == '?' || r[i] == ':' || r[i] == ';' {
			// Remove space before punctuation
			if r[i-1] == ' ' {
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
