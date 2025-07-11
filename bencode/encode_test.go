package bencode

import (
	"testing"
)

func TestEncodeString(t *testing.T) {
	result, err := Encode("spam")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "4:spam"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeInt(t *testing.T) {
	result, err := Encode(42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "i42e"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeList(t *testing.T) {
	input := []interface{}{"spam", 42}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "l4:spami42ee"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeDict(t *testing.T) {
	input := map[string]interface{}{
		"eggs": "toast",
		"spam": 42,
	}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Dictionary keys must be sorted
	expected := "d4:eggs5:toast4:spami42ee"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeUnsupportedType(t *testing.T) {
	_, err := Encode(3.14)
	if err == nil {
		t.Error("expected error for unsupported type, got nil")
	}
}
