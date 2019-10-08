package utils

import "strings"

// Contains - a string in a []string or not
func Contains(sli []string, target string) bool {
	for _, s := range sli {
		if s == target {
			return true
		}
	}

	return false
}

// SplitWithComma - split content with comma as spliter
func SplitWithComma(content string) []string {
	if len(content) == 0 {
		return []string{}
	}
	elements := strings.Split(content, ",")
	for i := 0; i < len(elements); i++ {
		elements[i] = strings.Trim(elements[i], " ")
	}

	return elements
}
