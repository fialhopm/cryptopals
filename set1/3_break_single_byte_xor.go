package set1

import (
	"unicode"
)

// BreakSingleByteXor decrypts a cipher encrypted with single-byte XOR.
func BreakSingleByteXor(data []byte) ([]byte, error) {
	var (
		key       uint8 = 1
		maxScore  float64
		decrypted []byte
	)
	for key > 0 {
		candidate := keyXor(data, key)
		score := freqScore(candidate)
		if score > maxScore {
			maxScore = score
			decrypted = candidate
		}
		key++
	}
	return decrypted, nil
}

// keyXor returns the result of XOR-ing every byte of the input buffer with the
// key.
func keyXor(data []byte, key byte) []byte {
	out := make([]byte, len(data))
	for i, b := range data {
		out[i] = b ^ key
	}
	return out
}

// freqScore computes a score that represents the likelihood of the input buffer
// being English text.
func freqScore(data []byte) float64 {
	// https://en.wikipedia.org/wiki/Letter_frequency.
	// The space is slightly more frequent than the top letter.
	var freq = map[rune]float64{
		'a': 8.167, 'b': 1.492, 'c': 2.782,
		'd': 4.253, 'e': 12.702, 'f': 2.228,
		'g': 2.015, 'h': 6.094, 'i': 6.966,
		'j': 0.153, 'k': 0.772, 'l': 4.025,
		'm': 2.406, 'n': 6.749, 'o': 7.507,
		'p': 1.929, 'q': 0.095, 'r': 5.987,
		's': 6.327, 't': 9.056, 'u': 2.758,
		'v': 0.978, 'w': 2.360, 'x': 0.150,
		'y': 1.974, 'z': 0.074, ' ': 13,
	}
	var score float64
	for _, r := range string(data) {
		if f, ok := freq[unicode.ToLower(r)]; ok {
			score += f
		}
	}
	return score
}
