package reloaded

import "fmt"

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
	fmt.Println(string(r[0]))
	for i := 0; i < len(r); i++ {
		if r[i] == '\'' {
			if !check {
				if i < len(r)-1 && r[i+1] == ' ' {
					r = []byte(string(r[:i+1]) + string(r[i+2:]))
				}
				if i > 0 && r[i-1] != ' ' {
					r = []byte(string(r[:i]) + " " + string(r[i:]))
					i++
				}
				check = true
			} else {
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
