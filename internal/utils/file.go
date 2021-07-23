package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// TryReadFile performs specified number of attempts to read a file
// with 1-second delay between attempts.
// Successful read returns file contents as a string.
func TryReadFile(filename string, attempts int) (string, error) {
	if attempts <= 0 {
		return "", fmt.Errorf("invalid number of attempts (%v)", attempts)
	}

	readFile := func(filename string) (string, error) {
		f, err := os.Open(filename)
		if err != nil {
			return "", err
		}
		defer f.Close()

		buf := new(strings.Builder)
		_, err = io.Copy(buf, f)
		if err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	var lastError error
	for i := 0; i < attempts; i++ {
		if lastError != nil {
			time.Sleep(1 * time.Second)
		}

		s, err := readFile(filename)
		if err != nil {
			lastError = err
		} else {
			return s, nil
		}
	}
	return "", fmt.Errorf("exceeded number of attempts (%v): %w", attempts, lastError)
}
