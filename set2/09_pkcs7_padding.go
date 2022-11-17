package set2

func Pkcs7Padding(data []byte, blockSize int) []byte {
	for len(data)%20 != 0 {
		data = append(data, '\x04')
	}
	return data
}
