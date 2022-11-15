package set1

// DetectSingleByteXor returns the input buffer that is most likely to have
// been encrypted with single-byte XOR.
func DetectSingleByteXor(data [][]byte) []byte {
	var (
		maxScore  float64
		decrypted []byte
	)
	for _, buffer := range data {
		candidate := BreakSingleByteXor(buffer)
		score := freqScore([]byte(candidate))
		if score > maxScore {
			maxScore = score
			decrypted = candidate
		}
	}
	return decrypted
}
