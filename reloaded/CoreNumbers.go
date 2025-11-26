package reloaded

import (
	"strconv"
	"strings"
)

func Core(s []string) []string {
	res := []string{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "(hex)":
			if len(res) > 0 {
				res[len(res)-1] = ConvertBase(res[len(res)-1], "hex")
			}
			continue

		case "(bin)":
			if len(res) > 0 {
				res[len(res)-1] = ConvertBase(res[len(res)-1], "bin")
			}
			continue

		case "(up)":
			if len(res) > 0 {
				res[len(res)-1] = strings.ToUpper(res[len(res)-1])
			}
			continue

		case "(low)":
			if len(res) > 0 {
				res[len(res)-1] = strings.ToLower(res[len(res)-1])
			}
			continue

		case "(cap)":
			if len(res) > 0 {
				res[len(res)-1] = ToCap(res[len(res)-1])
			}
			continue

		default:
			res = append(res, s[i])
		}
	}

	res = Numbered(res)
	return res
}

// processes tags like (cap, 2), (low, 3), (up, 5)
func Numbered(s []string) []string {
	var check bool
	for i := 0; i < len(s)-1; i++ {
		if s[i] == "(cap," && strings.HasSuffix(s[i+1], ")") {
			s, check = Do_n_times(s, i, "cap")
			if !check { // if no number (cap, )
				continue
			}
			s = append(s[:i], s[i+2:]...)
			i--

		} else if s[i] == "(low," && strings.HasSuffix(s[i+1], ")") {
			s, check = Do_n_times(s, i, "low")
			if !check {
				continue
			}
			s = append(s[:i], s[i+2:]...)
			i--

		} else if s[i] == "(up," && strings.HasSuffix(s[i+1], ")") {
			s, check = Do_n_times(s, i, "up")
			if !check {
				continue
			}
			s = append(s[:i], s[i+2:]...)
			i--

		}
	}
	return s
}

func Do_n_times(s []string, i int, typ string) ([]string, bool) {
	numIndex := i + 1

	valid := IsNumeric(s[numIndex])
	var x int

	if !valid {
		return s, false
	} else {
		str := s[numIndex][0 : len(s[numIndex])-1] // Remove last char ')'
		x, _ = strconv.Atoi(str)
	}

	if x > 0 {
		switch typ {
		case "low":
			for x > 0 && i > 0 {
				s[i-1] = strings.ToLower(s[i-1])
				i--
				x--
			}
		case "up":
			for x > 0 && i > 0 {
				s[i-1] = strings.ToUpper(s[i-1])
				i--
				x--
			}
		case "cap":
			for x > 0 && i > 0 {
				s[i-1] = ToCap(s[i-1])
				i--
				x--
			}
		}
	}
	return s, true
}

func IsNumeric(s string) bool {
	if len(s) <= 1 || s[len(s)-1] != ')' {
		return false
	}
	ss := s[:len(s)-1]
	for i, char := range ss {
		if i == 0 && char == '-' {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
