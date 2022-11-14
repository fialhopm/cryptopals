package set1_test

import (
	"bufio"
	"encoding/hex"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestDetectSingleByteXor(t *testing.T) {
	path, err := getTestDataPath("1_4.txt")
	if err != nil {
		t.Fatalf("getTestDataPath: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("os.Open: %v", err)
	}
	defer file.Close()

	// Input data is hex-encoded.
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	data := make([][]byte, len(lines))
	for i, line := range lines {
		data[i], err = hex.DecodeString(line)
		if err != nil {
			t.Fatalf("DecodeString: %v", err)
		}
	}
	decrypted, err := set1.DetectSingleByteXor(data)

	want := "Now that the party is jumping"
	got := strings.TrimSpace(string(decrypted))
	if err != nil {
		t.Fatalf("DetectSingleByteXor: %v", err)
	}
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}

func getTestDataPath(filename string) (string, error) {
	_, path, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(path), "..")
	return filepath.Join(root, "testdata", filename), nil
}
