package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func Core(s []string) []string {
	s = cleaner(s)
	s = handleComplexTags(s)
	return s
}

func cleaner(s []string) []string {
	changed := true

	for changed {
		changed = false
		newSlice := []string{}

		for i := 0; i < len(s); i++ {
			word := s[i]

			// Check if current word contains (hex)
			if strings.Contains(word, "(hex)") {
				// Extract the part before (hex)
				parts := strings.Split(word, "(hex)")
				if len(parts) >= 2 && len(newSlice) > 0 {
					// Convert previous word
					newSlice[len(newSlice)-1] = ConvertBase(newSlice[len(newSlice)-1], "(hex)")
					// Add the remaining part if it exists
					if parts[1] != "" {
						newSlice = append(newSlice, parts[1])
					}
					changed = true
					continue
				}
			}

			// Check if current word contains (bin)
			if strings.Contains(word, "(bin)") {
				parts := strings.Split(word, "(bin)")
				if len(parts) >= 2 && len(newSlice) > 0 {
					newSlice[len(newSlice)-1] = ConvertBase(newSlice[len(newSlice)-1], "(bin)")
					if parts[1] != "" {
						newSlice = append(newSlice, parts[1])
					}
					changed = true
					continue
				}
			}

			// Check if current word contains (up) - but NOT followed by comma
			if strings.Contains(word, "(up)") {
				parts := strings.Split(word, "(up)")
				if len(parts) >= 2 && len(newSlice) > 0 {
					newSlice[len(newSlice)-1] = strings.ToUpper(newSlice[len(newSlice)-1])
					if parts[1] != "" {
						newSlice = append(newSlice, parts[1])
					}
					changed = true
					continue
				}
			}

			// Check if current word contains (low) - but NOT followed by comma
			if strings.Contains(word, "(low)") {
				parts := strings.Split(word, "(low)")
				if len(parts) >= 2 && len(newSlice) > 0 {
					newSlice[len(newSlice)-1] = strings.ToLower(newSlice[len(newSlice)-1])
					if parts[1] != "" {
						newSlice = append(newSlice, parts[1])
					}
					changed = true
					continue
				}
			}

			// Check if current word contains (cap) - but NOT followed by comma
			if strings.Contains(word, "(cap)") {
				parts := strings.Split(word, "(cap)")
				if len(parts) >= 2 && len(newSlice) > 0 {
					newSlice[len(newSlice)-1] = ToCap(newSlice[len(newSlice)-1])
					if parts[1] != "" {
						newSlice = append(newSlice, parts[1])
					}
					changed = true
					continue
				}
			}

			// If no modification, keep the word
			newSlice = append(newSlice, word)
		}

		s = newSlice
	}

	return s
}

// Helper function to check if the next element starts with a comma
// func isFollowedByComma(s []string, index int) bool {
// 	if index+1 < len(s) {
// 		return strings.HasPrefix(s[index+1], ",")
// 	}
// 	return false
// }

// handleComplexTags processes tags like (cap, 2), (low, 3), (up, 5)
func handleComplexTags(s []string) []string {
	se := Punctuations([]byte((strings.Join(s, " "))))
	s = strings.Fields(string(se))
	var check bool
	for i := 1; i < len(s)-1; i++ {
		j := i
		if strings.HasSuffix(s[i], "(cap,") || strings.HasPrefix(s[i], "(") && strings.HasPrefix(s[i+1], "cap") {
			s, check = Do_n_times(s, j, "cap")
			if !check {
				continue
			}
		}
		if strings.HasSuffix(s[i], "(low,") || strings.HasPrefix(s[i], "(") && strings.HasPrefix(s[i+1], "low") {
			s, check = Do_n_times(s, j, "low")
			if !check {
				continue
			}
		}
		if strings.HasSuffix(s[i], "(up,") || strings.HasPrefix(s[i], "(") && strings.HasPrefix(s[i+1], "up") {
			fmt.Println("worked")
			fmt.Println("worked")
			fmt.Println("worked")
			fmt.Println("worked")
			s, check = Do_n_times(s, j, "up")
			if !check {
				continue
			}
		}

	}
	return s
}

func Do_n_times(s []string, i int, typ string) ([]string, bool) {
	check := true
	numIndex := i + 1

	// If i+1 is not numeric, try i+2
	val, bracket := IsNumeric(s[numIndex])
	if !val && numIndex+1 < len(s) {
		numIndex = i + 2
		val, bracket = IsNumeric(s[numIndex])
	}

	if !val {
		return s, false
	}
	var x int
	if bracket == 2 {
		str := s[numIndex][0 : len(s[numIndex])-1] // Remove last char ')'
		x, _ = strconv.Atoi(str)
	} else {
		x, _ = strconv.Atoi(s[numIndex])
	}

	if x > 0 {
		if typ == "low" {
			for i-1 >= 0 {
				s[i-1] = strings.ToLower(s[i-1])
				i--
			}
		} else if typ == "up" {
			for i-1 >= 0 {
				s[i-1] = strings.ToUpper(s[i-1])
				i--
			}
		} else {
			for i-1 >= 0 {
				s[i-1] = ToCap(s[i-1])
				i--
			}
		}
	}
	return s, check
}

func IsNumeric(s string) (bool, int) {
	if len(s) == 0 {
		return false, 0
	}

	for i, char := range s {
		if i == len(s)-1 {
			if char == ')' {
				return true, 2
			} else if char >= '0' && char <= '9' {
				return true, 1
			} else {
				return false, 0
			}
		}
		if char < '0' || char > '9' {
			return false, 0
		}
	}
	return true, 1
}
