package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

/*
The Advanced Encryption Standard (AES) is a symmetric block cipher.

A block cipher mode of operation is an algorithm for applying a cipher's
single-block operation repeatedly, which is necessary when the size of the data
is greater than a single block.

The Electronic Codebook (ECB) is a block cipher mode of operation where the data
is divided into blocks and each block is encrypted separately.

Why ECB should not be used anymore: https://crypto.stackexchange.com/questions/20941/why-shouldnt-i-use-ecb-encryption/20946#20946
*/

func DecryptAesEcbMode(data, key string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("DecodeString: %v", err)
	}
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("NewCipher: %v", err)
	}
	decrypted := make([]byte, len(bytes))
	size := len(key)
	start, end := 0, size
	for end < len(bytes) {
		cipher.Decrypt(decrypted[start:end], bytes[start:end])
		start, end = end, end+size
	}
	return string(decrypted), nil
}
