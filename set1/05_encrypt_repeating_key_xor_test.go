package set1_test

import (
	"encoding/hex"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestEncryptRepeatingKeyXor(t *testing.T) {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	encrypted := set1.EncryptRepeatingKeyXor(plaintext, key)

	want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	got := hex.EncodeToString(encrypted)
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
