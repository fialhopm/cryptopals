package set1

import (
	"encoding/hex"
)

func FixedXor(s1, s2 string) (string, error) {
	var (
		b1, b2 []byte
		err    error
	)
	b1, err = hex.DecodeString(s1)
	if err != nil {
		return "", err
	}
	b2, err = hex.DecodeString(s2)
	if err != nil {
		return "", err
	}

	for i := range b1 {
		b1[i] ^= b2[i]
	}
	return hex.EncodeToString(b1), nil
	// TODO: handle slices of different lenghts.
	/*
		if len(b1) != len(b2) {
			return
		output := make([]byte, len(b1))
		for i := range b1 {
			if b1[i] != b2[i] {
			}
		}
	*/

}
