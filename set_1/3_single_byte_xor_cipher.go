package set1

import (
	"encoding/hex"
	"unicode"
)

func SingleByteXorCipher(cipherHex string) (string, error) {
	cipherB, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}
	var (
		key       uint8 = 1
		maxScore  float64
		plaintext []byte
	)
	for key > 0 {
		b := singleByteXor(cipherB, key)
		s := freqScore(b)
		if s > maxScore {
			maxScore = s
			plaintext = b
		}
		key++
	}
	return string(plaintext), nil
}

func singleByteXor(bytes []byte, key byte) []byte {
	result := make([]byte, len(bytes))
	for i, b := range bytes {
		result[i] = b ^ key
	}
	return result
}

func freqScore(bytes []byte) float64 {
	// https://en.wikipedia.org/wiki/Letter_frequency
	var freq = map[rune]float64{
		'a': 8.167, 'b': 1.492, 'c': 2.782,
		'd': 4.253, 'e': 12.702, 'f': 2.228,
		'g': 2.015, 'h': 6.094, 'i': 6.966,
		'j': 0.153, 'k': 0.772, 'l': 4.025,
		'm': 2.406, 'n': 6.749, 'o': 7.507,
		'p': 1.929, 'q': 0.095, 'r': 5.987,
		's': 6.327, 't': 9.056, 'u': 2.758,
		'v': 0.978, 'w': 2.360, 'x': 0.150,
		'y': 1.974, 'z': 0.074, ' ': 13, // the space is slightly more frequent than the top letter
	}
	var score float64
	for _, r := range string(bytes) {
		if f, ok := freq[unicode.ToLower(r)]; ok {
			score += f
		}
	}
	return score
}