package set1

import (
	"encoding/hex"
	"fmt"
)

func FixedXor(hex1, hex2 string) (string, error) {
	var (
		b1, b2 []byte
		err    error
	)
	b1, err = hex.DecodeString(hex1)
	if err != nil {
		return "", fmt.Errorf("DecodeString: %v", err)
	}
	b2, err = hex.DecodeString(hex2)
	if err != nil {
		return "", fmt.Errorf("DecodeString: %v", err)
	}
	for i := range b1 {
		b1[i] ^= b2[i]
	}
	return hex.EncodeToString(b1), nil
}
