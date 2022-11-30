package set2

// Pkcs7Padding pads the input buffer to blockSize.
// The value used for padding is the number of padding bytes.
func Pkcs7Padding(data []byte, blockSize int) []byte {
	numPad := 0
	for i := len(data); i%blockSize != 0; i++ {
		numPad++
	}
	for i := 0; i < numPad; i++ {
		data = append(data, byte(numPad))
	}
	return data
}
