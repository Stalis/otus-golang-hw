package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"time"

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

	if limit > 0 && limit < size {
		size = limit
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
	progress := int64(0)
	for {
		n, err := io.CopyN(writer, reader, MinInt64(BufferSize, size-progress))
		progress += n
		pbar.Set64(progress)

		time.Sleep(time.Millisecond * 100)

		if progress >= size {
			break
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return ErrFileWriteNotAllowed
		}
	}

	return nil
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
