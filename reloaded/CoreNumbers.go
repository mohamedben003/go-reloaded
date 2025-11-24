package reloaded

import (
	"strconv"
	"strings"
)

func Core(s []string) []string {
	s = Cleaner(s)
	s = Nested(s)
	return s
}

func Cleaner(s []string) []string {
	changed := true

	for changed {
		changed = false
		newSlice := []string{}

		for i := 0; i < len(s); i++ {
			word := s[i]

			if strings.Contains(word, "(hex)") {
				if word == "(hex)" {
					if len(newSlice) > 0 { // first word is (hex)
						newSlice[len(newSlice)-1] = ConvertBase(newSlice[len(newSlice)-1], "(hex)")
						continue
					} else {
						continue
					}
				}
				parts := strings.Split(word, "(hex)")
				parts[0] = ConvertBase(parts[0], "(hex)")
				newSlice = append(newSlice, parts[0])
				newSlice = append(newSlice, parts[1])
				changed = true
				continue
			}

			if strings.Contains(word, "(bin)") {
				if word == "(bin)" {
					if len(newSlice) > 0 {
						newSlice[len(newSlice)-1] = ConvertBase(newSlice[len(newSlice)-1], "(bin)")
						continue
					} else {
						continue
					}
				}
				parts := strings.Split(word, "(bin)")
				parts[0] = ConvertBase(parts[0], "(bin)")
				newSlice = append(newSlice, parts[0])
				newSlice = append(newSlice, parts[1])
				changed = true
				continue
			}

			if strings.Contains(word, "(up)") {
				if word == "(up)" {
					if len(newSlice) > 0 {
						newSlice[len(newSlice)-1] = strings.ToUpper(newSlice[len(newSlice)-1])
						continue
					} else {
						continue
					}
				}
				parts := strings.Split(word, "(up)")
				parts[0] = strings.ToUpper(parts[0])
				newSlice = append(newSlice, parts[0])
				newSlice = append(newSlice, parts[1])
				changed = true
				continue
			}

			if strings.Contains(word, "(low)") {
				if word == "(low)" {
					if len(newSlice) > 0 {
						newSlice[len(newSlice)-1] = strings.ToLower(newSlice[len(newSlice)-1])
						continue
					} else {
						continue
					}
				}
				parts := strings.Split(word, "(low)")
				parts[0] = strings.ToLower(parts[0])
				newSlice = append(newSlice, parts[0])
				newSlice = append(newSlice, parts[1])
				changed = true
				continue
			}

			if strings.Contains(word, "(cap)") {
				if word == "(cap)" {
					if len(newSlice) > 0 {
						newSlice[len(newSlice)-1] = ToCap(newSlice[len(newSlice)-1])
						continue
					} else {
						continue
					}
				}
				parts := strings.Split(word, "(cap)")
				parts[0] = ToCap(parts[0])
				newSlice = append(newSlice, parts[0])
				newSlice = append(newSlice, parts[1])
				changed = true
				continue
			}

			// If no modification, keep the word
			newSlice = append(newSlice, word)
		}

		s = newSlice
	}
	return s
}

// Nested processes tags like (cap, 2), (low, 3), (up, 5)
func Nested(s []string) []string {
	se := Punctuations([]byte((strings.Join(s, " "))))
	s = strings.Fields(string(se))

	for i := 1; i < len(s)-1; i++ {
		check := false
		if strings.HasSuffix(s[i], "(cap,") || strings.HasPrefix(s[i], "(") && strings.HasPrefix(s[i+1], "cap") {
			s = Do_n_times(s, i, "cap")
			check = true
		}
		if strings.HasSuffix(s[i], "(low,") || strings.HasPrefix(s[i], "(") && strings.HasPrefix(s[i+1], "low") {
			s = Do_n_times(s, i, "low")
			check = true
		}
		if strings.HasSuffix(s[i], "(up,") || strings.HasPrefix(s[i], "(") && strings.HasPrefix(s[i+1], "up") {
			s = Do_n_times(s, i, "up")
			check = true
		}
		if check {
			start := i
			elementsToRemove := 0

			for j := i; j < len(s); j++ {
				elementsToRemove++
				if strings.Contains(s[j], ")") {
					if idx := strings.Index(s[j], ")"); idx < len(s[j])-1 {
						s[j] = s[j][idx+1:]
						elementsToRemove--
					}
					break
				}
			}

			s = append(s[:start], s[start+elementsToRemove:]...)
			i = start - 1
		}
	}
	return s
}

func Do_n_times(s []string, i int, typ string) []string {
	numIndex := i + 1

	// If i+1 is not numeric, try i+2
	val, bracket := IsNumeric(s[numIndex])
	if !val && numIndex+1 < len(s) {
		numIndex = i + 2
		val, bracket = IsNumeric(s[numIndex])
	}

	if !val {
		return s
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
	return s
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
