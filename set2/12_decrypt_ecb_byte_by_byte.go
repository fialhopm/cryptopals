package set2

import (
	"fmt"
	"math"

	"github.com/fialhopm/cryptopals/util"
)

const null = byte(0)

type Oracle struct {
	// unknown is a buffer that is appended to the data before encryption.
	unknown []byte
	// key is the encryption key.
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

// DecryptEcbByteByByte decrypts the input oracle's unknown buffer.
func DecryptEcbByteByByte(oracle *Oracle) ([]byte, error) {
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

// decryptUnknown decrypts the input oracle's unknown buffer.
//
// It decrypts one byte at a time by repeatedly calling the oracle's encryption
// method.
func decryptUnknown(oracle *Oracle, blockSize int) ([]byte, error) {
	byteIdx := 0
	known := make([]byte, 0)
	for {
		nextByte, err := decryptNextByte(oracle, byteIdx, known, blockSize)
		if err != nil {
			return nil, fmt.Errorf("decryptNextByte: %v", err)
		}
		// Null byte indicates that there are no more bytes to decrypt.
		if nextByte == null {
			break
		}
		known = append(known, nextByte)
		byteIdx++
	}
	return known, nil
}

// decryptNextByte decrypts the next byte of the input oracle's unknown buffer.
func decryptNextByte(oracle *Oracle, byteIdx int, known []byte, blockSize int) (byte, error) {
	// Assume that all bytes are contained in the ASCII table.
	const (
		asciiStart = 0
		asciiEnd   = 127
	)
	// Create a buffer with the last blockSize-1 bytes that have been
	// decrypted. If the total number of decrypted bytes is less than
	// blockSize-1, use 'A' as padding.
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

	// Compute all possible ciphertexts by encrypting the concatenation of the
	// last blockSize-1 known bytes with each byte in the ASCII table. Store
	// each mapping of ciphertext to byte.
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

	// Resize the block to the known set of bytes that we want to encrypt
	// such that
	block = block[:blockSize-1-(byteIdx%blockSize)]
	encrypted, err := oracle.Encrypt(block)
	if err != nil {
		return 0, fmt.Errorf("oracle.Encrypt: %v", err)
	}
	// Find the encrypted block that we're interested in.
	blockNum := int(math.Floor(float64(byteIdx) / float64(blockSize)))
	encBlock := encrypted[blockSize*blockNum : blockSize*(blockNum+1)]
	// An empty buffer indicates that there are no more bytes to decrypt.
	nullBlock := make([]byte, blockSize)
	if util.IsEqual(encBlock, nullBlock) {
		return null, nil
	}
	// Lookup ciphertext to find the next byte of unknown.
	nextByte, ok := memo[string(encBlock)]
	if !ok {
		return 0, fmt.Errorf("unknown contains bytes that are not in the ASCII table")
	}
	return nextByte, nil
}
