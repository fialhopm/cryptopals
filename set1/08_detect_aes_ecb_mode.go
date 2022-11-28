package set1

// DetectAesEcbMode detects if a buffer is encrypted with AES-128 in ECB mode.
//
// The detection algorithm is very naive: a buffer is considered to have been
// encrypted with ECB iff it contains at least one duplicate block.
func DetectAesEcbMode(data []byte, blockSize int) bool {
	memo := make(map[string]struct{})
	start, end := 0, blockSize
	for end < len(data) {
		block := string(data[start:end])
		if _, ok := memo[block]; ok {
			return true
		} else {
			memo[block] = struct{}{}
		}
		start = end
		end += blockSize
	}
	return false
}
