package bencode

import (
	"fmt"
	"strconv"
	"sort"
)

func Encode(input interface{}) (string, error) {
	switch i := input.(type) {
	case string:
		return encodeString(i)
	case int:
		return encodeInt(i)
	case []interface{}:
		return encodeList(i)
	case map[string]interface{}:
		return encodeDict(i)
	}
	return "", fmt.Errorf("unsupported type: %T", input)
}

func encodeString(input string) (string, error) {
	return strconv.Itoa(len(input)) + ":" + input, nil
}

func encodeInt(input int) (string, error) {
	return "i" + strconv.Itoa(input) + "e", nil
}

func encodeList(input []interface{}) (string, error) {
	result := "l"
	for _, item := range input {
		encoded, err := Encode(item)
		if err != nil {
			return "", err
		}
		result += encoded
	}
	return result + "e", nil
}

func encodeDict(input map[string]interface{}) (string, error) {
	result := "d"
	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		encodedKey, err := encodeString(key)
		if err != nil {
			return "", err
		}
		result += encodedKey
		encodedValue, err := Encode(input[key])
		if err != nil {
			return "", err
		}
		result += encodedValue
	}
	result += "e"
	return result, nil
}
