package bencode

import (
  "reflect"
  "testing"
)

func TestDecodeString(t *testing.T) {
  input := "4:spam"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := "spam"
  if result != expected {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeEmptyString(t *testing.T) {
  input := "0:"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := ""
  if result != expected {
    t.Errorf("got %q, want %q", result, expected)
  }
}

func TestDecodeUnicodeString(t *testing.T) {
  input := "4:🍳"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := "🍳"
  if result != expected {
    t.Errorf("got %q, want %q", result, expected)
  }
}

func TestDecodeStringWithColon(t *testing.T) {
  input := "9:eggs:spam"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := "eggs:spam"
  if result != expected {
    t.Errorf("got %q, want %q", result, expected)
  }
}

func TestDecodeInt(t *testing.T) {
  input := "i42e"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := 42
  if result != expected {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeZeroInt(t *testing.T) {
  input := "i0e"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := 0
  if result != expected {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeNegativeInt(t *testing.T) {
  input := "i-42e"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := -42
  if result != expected {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeLargeInt(t *testing.T) {
  input := "i1234567890e"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := 1234567890
  if result != expected {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeList(t *testing.T) {
  input := "l4:spami42e4:eggse"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := []interface{}{"spam", 42, "eggs"}
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeEmptyList(t *testing.T) {
  input := "le"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := []interface{}{}
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeNestedList(t *testing.T) {
  input := "ll4:spami42ee4:eggse"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := []interface{}{
    []interface{}{"spam", 42},
    "eggs",
  }
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeDict(t *testing.T) {
  input := "d4:eggs5:toast4:spami42ee"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := map[string]interface{}{
    "eggs": "toast",
    "spam": 42,
  }
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeEmptyDict(t *testing.T) {
  input := "de"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := map[string]interface{}{}
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeNestedDict(t *testing.T) {
  input := "d4:eggsd5:toasti42ee4:spaml4:spam4:eggsee"
  result, err := Decode(input)
  if err != nil {
    t.Fatalf("unexpected error: %v", err)
  }
  expected := map[string]interface{}{
    "eggs": map[string]interface{}{"toast": 42},
    "spam": []interface{}{"spam", "eggs"},
  }
  if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v, want %v", result, expected)
  }
}

func TestDecodeInvalidInput(t *testing.T) {
  cases := []string{
    "",
    "x42e",
    "i42",
    "l4:spami42e", // missing end marker
    "d4:eggs5:toast4:spami42e", // missing end marker
    "999spam", // missing colon
    "i4twoe", // invalid integer
  }
  for _, input := range cases {
    _, err := Decode(input)
    if err == nil {
      t.Errorf("expected error for input %q, got nil", input)
    }
  }
}
