package set1

import (
	"fmt"
)

// DetectSingleByteXor returns the input buffer that is most likely to have
// been encrypted with single-byte XOR.
func DetectSingleByteXor(data [][]byte) ([]byte, error) {
	var (
		maxScore  float64
		decrypted []byte
	)
	for _, buffer := range data {
		candidate, err := BreakSingleByteXor(buffer)
		if err != nil {
			return nil, fmt.Errorf("BreakSingleByteXor: %v", err)
		}
		score := freqScore([]byte(candidate))
		if score > maxScore {
			maxScore = score
			decrypted = candidate
		}
	}
	return decrypted, nil
}
