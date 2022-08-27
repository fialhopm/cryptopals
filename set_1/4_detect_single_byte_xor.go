package set1

import "strings"

func DetectSingleByteXor(lines []string) (string, error) {
	var (
		maxScore  float64
		plaintext string
	)
	for _, line := range lines {
		candidate, err := SingleByteXorCipher(line)
		if err != nil {
			return "", err
		}
		score := freqScore([]byte(candidate))
		if score > maxScore {
			maxScore = score
			plaintext = candidate
		}
	}
	return strings.TrimSpace(plaintext), nil
}