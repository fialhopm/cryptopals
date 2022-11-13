package set1

import (
	"encoding/hex"
)

func RepeatingKeyXor(plaintext, key string) string {
	bytes := []byte(plaintext)
	enc := make([]byte, len(bytes))
	for i, b := range bytes {
		k := i % len(key)
		enc[i] = b ^ key[k]
	}
	return hex.EncodeToString(enc)
}
