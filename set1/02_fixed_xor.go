package set1

import "fmt"

// FixedXor returns the result of XOR-ing every byte of two input buffers.
func FixedXor(data1, data2 []byte) ([]byte, error) {
	if len(data1) != len(data2) {
		return nil, fmt.Errorf("buffers must have equal length")
	}
	for i := range data1 {
		data1[i] ^= data2[i]
	}
	return data1, nil
}
