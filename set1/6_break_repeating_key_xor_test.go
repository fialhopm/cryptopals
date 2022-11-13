package set1_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestBreakRepeatingKeyXor(t *testing.T) {
	path, err := getTestDataPath("1_6.txt")
	if err != nil {
		t.Fatalf("getTestDataPath: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("os.Open: %v", err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}
	results, err := set1.BreakRepeatingKeyXor(sb.String(), 3)
	if err != nil {
		t.Fatalf("BreakRepeatingKeyXor: %v", err)
	}
	// For this sample input, the key size scoring function returns the correct
	// output as the top result.
	got := results[0]
	// It's sufficient to assert only on the first line.
	want := "I'm back and I'm ringin' the bell"
	if got[:len(want)] != "I'm back and I'm ringin' the bell" {
		t.Errorf("want %#v got %#v", want, got)
	}
}

func TestHammingDistance(t *testing.T) {
	b1 := []byte("this is a test")
	b2 := []byte("wokka wokka!!!")
	want := 37
	got, err := set1.HammingDistance(b1, b2)
	if err != nil {
		t.Fatalf("hammingDistace: %v", err)
	}
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}
