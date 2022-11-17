package set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// HexToBase64 converts a hex string to a base64 encoded string.
func HexToBase64(data string) (string, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("hex.DecodeString: %v", err)
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
