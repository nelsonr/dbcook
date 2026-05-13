package lib

import (
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

var scenario01 = `CREATE TABLE IF NOT EXISTS blog_posts (
  id INTEGER PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
);
`

var scenario02 = `CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY,
  age INTEGER NOT NULL,
);
`

var scenario03 = `CREATE TABLE IF NOT EXISTS articles (
  id INTEGER PRIMARY KEY,
  body TEXT NOT NULL,
);
`

var scenario04 = `CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  price INTEGER NOT NULL,
);
`

var scenario05 = `CREATE TABLE IF NOT EXISTS tags (
  id INTEGER PRIMARY KEY,
);
`

var scenario06 = `CREATE TABLE IF NOT EXISTS posts (
  id INTEGER PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
);
`

func assertSqlMatch(t *testing.T, expected, result string) {
	t.Helper()
	if result != expected {
		diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
			A:        difflib.SplitLines(expected),
			B:        difflib.SplitLines(result),
			FromFile: "Expected",
			ToFile:   "Result",
			Context:  3,
		})
		t.Errorf("Generated SQL mismatch:\n%s", diff)
	}
}

func TestGenerateTableSqlExplicitStringType(t *testing.T) {
	result, _ := GenerateTableSql("blog_posts", []string{"title:string", "url:string"})
	assertSqlMatch(t, scenario01, result)
}

func TestGenerateTableSqlDefaultStringType(t *testing.T) {
	result, _ := GenerateTableSql("blog_posts", []string{"title", "url"})
	assertSqlMatch(t, scenario01, result)
}

func TestGenerateTableSqlUnknownType(t *testing.T) {
	_, err := GenerateTableSql("blog_posts", []string{"title:string", "url:lolbanana"})
	if err == nil {
		t.Errorf("An invalid field type should return an error")
	}
}

func TestGenerateTableSqlIntType(t *testing.T) {
	result, err := GenerateTableSql("users", []string{"age:int"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	assertSqlMatch(t, scenario02, result)
}

func TestGenerateTableSqlTextType(t *testing.T) {
	result, err := GenerateTableSql("articles", []string{"body:text"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	assertSqlMatch(t, scenario03, result)
}

func TestGenerateTableSqlMixedTypes(t *testing.T) {
	result, err := GenerateTableSql("products", []string{"name:string", "description:text", "price:int"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	assertSqlMatch(t, scenario04, result)
}

func TestGenerateTableSqlEmptyFields(t *testing.T) {
	result, err := GenerateTableSql("tags", []string{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	assertSqlMatch(t, scenario05, result)
}

func TestGenerateTableSqlInvalidTypes(t *testing.T) {
	cases := []string{"float", "bool", "json", "uuid"}
	for _, typeName := range cases {
		_, err := GenerateTableSql("test_table", []string{"field:" + typeName})
		if err == nil {
			t.Errorf("Expected error for unsupported type %q, got nil", typeName)
		}
	}
}

func TestGenerateTableSqlEmptyName(t *testing.T) {
	_, err := GenerateTableSql("", []string{"title:string"})
	if err == nil {
		t.Errorf("Expected error for empty table name, got nil")
	}
}

func TestGenerateTableSqlEmptyFieldName(t *testing.T) {
	_, err := GenerateTableSql("posts", []string{":string"})
	if err == nil {
		t.Errorf("Expected error for empty field name, got nil")
	}
}

func TestGenerateTableSqlColonOnlyField(t *testing.T) {
	_, err := GenerateTableSql("posts", []string{":"})
	if err == nil {
		t.Errorf("Expected error for colon-only field, got nil")
	}
}

func TestGenerateTableSqlMultipleColons(t *testing.T) {
	result, err := GenerateTableSql("posts", []string{"name:string:extra"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	assertSqlMatch(t, scenario06, result)
}

func TestGenerateTableSqlUppercaseType(t *testing.T) {
	_, err := GenerateTableSql("users", []string{"age:INT"})
	if err == nil {
		t.Errorf("Expected error for uppercase type %q, got nil", "INT")
	}
}

func TestGenerateTableSqlDuplicateFields(t *testing.T) {
	_, err := GenerateTableSql("posts", []string{"title:string", "title:int"})
	if err == nil {
		t.Errorf("Expected error for duplicated field names, got nil")
	}
}

func TestContainsDuplicatedFieldNames(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect bool
	}{
		{"empty slice", []string{}, false},
		{"single field", []string{"title"}, false},
		{"single field with type", []string{"title:string"}, false},
		{"no duplicates", []string{"title:string", "body:text", "age:int"}, false},
		{"duplicate names same type", []string{"title:string", "title:string"}, true},
		{"duplicate names different types", []string{"title:string", "title:int"}, true},
		{"duplicate plain names no colon", []string{"title", "title"}, true},
		{"duplicate via mixed colon and plain", []string{"title:string", "title"}, true},
		{"multiple colons uses first segment", []string{"title:string:extra", "title:int"}, true},
		{"case insensitive duplicate", []string{"Title:string", "title:string"}, true},
		{"all same name", []string{"x:string", "x:int", "x:text"}, true},
		{"duplicate not adjacent", []string{"a:string", "b:int", "a:text"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsDuplicatedFieldNames(tc.input)
			if result != tc.expect {
				t.Errorf("ContainsDuplicatedFieldNames(%v) = %v, expected %v", tc.input, result, tc.expect)
			}
		})
	}
}
