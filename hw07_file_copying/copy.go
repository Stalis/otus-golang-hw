package main

import (
	"bufio"
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrFileWriteNotAllowed   = errors.New("file write is not allowed")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrReadFailed            = errors.New("read failed")
)

const BufferSize = 100

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer fromFile.Close()

	info, err := fromFile.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}

	size := info.Size()
	if offset > size {
		return ErrOffsetExceedsFileSize
	}

	if limit > size {
		limit = size
	}

	fromFile.Seek(offset, 0)

	toFile, err := os.Create(toPath)
	if err != nil {
		return ErrFileWriteNotAllowed
	}
	defer toFile.Close()

	reader := bufio.NewReader(fromFile)
	writer := bufio.NewWriter(toFile)

	pbar := progressbar.Default(size)
	for {
		n, err := io.CopyN(writer, reader, BufferSize)
		pbar.Add64(n)
		if err != nil {
			if err == io.EOF {
				break
			}
			return ErrFileWriteNotAllowed
		}
	}

	return nil
}
