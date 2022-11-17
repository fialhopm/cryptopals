package set1

import (
	"crypto/aes"
	"fmt"
)

// DecryptAesEcbMode decrypts a cipher encrypted with AES-128 in ECB mode.
//
// The Advanced Encryption Standard (AES) is a symmetric block cipher.
//
// A block cipher mode of operation is an algorithm for applying a cipher's
// single-block operation repeatedly, which is necessary when the size of the
// data is greater than a single block.
//
// The Electronic Codebook (ECB) is a block cipher mode of operation where the
// data is divided into blocks and each block is encrypted separately.
//
// Why ECB should not be used anymore:
// https://crypto.stackexchange.com/questions/20941/why-shouldnt-i-use-ecb-encryption/20946#20946
func DecryptAesEcbMode(data, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher: %v", err)
	}
	decrypted := make([]byte, len(data))
	size := len(key)
	start, end := 0, size
	for end < len(data) {
		cipher.Decrypt(decrypted[start:end], data[start:end])
		start, end = end, end+size
	}
	return decrypted, nil
}
