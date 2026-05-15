package lib

import (
	"testing"

	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", ".", "output directory")
	return cmd
}

func TestResolveOutputPathDefault(t *testing.T) {
	expected := "."
	cmd := newTestCmd()
	result := ResolveOutputPath(cmd, ".")
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestResolveOutputPathFromConfig(t *testing.T) {
	expected := "db/migrations"
	cmd := newTestCmd()
	result := ResolveOutputPath(cmd, "testdata")
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestResolveOutputPathFlagOverridesConfig(t *testing.T) {
	expected := "custom/path"
	cmd := newTestCmd()
	cmd.Flags().Set("output", "custom/path")
	result := ResolveOutputPath(cmd, "testdata")
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}

func TestResolveOutputPathFlagOverridesDefault(t *testing.T) {
	expected := "custom/path"
	cmd := newTestCmd()
	cmd.Flags().Set("output", "custom/path")
	result := ResolveOutputPath(cmd, ".")
	if result != expected {
		t.Errorf("expected '%s', got '%s'", expected, result)
	}
}
