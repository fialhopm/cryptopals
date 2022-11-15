package testutil

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ReadAndHexDecode reads a hex-encoded file and returns the decoded data.
// Each line is returned as its own buffer.
func ReadAndHexDecode(filename string) ([][]byte, error) {
	path := getTestDataPath(filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	data := make([][]byte, len(lines))
	for i, line := range lines {
		data[i], err = hex.DecodeString(line)
		if err != nil {
			return nil, fmt.Errorf("hex.DecodeString: %v", err)
		}
	}
	return data, nil
}

// ReadAndBase64Decode reads a base64-encoded file and returns the decoded data.
// The entire content is returned as a single buffer.
func ReadAndBase64Decode(filename string) ([]byte, error) {
	path := getTestDataPath(filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	return base64.StdEncoding.DecodeString(sb.String())
}

// getTestDataPath returns the absolute path of the testdata directory.
func getTestDataPath(filename string) string {
	_, path, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(path), "..")
	return filepath.Join(root, "testdata", filename)
}
