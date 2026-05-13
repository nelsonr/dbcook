package lib

import (
	"fmt"
	"slices"
	"strings"
)

var fieldTypes = map[string]string{
	"int":      "INTEGER",
	"string":   "VARCHAR(255)",
	"text":     "TEXT",
	"date":     "DATE",
	"datetime": "DATETIME",
}

func GenerateTableSql(name string, fields []string) (string, error) {
	var b strings.Builder

	if name == "" {
		return "", fmt.Errorf("error: table name cannot be empty")
	}

	if containsDuplicatedNames(fields) {
		return "", fmt.Errorf("error: field names must be unique")
	}

	fmt.Fprintf(&b, "CREATE TABLE IF NOT EXISTS %s (\n", name)
	fmt.Fprintf(&b, "  id INTEGER PRIMARY KEY,\n")

	for _, f := range fields {
		fieldName := f
		fieldType := fieldTypes["string"]
		hasTypeSep := strings.Contains(f, ":")
		parts := strings.Split(f, ":")

		if hasTypeSep && len(parts) < 2 {
			return "", fmt.Errorf("error: invalid field '%s'", f)
		}

		var err error
		if len(parts) > 1 {
			fieldName, err = NormalizeName(parts[0])
			if err != nil {
				return "", fmt.Errorf("%s", err)
			}

			mappedType, ok := fieldTypes[parts[1]]
			if ok {
				fieldType = mappedType
			} else {
				return "", fmt.Errorf("error: invalid field type '%s'", parts[1])
			}
		}

		fmt.Fprintf(&b, "  %s %s NOT NULL,\n", fieldName, fieldType)
	}

	fmt.Fprintf(&b, ");\n")

	return b.String(), nil
}

// Reports if slice s contains duplicated field names
func containsDuplicatedNames(s []string) bool {
	names := make([]string, 0)

	for _, field := range s {
		name, _, _ := strings.Cut(field, ":")
		names = append(names, name)
	}

	for i, name := range names {
		if slices.Contains(names[i+1:], name) {
			return true
		}
	}

	return false
}
