package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(h string) (string, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return "", fmt.Errorf("DecodeString: %v", err)
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
