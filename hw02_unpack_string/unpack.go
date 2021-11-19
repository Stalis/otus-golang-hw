package hw02unpackstring

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var prev string

	builder := &strings.Builder{}
	for _, v := range input {
		current := string(v)
		isDigit := unicode.IsDigit(v)

		if prev == "" {
			if isDigit {
				return "", ErrInvalidString
			}
			prev = current
			continue
		}

		if !isDigit {
			builder.WriteString(prev)
			prev = current
			continue
		}

		count, err := strconv.Atoi(current)
		if err != nil {
			return "", errors.Wrap(err, ErrInvalidString.Error())
		}
		builder.WriteString(strings.Repeat(prev, count))

		prev = ""
	}

	builder.WriteString(prev)

	return builder.String(), nil
}
