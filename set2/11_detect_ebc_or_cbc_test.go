package set2_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set2"
)

func TestEncryptAndDetectEbcOrCbc(t *testing.T) {
	numEcb := 0
	plaintext := make([]byte, 64)
	for i := range plaintext {
		plaintext[i] = 42
	}
	for i := 0; i < 100; i++ {
		mode, err := set2.EncryptAndDetectEbcOrCbc(plaintext)
		if err != nil {
			t.Fatalf("set2.EncryptAndDetectEbcOrCbc: %v", err)
		}
		if mode == set2.Ecb {
			numEcb++
		}
	}
	// Can't assume that ECB will be picked exactly 50% of the time.
	if numEcb < 40 || numEcb > 60 {
		t.Errorf("want 40 < numEcb < 60, got %d", numEcb)
	}
}
