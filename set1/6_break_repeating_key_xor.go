package set1

import (
	"bytes"
	"container/heap"
	"errors"
	"fmt"
	"math/bits"
)

type candidateKey struct {
	size  int
	score int
}

// keyHeap is a min-heap of candidateKeys ordered by score.
// Based on https://cs.opensource.google/go/go/+/refs/tags/go1.18.5:src/container/heap/example_intheap_test.go
type keyHeap []candidateKey

func (h keyHeap) Len() int           { return len(h) }
func (h keyHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h keyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *keyHeap) Push(x interface{}) {
	*h = append(*h, x.(candidateKey))
}

func (h *keyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// BreakRepeatingKeyXor decrypts a cipher encrypted with repeating-key XOR.
//
// Given that we don't know the key size, it returns the N (numCandidates)
// buffers that are most likely to be the decrypted plaintext, ordered by
// decreasing likelihood with respect to the scoring function.
func BreakRepeatingKeyXor(data []byte, numCandidates int) ([][]byte, error) {
	// Compute a score for every key size and store it in a min-heap.
	h := &keyHeap{}
	heap.Init(h)
	const minKeySize, maxKeySize = 2, 40
	for size := minKeySize; size <= maxKeySize; size++ {
		score, err := scoreKeySize(data, size)
		if err != nil {
			return nil, fmt.Errorf("scoreKeySize: %v", err)
		}
		cand := candidateKey{size: size, score: score}
		heap.Push(h, cand)
	}

	// Get the key sizes with the best (lowest) scores.
	candidateKeySizes := make([]int, numCandidates)
	for i := 0; i < numCandidates; i++ {
		candidateKeySizes[i] = heap.Pop(h).(candidateKey).size
	}

	// Decrypt the cipher with each candidate key size.
	candidates := make([][]byte, 0)
	for _, keySize := range candidateKeySizes {
		// Break the cipher into blocks of keySize length and transpose them.
		blocks := make([][]byte, keySize)
		for i := 0; i < keySize; i++ {
			blocks[i] = make([]byte, 0)
		}
		for i, b := range data {
			idx := i % keySize
			blocks[idx] = append(blocks[idx], b)
		}

		// Decrypt each block using single-byte XOR.
		decryptedBlocks := make([][]byte, len(blocks))
		for i, block := range blocks {
			decrypted, err := BreakSingleByteXor(block)
			if err != nil {
				return nil, fmt.Errorf("SingleByteXorCipher: %v", err)
			}
			decryptedBlocks[i] = decrypted
		}

		// Merge the blocks back into a single buffer.
		var bb bytes.Buffer
		var i int
		for {
			if i >= len(decryptedBlocks[0]) {
				break
			}
			for _, block := range decryptedBlocks {
				if i < len(block) {
					bb.WriteByte(block[i])
				}
			}
			i++
		}
		candidates = append(candidates, bb.Bytes())
	}

	return candidates, nil
}

// scoreKeySize computes a score for a key size based on the normalized edit
// distances between the first 10 blocks of encrypted data.
//
// This is a variation of the algorithms suggested in
// https://cryptopals.com/sets/1/challenges/6 that attributes the best score
// to the key that correctly decrypts the sample cipher (1_6.txt).
func scoreKeySize(data []byte, size int) (int, error) {
	const numBlocks = 10
	blocks := make([][]byte, numBlocks)
	for i := 0; i < numBlocks; i++ {
		left := size * i
		right := size * (i + 1)
		blocks[i] = data[left:right]
	}

	var distances int
	for i := 0; i < numBlocks; i++ {
		for j := i + 1; j < numBlocks; j++ {
			d, err := HammingDistance(blocks[i], blocks[j])
			if err != nil {
				return 0, fmt.Errorf("HammingDistance: %v", err)
			}
			distances += d
		}
	}
	average := distances / 45 // C(10, 2)
	normalized := average / size
	return normalized, nil
}

// HammingDistance computes the Hamming distance between two buffers.
func HammingDistance(b1, b2 []byte) (int, error) {
	if len(b1) != len(b2) {
		return 0, errors.New("input strings must have equal length")
	}
	var dist int
	for i, _ := range b1 {
		xor := b1[i] ^ b2[i]
		dist += bits.OnesCount(uint(xor))
	}
	return dist, nil
}
