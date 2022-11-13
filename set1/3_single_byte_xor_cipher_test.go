package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestSingleByteXorCipher(t *testing.T) {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"
	got, err := set1.SingleByteXorCipher(cipher)
	if err != nil {
		t.Fatalf("SingleByteXorCipher: %v", err)
	}
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}
