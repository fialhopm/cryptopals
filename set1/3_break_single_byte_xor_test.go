package set1_test

import (
	"encoding/hex"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestBreakSingleByteXor(t *testing.T) {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, err := hex.DecodeString(cipher)
	if err != nil {
		t.Fatalf("DecodeString: %v", err)
	}
	decrypted, err := set1.BreakSingleByteXor(bytes)
	if err != nil {
		t.Fatalf("SingleByteXorCipher: %v", err)
	}

	want := "Cooking MC's like a pound of bacon"
	got := string(decrypted)
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
