package set1

import (
	"encoding/hex"
	"unicode"
)

func SingleByteXorCipher(cipher string) (string, error) {
	bytes, err := hex.DecodeString(cipher)
	if err != nil {
		return "", err
	}
	var (
		key       uint8 = 1
		maxScore  float64
		plaintext []byte
	)
	for key > 0 {
		candidate := byteXor(bytes, key)
		score := freqScore(candidate)
		if score > maxScore {
			maxScore = score
			plaintext = candidate
		}
		key++
	}
	return string(plaintext), nil
}

func byteXor(in []byte, key byte) []byte {
	out := make([]byte, len(in))
	for i, b := range in {
		out[i] = b ^ key
	}
	return out
}

func freqScore(bytes []byte) float64 {
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
	for _, r := range string(bytes) {
		if f, ok := freq[unicode.ToLower(r)]; ok {
			score += f
		}
	}
	return score
}
