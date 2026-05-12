package lib

import (
	"reflect"
	"testing"
)

func TestSaneSplitBasic(t *testing.T) {
	result := SaneSplit("a,b,c", ",")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSaneSplitFiltersEmptyParts(t *testing.T) {
	result := SaneSplit(",a,,b,", ",")
	expected := []string{"a", "b"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSaneSplitEmptyString(t *testing.T) {
	result := SaneSplit("", ",")
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestSaneSplitNoSeparator(t *testing.T) {
	result := SaneSplit("abc", ",")
	expected := []string{"abc"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSaneSplitOnlySeparators(t *testing.T) {
	result := SaneSplit(",,,", ",")
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %v", result)
	}
}

func TestSaneSplitMultiCharSeparator(t *testing.T) {
	result := SaneSplit("a::b::c", "::")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
