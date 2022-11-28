package set2_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set2"
)

func TestPkcs7Padding(t *testing.T) {
	data := []byte("YELLOW SUBMARINE")
	blockSize := 20
	padded := set2.Pkcs7Padding(data, blockSize)

	want := "YELLOW SUBMARINE\x04\x04\x04\x04"
	got := string(padded)
	if want != string(got) {
		t.Fatalf("want %q got %q", want, got)
	}
}
