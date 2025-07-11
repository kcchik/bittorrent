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

func TestEncodeEmptyString(t *testing.T) {
	result, err := Encode("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "0:"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeUnicodeString(t *testing.T) {
	result, err := Encode("🍳")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "4:🍳"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeStringWithColon(t *testing.T) {
	result, err := Encode("eggs:spam")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "9:eggs:spam"
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

func TestEncodeZeroInt(t *testing.T) {
	result, err := Encode(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "i0e"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeNegativeInt(t *testing.T) {
	result, err := Encode(-42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "i-42e"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeLargeInt(t *testing.T) {
	result, err := Encode(1234567890)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "i1234567890e"
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

func TestEncodeEmptyList(t *testing.T) {
	input := []interface{}{}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "le"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeNestedList(t *testing.T) {
	input := []interface{}{[]interface{}{"spam", 42}, "eggs"}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "ll4:spami42ee4:eggse"
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
	expected := "d4:eggs5:toast4:spami42ee"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeEmptyDict(t *testing.T) {
	input := map[string]interface{}{}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "de"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeNestedDict(t *testing.T) {
	input := map[string]interface{}{
		"eggs": map[string]interface{}{"toast": 42},
		"spam": []interface{}{"spam", "eggs"},
	}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "d4:eggsd5:toasti42ee4:spaml4:spam4:eggsee"
	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}

func TestEncodeDictKeyOrdering(t *testing.T) {
	input := map[string]interface{}{
		"spam": 42,
		"eggs": "toast",
		"toast": "eggs",
	}
	result, err := Encode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "d4:eggs5:toast4:spami42e5:toast4:eggse"
	if result != expected {
		t.Errorf("dictionary keys not sorted: %q", result)
	}
}

func TestEncodeUnsupportedType(t *testing.T) {
	_, err := Encode(3.14)
	if err == nil {
		t.Error("expected error for unsupported type, got nil")
	}
}
