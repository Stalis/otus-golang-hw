package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var prev string

	builder := &strings.Builder{}
	for _, v := range input {
		char := string(v)
		isDigit := unicode.IsDigit(v)

		if prev == "" && isDigit {
			return "", ErrInvalidString
		}

		if prev == "" {
			prev = char
			continue
		}

		if !isDigit {
			builder.WriteString(prev)
			prev = char
			continue
		}

		count, err := strconv.Atoi(char)
		if err != nil {
			return "", err
		}
		builder.WriteString(strings.Repeat(prev, count))

		prev = ""
	}

	if prev != "" {
		builder.WriteString(prev)
	}

	return builder.String(), nil
}
