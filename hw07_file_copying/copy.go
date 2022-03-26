package main

import (
	"errors"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrFileWriteNotAllowed   = errors.New("file write is not allowed")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	data, err := os.ReadFile(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	if offset > int64(len(data)) {
		return ErrOffsetExceedsFileSize
	}

	file, err := os.Create(toPath)
	if err != nil {
		return ErrFileWriteNotAllowed
	}

	data = data[offset:]
	if limit > 0 && limit < int64(len(data)) {
		data = data[:limit]
	}

	_, err = file.Write(data)
	if err != nil {
		return ErrFileWriteNotAllowed
	}

	return nil
}
