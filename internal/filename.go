package lib

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func GenerateFileName(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", fmt.Errorf("error: name cannot be empty")
	}

	ts := time.Now().UnixMilli()

	filename, err := NormalizeName(name)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	var result strings.Builder
	fmt.Fprintf(&result, "%v_create_%s.sql", ts, filename)

	return result.String(), nil
}

func NormalizeName(name string) (string, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return "", fmt.Errorf("error: name cannot be empty")
	}

	sep := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	result := string(sep.ReplaceAll([]byte(name), []byte("_")))
	result = ConvertToSnakeCase(result)

	return result, nil
}

// Converts the string s to snake_case
//
// Example: "theBestName" -> "the_best_name"
func ConvertToSnakeCase(s string) string {
	var result strings.Builder

	for i, char := range s {
		if char >= 'A' && char <= 'Z' && i > 0 {
			fmt.Fprint(&result, "_")
		}
		fmt.Fprintf(&result, "%v", strings.ToLower(string(char)))
	}

	return result.String()
}
