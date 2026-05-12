package lib

import "strings"

func SaneSplit(str string, sep string) []string {
	result := make([]string, 0)
	parts := strings.SplitSeq(str, sep)

	for part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}
