package hw02unpackstring

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

var ErrInvalidString = errors.New("invalid string")

const EscapeChar = `\`

type token struct {
	char    string
	isDigit bool
	escaped bool
}

func Unpack(input string) (string, error) {
	var prev token

	builder := &strings.Builder{}
	for _, v := range input {
		current := token{char: string(v), isDigit: unicode.IsDigit(v)}

		if prev.char == "" {
			if current.isDigit {
				return "", ErrInvalidString
			}
			prev = current
			continue
		}

		if prev.char == EscapeChar && !prev.escaped {
			if !current.isDigit && current.char != EscapeChar {
				return "", ErrInvalidString
			}

			current.escaped = true
			prev = current
			continue
		}

		if !current.isDigit {
			builder.WriteString(prev.char)
			prev = current
			continue
		}

		count, err := strconv.Atoi(current.char)
		if err != nil {
			return "", errors.Wrap(err, ErrInvalidString.Error())
		}
		builder.WriteString(strings.Repeat(prev.char, count))

		prev = token{}
	}

	builder.WriteString(prev.char)

	return builder.String(), nil
}
