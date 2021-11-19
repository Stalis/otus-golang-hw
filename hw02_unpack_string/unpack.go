package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type node struct {
	empty   bool
	char    rune
	repeats int
	next    *node
}

func emptyNode() *node {
	return &node{empty: true, repeats: -1}
}

func newNode(r rune) *node {
	return &node{empty: false, char: r, repeats: -1}
}

func parseString(input string) (*node, error) {
	first := emptyNode()
	current := first

	for _, v := range input {
		if unicode.IsDigit(v) {
			if current.empty || current.repeats > 0 {
				return nil, ErrInvalidString
			}

			val, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}

			current.repeats = val
		} else {
			if current.empty {
				current.empty = false
				current.char = v
			} else {
				if current.repeats < 0 {
					current.repeats = 1
				}
				current.next = newNode(v)
				current = current.next
			}
		}
	}

	if current.repeats <= 0 {
		current.repeats = 1
	}

	return first, nil
}

func Unpack(input string) (string, error) {
	first, err := parseString(input)
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	for node := first; node != nil; node = node.next {
		if node.repeats <= 0 || node.empty {
			continue
		}
		builder.WriteString(strings.Repeat(string(node.char), node.repeats))
	}

	return builder.String(), nil
}
