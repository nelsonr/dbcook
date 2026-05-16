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

var reservedFields = []string{
	"created_at",
	"updated_at",
}

func GenerateTableSql(name string, fields []string) (string, error) {
	var (
		output    strings.Builder
		fieldName string
		fieldType string
		err       error
		ok        bool
	)

	if name == "" {
		return "", fmt.Errorf("error: table name cannot be empty")
	}

	if ContainsDuplicatedFieldNames(fields) {
		return "", fmt.Errorf("error: field names must be unique")
	}

	fmt.Fprintf(&output, "CREATE TABLE IF NOT EXISTS %s (\n", name)
	fmt.Fprintf(&output, "  id INTEGER PRIMARY KEY,\n")

	for _, f := range fields {
		fieldName = f
		fieldType = fieldTypes["string"]
		hasTypeSep := strings.Contains(f, ":")
		parts := strings.Split(f, ":")

		if hasTypeSep && len(parts) < 2 {
			return "", fmt.Errorf("error: invalid field '%s'", f)
		}

		if len(parts) > 1 {
			fieldName, err = NormalizeName(parts[0])
			if err != nil {
				return "", fmt.Errorf("%s", err)
			}

			fieldType, ok = fieldTypes[parts[1]]
			if !ok {
				return "", fmt.Errorf("error: invalid field type '%s'", parts[1])
			}
		}

		if slices.Contains(reservedFields, fieldName) {
			return "", fmt.Errorf("error: '%s' is a reserved field name", fieldName)
		}

		fmt.Fprintf(&output, "  %s %s NOT NULL,\n", fieldName, fieldType)
	}

	fmt.Fprintf(&output, "  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,\n")
	fmt.Fprintf(&output, "  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP\n")
	fmt.Fprintf(&output, ");\n")

	return output.String(), nil
}

// Reports if fields contains duplicated names
func ContainsDuplicatedFieldNames(fields []string) bool {
	checkList := make(map[string]bool, len(fields))

	for _, field := range fields {
		name, _, _ := strings.Cut(field, ":")
		name = strings.ToLower(name)

		if checkList[name] {
			return true
		}

		checkList[name] = true
	}

	return false
}
