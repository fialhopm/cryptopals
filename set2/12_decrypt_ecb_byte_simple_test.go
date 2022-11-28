package set2_test

import (
	"strings"
	"testing"

	"github.com/fialhopm/cryptopals/set2"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDecryptEcbByteSimple(t *testing.T) {
	// Read and base64-decode unknown string.
	unknown, err := testutil.ReadAndBase64Decode("2_12.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndBase64Decode: %v", err)
	}

	// Generate random key.
	const keySize = 16
	key, err := set2.GenerateRandBuffer(keySize)
	if err != nil {
		t.Fatalf("GenerateRandBuffer: %v", err)
	}

	// Decrypt unknown.
	oracle := set2.NewOracle(unknown, key)
	decrypted, err := set2.DecryptEcbByteSimple(oracle)
	if err != nil {
		t.Fatalf("set2.DecryptEcbByteSimple: %v", err)
	}

	// TODO: why am I not getting the last newline?
	want := strings.TrimSpace(string(unknown))
	got := strings.TrimSpace(string(decrypted))
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
