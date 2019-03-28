package template

import "strings"

// StringSPLIT split string
func StringSPLIT(s string, n int) string {
	rs := []rune(s)
	n2 := n + 1
	rss := string(rs[n:n2])
	return rss
}

// StringTOSHUZU string to shuzu
func StringTOSHUZU(s string) []string {
	var StartString []string
	num := len(s)
	for i := 0; i < num; i++ {
		StartString = append(StartString, (StringSPLIT(s, i)))
	}
	return StartString
}

// StringCHANGE make string change seats
func StringCHANGE(s []string) []string {
	num := len(s)
	for i := 0; i < num/2; i++ {
		s[i], s[num-i-1] = s[num-i-1], s[i]
	}
	return s

}

// StringCHANGE2 make string change seats next
func StringCHANGE2(s []string) []string {
	num := len(s)
	for i := 0; i <= 2*(num/2-1); i++ {
		s[i], s[i+1] = s[i+1], s[i]
		i += 2
	}
	return s
}

// SHUZUtoSTring make shuzu to string
func SHUZUtoSTring(s []string) string {
	strs := strings.Join(s, "")
	return strs
}
