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
