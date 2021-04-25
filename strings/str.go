package str

import "strings"

func Lower(s string) string {
	return strings.ToLower((s)) 
}

func Contain(s, sub string) bool {
	return strings.Contains(s, sub)
}

func Joined(elems []string, sep string) string {
	return strings.Join(elems, sep)
}