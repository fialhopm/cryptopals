package set1

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"testing"
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

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	want := "Now that the party is jumping"
	got, err := DetectSingleByteXor(lines)
	if err != nil {
		t.Fatalf("DetectSingleByteXor: %v", err)
	}
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}

func getTestDataPath(filename string) (string, error) {
	_, path, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(path), "..")
	return filepath.Join(root, "testdata", filename), nil
}
