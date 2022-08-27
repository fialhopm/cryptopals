package set1

import (
	"encoding/hex"
)

func RepeatingKeyXor(ptext, key string) string {
	enc := make([]byte, len(ptext))
	for i := 0; i < len(ptext); i++ {
		k := i % len(key)
		enc[i] = ptext[i] ^ key[k]
	}
	return hex.EncodeToString(enc)
}
