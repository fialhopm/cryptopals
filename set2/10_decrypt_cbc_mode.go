package set2

import (
	"crypto/aes"
	"fmt"

	"github.com/fialhopm/cryptopals/set1"
)

// DecryptCbcMode decrypts a cipher encrypted with AES-128 in CBC mode.
//
// The cipher-block chaining (CBC) mode is more secure than EBC because each
// input to the cipher core is randomized, hiding patterns in the plaintext.
// Nice write-up: https://crypto.stackexchange.com/questions/1129/can-cbc-ciphertext-be-decrypted-if-the-key-is-known-but-the-iv-not
func DecryptCbcMode(data, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher: %v", err)
	}
	decrypted := make([]byte, len(data))
	size := len(key)
	start, end := 0, size
	prevCipherBlock := iv
	for end < len(data) {
		block := make([]byte, size)
		cipher.Decrypt(block, data[start:end])
		block, err = set1.FixedXor(block, prevCipherBlock)
		if err != nil {
			return nil, fmt.Errorf("set1.FixedXor: %v", err)
		}
		copy(decrypted[start:end], block)
		prevCipherBlock = data[start:end]
		start, end = end, end+size
	}
	return decrypted, nil
}
