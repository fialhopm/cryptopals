package set1_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestDecryptAesEcbMode(t *testing.T) {
	path, err := getTestDataPath("1_7.txt")
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
	key := "YELLOW SUBMARINE"
	decrypted, err := set1.DecryptAesEcbMode(sb.String(), key)
	if err != nil {
		t.Fatalf("DecryptAesEcbMode: %v", err)
	}
	want := "I'm back and I'm ringin' the bell"
	got := decrypted[:len(want)]
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
