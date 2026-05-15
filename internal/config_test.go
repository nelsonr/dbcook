package lib

import (
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	conf, err := ReadConfigFile("testdata")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if conf.OutputPath != "db/migrations" {
		t.Errorf("expected 'db/migrations', got '%s'", conf.OutputPath)
	}
}

func TestReadConfigFileMissingDir(t *testing.T) {
	_, err := ReadConfigFile("testdata/nonexistent")
	if err == nil {
		t.Errorf("expected error for missing directory, got nil")
	}
}

func TestDecodeConfigValidToml(t *testing.T) {
	expected := "db/migrations"
	data := `output_path = "db/migrations"`
	conf, err := DecodeConfig(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if conf.OutputPath != expected {
		t.Errorf("expected '%s', got '%s'", expected, conf.OutputPath)
	}
}

func TestDecodeConfigEmptyString(t *testing.T) {
	conf, err := DecodeConfig("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if conf.OutputPath != "" {
		t.Errorf("expected empty OutputPath, got '%s'", conf.OutputPath)
	}
}

func TestDecodeConfigInvalidToml(t *testing.T) {
	_, err := DecodeConfig("OutputPath = [invalid")
	if err == nil {
		t.Errorf("expected error for invalid TOML, got nil")
	}
}

func TestDecodeConfigUnknownFieldIgnored(t *testing.T) {
	expected := "db/migrations"
	data := `
output_path = "db/migrations"
UnknownField = "ignored"
`
	conf, err := DecodeConfig(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if conf.OutputPath != expected {
		t.Errorf("expected '%s', got '%s'", expected, conf.OutputPath)
	}
}

func TestDecodeConfigSnakeCaseField(t *testing.T) {
	expected := "db/migrations"
	data := `
output_path = "db/migrations"
`
	conf, err := DecodeConfig(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if conf.OutputPath != expected {
		t.Errorf("expected '%s', got '%s'", expected, conf.OutputPath)
	}
}
