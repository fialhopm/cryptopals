package set1_test

import (
	"strings"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDetectSingleByteXor(t *testing.T) {
	data, err := testutil.ReadAndHexDecode("1_04.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndHexDecode: %v", err)
	}
	decrypted := set1.DetectSingleByteXor(data)

	want := "Now that the party is jumping"
	got := strings.TrimSpace(string(decrypted))
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
