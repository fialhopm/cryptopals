package set1

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

func TestDetectSingleByteXor(t *testing.T) {
	path, err := getTestDataPath("4.txt")
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
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cur, "test_data", filename), nil
}
