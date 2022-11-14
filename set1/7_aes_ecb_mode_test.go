package set1_test

import (
	"bufio"
	"encoding/base64"
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
	data, err := base64.StdEncoding.DecodeString(sb.String())
	if err != nil {
		t.Fatalf("DecodeString: %v", err)
	}
	key := []byte("YELLOW SUBMARINE")
	decrypted, err := set1.DecryptAesEcbMode(data, key)
	if err != nil {
		t.Fatalf("DecryptAesEcbMode: %v", err)
	}

	want := "I'm back and I'm ringin' the bell"
	got := string(decrypted[:len(want)])
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
