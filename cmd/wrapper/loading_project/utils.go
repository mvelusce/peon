package loading_project

import "strings"

func TrimSuffix(s string, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func TrimPrefix(s string, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		s = s[0+len(prefix):]
	}
	return s
}
