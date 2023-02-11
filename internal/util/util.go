package util

import "strings"

func BaseName(filepath string) string {
	for j := len(filepath) - 1; j >= 0; j-- {
		if filepath[j] == '\\' {
			return filepath[j+1:]
		}
	}
	return ""
}

func Equals(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if strings.ToLower(s) != strings.ToLower(t) {
			return false
		}
	}
	return true
}
