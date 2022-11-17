package set1

import "fmt"

// DetectAesEcbMode detects which input buffer has been encrypted with AES-128
// in ECB mode and returns it.
//
// The detection algorithm is very naive: a buffer is considered to have been
// encrypted with ECB iff it contains at least one duplicate 16-byte block.
func DetectAesEcbMode(data [][]byte) (int, error) {
	const keySize = 16
	candidates := make([]int, 0)
	for i, buffer := range data {
		memo := make(map[string]struct{})
		start, end := 0, keySize
		for end < len(buffer) {
			block := string(buffer[start:end])
			if _, ok := memo[block]; ok {
				candidates = append(candidates, i)
				break
			} else {
				memo[block] = struct{}{}
			}
			start = end
			end += keySize
		}
	}
	if len(candidates) == 0 {
		return 0, fmt.Errorf("failed to detect AES ECB encrypted ciphertext")
	}
	if len(candidates) > 1 {
		return 0, fmt.Errorf("detected more than one AES ECB encrypted ciphertext")
	}
	return candidates[0], nil
}
