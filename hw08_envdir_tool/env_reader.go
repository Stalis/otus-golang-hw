package main

import (
	"bytes"
	"io"
	"os"
	"path"
	"strings"
)

const RestrictedNameChars = "="

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ValidateFileName(fileName string) bool {
	return !strings.ContainsAny(fileName, RestrictedNameChars)
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entry, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	res := make(Environment)

	for _, v := range entry {
		if v.IsDir() {
			continue
		}

		if !ValidateFileName(v.Name()) {
			continue
		}

		file, err := os.Open(path.Join(dir, v.Name()))
		if err != nil {
			return nil, err
		}

		readed, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		if strlen := bytes.IndexRune(readed, '\n'); strlen > 0 {
			readed = readed[:strlen]
		}
		file.Close()

		readed = bytes.ReplaceAll(readed, []byte{0x00}, []byte{'\n'})

		value := strings.Split(string(readed), "\n")[0]
		value = strings.TrimRight(value, "\t ")

		res[v.Name()] = EnvValue{
			Value:      value,
			NeedRemove: len(value) == 0,
		}
	}

	return res, nil
}
