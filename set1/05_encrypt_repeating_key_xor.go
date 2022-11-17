package set1

// EncryptRepeatingKeyXor encrypts plaintext with repeating-key XOR.
func EncryptRepeatingKeyXor(plaintext, key []byte) []byte {
	encrypted := make([]byte, len(plaintext))
	for i, b := range plaintext {
		k := i % len(key)
		encrypted[i] = b ^ key[k]
	}
	return encrypted
}
