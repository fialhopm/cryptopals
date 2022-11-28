package set2

import (
	"fmt"
	"math"

	"github.com/fialhopm/cryptopals/util"
)

type Oracle struct {
	// unknown is a buffer that is appended to the data before encryption.
	unknown []byte
	// key is encryption key.
	key []byte
}

func NewOracle(unknown, key []byte) *Oracle {
	return &Oracle{unknown: unknown, key: key}
}

// Encrypt appends an unknown buffer to the data and then encrypts it using
// ECB mode.
func (o *Oracle) Encrypt(data []byte) ([]byte, error) {
	data = append(data, o.unknown...)
	return encryptEcb(data, o.key)
}

// TODO: add doc comment.
func DecryptEcbByteSimple(oracle *Oracle) ([]byte, error) {
	// Detect block size.
	blocksize, err := detectBlockSize(oracle)
	if err != nil {
		return nil, fmt.Errorf("detectBlockSize: %v", err)
	}

	// Ensure oracle is using ECB encryption.
	result, err := isEcb(oracle, blocksize)
	if err != nil {
		return nil, fmt.Errorf("isEcb: %v", err)
	}
	if !result {
		return nil, fmt.Errorf("oracle is not using ECB encryption")
	}

	return decryptUnknown(oracle, blocksize)
}

// detectBlockSize detects the block size used by the input Oracle.
//
// It assumes that the only possible sizes are 64, 128, and 256 bits.
// TODO: is block size equivalent to key size in ECB?
func detectBlockSize(oracle *Oracle) (int, error) {
	var blockSize int
	candidateSizes := []int{8, 12, 16}
	for _, size := range candidateSizes {
		// Create buffer with size equal to the candidate block size.
		data := make([]byte, size, size+1)
		for i := 0; i < size; i++ {
			data[i] = 'A'
		}
		// Encrypt buffer.
		encSize, err := oracle.Encrypt(data)
		if err != nil {
			return 0, fmt.Errorf("oracle: %v", err)
		}
		// Add one more element to the buffer and encrypt again.
		data = append(data, 'A')
		encSizePlus1, err := oracle.Encrypt(data)
		if err != nil {
			return 0, fmt.Errorf("oracle: %v", err)
		}
		// If the candidate is the actual block size, then the first "size"
		// bytes should be the same across the two encrypted buffers.
		if util.IsEqual(encSize[:size], encSizePlus1[:size]) {
			blockSize = size
			break
		}
	}
	if blockSize == 0 {
		return 0, fmt.Errorf("oracle block size is not 64, 128, nor 256 bits")
	}
	return blockSize, nil
}

// isEcb returns whether the input Oracle uses ECB encryption.
//
// It returns true if, given a buffer where the first two blocks are equal,
// the Oracle produces a ciphertext where the first two blocks are also equal.
func isEcb(oracle *Oracle, blockSize int) (bool, error) {
	// Create a buffer with space for two blocks and set each block to the same
	// sequence of bytes.
	data := make([]byte, 2*blockSize)
	for i := 0; i < 2; i++ {
		for j := 0; j < blockSize; j++ {
			idx := (blockSize * i) + j
			data[idx] = byte(j)
		}
	}
	// Encrypt and ensure that the first two blocks of ciphertext are equal.
	encrypted, err := oracle.Encrypt(data)
	if err != nil {
		return false, fmt.Errorf("oracle: %v", err)
	}
	first := encrypted[:blockSize]
	second := encrypted[blockSize : 2*blockSize]
	return util.IsEqual(first, second), nil
}

// TODO: add doc comment.
func decryptUnknown(oracle *Oracle, blockSize int) ([]byte, error) {
	byteIdx := 0
	known := make([]byte, 0)
	for {
		nextByte, err := decryptNextByte(oracle, byteIdx, known, blockSize)
		if err != nil {
			return nil, fmt.Errorf("decryptNextByte: %v", err)
		}
		if nextByte == byte(3) {
			break
		}
		known = append(known, nextByte)
		byteIdx++
	}

	return known, nil
}

// TODO: clean up and add doc comment.
func decryptNextByte(oracle *Oracle, byteIdx int, known []byte, blockSize int) (byte, error) {
	// Assume that each byte of Oracle.unknown is either an Engligh letter,
	// digit, or special character. TODO: update me
	const (
		asciiStart = 0
		asciiEnd   = 2000
	)
	// Compute and store all possible ciphertexts.
	block := make([]byte, blockSize)
	blockIdx := blockSize - 2
	knownIdx := len(known) - 1
	for knownIdx >= 0 && blockIdx >= 0 {
		block[blockIdx] = known[knownIdx]
		blockIdx--
		knownIdx--
	}
	blockIdx = 0
	for blockIdx < blockSize-1 && block[blockIdx] == byte(0) {
		block[blockIdx] = 'A'
		blockIdx++
	}
	memo := make(map[string]byte, asciiEnd-asciiStart)
	for i := asciiStart; i < asciiEnd; i++ {
		candidate := byte(i)
		block[blockSize-1] = candidate
		encrypted, err := oracle.Encrypt(block)
		if err != nil {
			return 0, fmt.Errorf("oracle.Encrypt: %v", err)
		}
		memo[string(encrypted[:blockSize])] = candidate
	}
	// Remove last byte from block and encrypt.
	block = block[:blockSize-1-(byteIdx%blockSize)]
	encrypted, err := oracle.Encrypt(block)
	if err != nil {
		return 0, fmt.Errorf("oracle.Encrypt: %v", err)
	}
	// Lookup ciphertext to find next unknown byte.
	blockNum := int(math.Floor(float64(byteIdx) / float64(blockSize)))
	encBlock := encrypted[blockSize*blockNum : blockSize*(blockNum+1)]
	nullBlock := make([]byte, blockSize)
	if util.IsEqual(encBlock, nullBlock) {
		return byte(3), nil
	}
	nextByte, ok := memo[string(encBlock)]
	if !ok {
		return 0, fmt.Errorf("unknown contains bytes lower than %d or higher than %d", asciiStart, asciiEnd)
	}
	return nextByte, nil
}
