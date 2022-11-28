package set2

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/fialhopm/cryptopals/set1"
)

const blockSize = 16 // block size is always 128 bits.

type Mode int

const (
	Ecb Mode = iota
	Cbc
)

// EncryptAndDetectEbcOrCbc encrypts a buffer with either ECB or CBC mode
// (randomly picked) and then detects which encryption mode was used.
func EncryptAndDetectEbcOrCbc(data []byte) (Mode, error) {
	encrypted, err := EncryptEbcOrCbc(data)
	if err != nil {
		return 0, fmt.Errorf("EncryptEbcOrCbc: %v", err)
	}

	const blockSize = 16
	if set1.DetectAesEcbMode(encrypted, blockSize) {
		return Ecb, nil
	}
	return Cbc, nil
}

// EncryptEbcOrCbc pads the beginning and end of the data with random buffers
// of random length, and then encrypts the result using either ECB or
// CBC mode (also randomly picked).
func EncryptEbcOrCbc(data []byte) ([]byte, error) {
	// Generate random key.
	const keySize = 16
	key, err := GenerateRandBuffer(keySize)
	if err != nil {
		return nil, fmt.Errorf("GenerateRandBuffer: %v", err)
	}

	// Generate 5-10 random bytes to pad the beginning of the plaintext.
	const minPadding, maxPadding = 5, 10
	n, err := generateRandInt(maxPadding - minPadding)
	if err != nil {
		return nil, fmt.Errorf("generateRandInt(: %v", err)
	}
	paddingLeft, err := GenerateRandBuffer(minPadding + n)
	if err != nil {
		return nil, fmt.Errorf("GenerateRandBuffer: %v", err)
	}

	// Generate 5-10 random bytes to pad the end of the plaintext.
	n, err = generateRandInt(maxPadding - minPadding)
	if err != nil {
		return nil, fmt.Errorf("generateRandInt(: %v", err)
	}
	paddingRight, err := GenerateRandBuffer(minPadding + n)
	if err != nil {
		return nil, fmt.Errorf("GenerateRandBuffer: %v", err)
	}

	// Pad plaintext.
	// TODO: do I need to check if final size is a multiple of blockSize?
	newData := make([]byte, len(data)+len(paddingLeft)+len(paddingRight))
	copy(newData, paddingLeft)
	copy(newData, data)
	copy(newData, paddingRight)

	// Randomly pick between ECB and CBC encryption.
	n, err = generateRandInt(1)
	if err != nil {
		return nil, fmt.Errorf("generateRandInt: %v", err)
	}
	var encrypted []byte
	if n == int(Ecb) {
		encrypted, err = encryptEcb(newData, key)
		if err != nil {
			return nil, fmt.Errorf("encryptEcb: %v", err)
		}
	} else {
		iv, err := GenerateRandBuffer(blockSize)
		if err != nil {
			return nil, fmt.Errorf("GenerateRandBuffer: %v", err)
		}
		encrypted, err = encryptCbc(newData, key, iv)
		if err != nil {
			return nil, fmt.Errorf("encryptCbc: %v", err)
		}
	}
	return encrypted, nil
}

// GenerateRandBuffer generates a random buffer.
func GenerateRandBuffer(size int) ([]byte, error) {
	buffer := make([]byte, size)
	_, err := rand.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("rand.Read: %v", err)
	}
	return buffer, nil
}

// generateRandInt generates a random integer.
func generateRandInt(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max+1)))
	if err != nil {
		return 0, fmt.Errorf("rand.Int: %v", err)
	}
	return int(n.Int64()), nil
}

// encryptEcb encrypts data using ECB mode.
func encryptEcb(data, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher: %v", err)
	}
	encrypted := make([]byte, len(data))
	start, end := 0, blockSize
	for end < len(data) {
		cipher.Encrypt(encrypted[start:end], data[start:end])
		start, end = end, end+blockSize
	}
	return encrypted, nil
}

// encryptCbc encrypts data using CBC mode.
func encryptCbc(data, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher: %v", err)
	}
	encrypted := make([]byte, len(data))
	start, end := 0, blockSize
	prevCipherBlock := iv
	for end < len(data) {
		block, err := set1.FixedXor(data[start:end], prevCipherBlock)
		if err != nil {
			return nil, fmt.Errorf("set1.FixedXor: %v", err)
		}
		cipher.Encrypt(encrypted[start:end], block)
		prevCipherBlock = encrypted[start:end]
		start, end = end, end+blockSize
	}
	return encrypted, nil
}
