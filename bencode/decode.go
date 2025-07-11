package bencode

import (
	"fmt"
	"strconv"
	"strings"
)

func Decode(input string) (interface{}, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("input cannot be empty")
	}
	result, length, err := decodeNext(input)
	if err != nil {
		return nil, err
	}
	if length != len(input) {
		return nil, fmt.Errorf("extra characters after valid bencode: %s", input[length:])
	}
	return result, nil
}

func decodeNext(input string) (interface{}, int, error) {
	switch input[0] {
	case 'i':
		return decodeInt(input)
	case 'l':
		return decodeList(input)
	case 'd':
		return decodeDict(input)
	default:
		if input[0] >= '0' && input[0] <= '9' {
			return decodeString(input)
		}
		return nil, 0, fmt.Errorf("invalid character %c in input", input[0])
	}
}

func decodeString(input string) (string, int, error) {
	lengthEnd := strings.Index(input, ":")
	if lengthEnd == -1 {
		return "", 0, fmt.Errorf("invalid string format")
	}
	length, err := strconv.Atoi(input[:lengthEnd])
	if err != nil {
		return "", 0, fmt.Errorf("invalid string length: %v", err)
	}
	lengthTotal := lengthEnd + 1 + length
	return input[lengthEnd + 1 : lengthTotal], lengthTotal, nil
}

func decodeInt(input string) (int, int, error) {
	numEnd := strings.Index(input, "e")
	if numEnd < 1 {
		return 0, 0, fmt.Errorf("invalid integer format")
	}
	numStr := input[1:numEnd]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid integer value: %v", err)
	}
	return num, len(numStr) + 2, nil
}

func decodeList(input string) ([]interface{}, int, error) {
	result := []interface{}{}
	for i := 1; i < len(input); {
		if input[i] == 'e' {
			return result, i + 1, nil
		}
		item, length, err := decodeNext(input[i:])
		if err != nil {
			return nil, 0, err
		}
		i += length
		result = append(result, item)
	}
	return nil, 0, fmt.Errorf("missing end marker for list")
}

func decodeDict(input string) (map[string]interface{}, int, error) {
	result := make(map[string]interface{})
	for i := 1; i < len(input); {
		if input[i] == 'e' {
			return result, i + 1, nil
		}
		key, keyLength, err := decodeString(input[i:])
		if err != nil {
			return nil, 0, err
		}
		i += keyLength
		value, valueLength, err := decodeNext(input[i:])
		if err != nil {
			return nil, 0, err
		}
		i += valueLength
		result[key] = value
	}
	return nil, 0, fmt.Errorf("missing end marker for dictionary")
}
