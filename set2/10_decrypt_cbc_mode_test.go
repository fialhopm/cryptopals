package set2_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set2"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDecryptCbcMode(t *testing.T) {
	// The file is base64-encoded, even though it's not specified in the
	// problem statement.
	data, err := testutil.ReadAndBase64Decode("2_10.txt")
	if err != nil {
		t.Fatalf("testutil.Read: %v", err)
	}
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16) // initialization vector is all ASCII 0.
	decrypted, err := set2.DecryptCbcMode(data, key, iv)
	if err != nil {
		t.Fatalf("set2.DecryptCbcMode: %v", err)
	}

	// It's sufficient to assert only on the first line.
	want := "I'm back and I'm ringin' the bell"
	got := string(decrypted[:len(want)])
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
