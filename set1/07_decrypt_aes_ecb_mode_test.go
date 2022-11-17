package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDecryptAesEcbMode(t *testing.T) {
	data, err := testutil.ReadAndBase64Decode("1_07.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndBase64Decode: %v", err)
	}
	key := []byte("YELLOW SUBMARINE")
	decrypted, err := set1.DecryptAesEcbMode(data, key)
	if err != nil {
		t.Fatalf("set1.DecryptAesEcbMode: %v", err)
	}

	want := "I'm back and I'm ringin' the bell"
	got := string(decrypted[:len(want)])
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
