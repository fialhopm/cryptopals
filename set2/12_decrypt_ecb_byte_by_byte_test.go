package set2_test

import (
	"strings"
	"testing"

	"github.com/fialhopm/cryptopals/set2"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDecryptEcbByteByByte(t *testing.T) {
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

	// Initialize the oracle and decrypt its unknown buffer.
	oracle := set2.NewOracle(unknown, key)
	decrypted, err := set2.DecryptEcbByteByByte(oracle)
	if err != nil {
		t.Fatalf("set2.DecryptEcbByteByByte: %v", err)
	}

	// TODO: make assertion work without removing trailing newline.
	want := strings.TrimSpace(string(unknown))
	got := strings.TrimSpace(string(decrypted))
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
