package set2

import (
	"errors"
)

var ErrInvalidPadding = errors.New("invalid PKCS7 padding")

// StripPkcs7Padding removes the padding bytes from the input buffer.
// Returns an error if the padding is not valid PKCS#7.
//
// TODO: how should inputs without padding be handled? Assuming that padding is
// not added to inputs with length equal to block size.
func StripPkcs7Padding(data []byte) ([]byte, error) {
	size := len(data)
	padByte := data[size-1]
	numPads := int(padByte)
	// Validate padding.
	for i := size - numPads; i < size; i++ {
		if data[i] != padByte {
			return nil, ErrInvalidPadding
		}
	}
	// Strip padding.
	return data[:size-numPads], nil
}
