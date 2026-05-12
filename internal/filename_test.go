package lib

import (
	"strings"
	"testing"
)

func TestGenerateFileNameEmptyString(t *testing.T) {
	_, err := GenerateFileName("")
	if err == nil {
		t.Errorf("expected error for empty table name, got nil")
	}
}

func TestGenerateFileNameWhitespaceOnly(t *testing.T) {
	_, err := GenerateFileName("   ")
	if err == nil {
		t.Errorf("expected error for whitespace-only table name, got nil")
	}
}

func TestGenerateFileNameBasic(t *testing.T) {
	expected := "_create_users.sql"
	result, err := GenerateFileName("users")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasSuffix(result, expected) {
		t.Errorf("expected suffix '%s', got '%s'", expected, result)
	}
}

func TestGenerateFileNameSpacesReplacedWithUnderscores(t *testing.T) {
	expected := "_create_my_table.sql"
	result, err := GenerateFileName("my table")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasSuffix(result, expected) {
		t.Errorf("expected suffix '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameEmptyString(t *testing.T) {
	_, err := NormalizeName("")
	if err == nil {
		t.Errorf("expected error for empty name, got nil")
	}
}

func TestNormalizeNameWhitespaceOnly(t *testing.T) {
	_, err := NormalizeName("   ")
	if err == nil {
		t.Errorf("expected error for whitespace-only name, got nil")
	}
}

func TestNormalizeNameSpacesReplacedWithUnderscores(t *testing.T) {
	expected := "my_table"
	result, err := NormalizeName("my table")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameHyphensReplacedWithUnderscores(t *testing.T) {
	expected := "my_table"
	result, err := NormalizeName("my-table")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameAlreadyNormalized(t *testing.T) {
	expected := "my_table"
	result, err := NormalizeName("my_table")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameTabReplacedWithUnderscore(t *testing.T) {
	expected := "my_table"
	result, err := NormalizeName("my\ttable")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameMixedSeparators(t *testing.T) {
	expected := "my___table"
	result, err := NormalizeName("my - table")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameLeadingTrailingSpaces(t *testing.T) {
	expected := "orders"
	result, err := NormalizeName(" orders ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameNumbers(t *testing.T) {
	expected := "orders2024"
	result, err := NormalizeName("orders2024")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameDot(t *testing.T) {
	expected := "user_profile"
	result, err := NormalizeName("user.profile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestNormalizeNameSlash(t *testing.T) {
	expected := "my_table"
	result, err := NormalizeName("my/table")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}
