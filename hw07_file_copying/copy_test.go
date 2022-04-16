package main

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TempFilesDir = "otus-copy-test"
)

func CreateTempFile(content string) (string, error) {
	file, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func TestCopy(t *testing.T) {
	destPath := path.Join(os.TempDir(), "dest")

	t.Run("empty file", func(t *testing.T) {
		source, err := CreateTempFile("")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		err = Copy(source, destPath, 0, 0)
		require.Nil(t, err)

		dest, err := os.Open(destPath)
		require.Nil(t, err)

		data, err := io.ReadAll(dest)
		require.Nil(t, err)
		require.Empty(t, data)
	})

	t.Run("not existing file", func(t *testing.T) {
		err := Copy("/unexisting_file", destPath, 0, 0)
		require.Error(t, err)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("existing file to not accessed place", func(t *testing.T) {
		source, err := CreateTempFile("")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		err = Copy(source, "/root/not_accessed_file", 0, 0)
		require.Error(t, err)
		require.Equal(t, ErrFileWriteNotAllowed, err)
	})

	t.Run("existing file smaller than buffer", func(t *testing.T) {
		input, err := os.Open("testdata/input.txt")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		buf := make([]byte, BufferSize-5)
		_, err = input.Read(buf)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		str := string(buf)
		source, err := CreateTempFile(str)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		err = Copy(source, destPath, 0, 0)
		require.Nil(t, err)

		output, err := ioutil.ReadFile(destPath)
		require.Nil(t, err)
		require.Equal(t, buf, output)
	})
}
