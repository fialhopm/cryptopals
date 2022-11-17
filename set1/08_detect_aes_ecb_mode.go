package set1

// DetectAesEcbMode detects which input buffers, if any, have been encrypted
// with AES-128 in ECB mode and returns them.
//
// The detection algorithm is very naive: a buffer is considered to have been
// encrypted with ECB iff it contains at least one duplicate 16-byte block.
// TODO: change signature to take in a single buffer. This will improve 2_11.
func DetectAesEcbMode(data [][]byte) ([]int, error) {
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
	return candidates, nil
}
