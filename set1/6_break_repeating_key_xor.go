package set1

import (
	"container/heap"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/bits"
	"strings"
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

func BreakRepeatingKeyXor(data string, numCandidateKeySizes int) ([]string, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	// Compute a score for every key size and store it in a min-heap.
	h := &keyHeap{}
	heap.Init(h)
	const minKeySize, maxKeySize = 2, 40
	for size := minKeySize; size <= maxKeySize; size++ {
		score, err := scoreKeySize(bytes, size)
		if err != nil {
			return nil, err
		}
		cand := candidateKey{size: size, score: score}
		heap.Push(h, cand)
	}

	// Decrypt using the key sizes with the best (lowest) scores.
	candidateKeySizes := make([]int, numCandidateKeySizes)
	for i := 0; i < numCandidateKeySizes; i++ {
		candidateKeySizes[i] = heap.Pop(h).(candidateKey).size
	}
	candidatePlaintexts := make([]string, 0)
	for _, keySize := range candidateKeySizes {
		// Break the cipher into blocks of length keySize and transpose them.
		blocks := make([][]byte, keySize)
		for i := 0; i < keySize; i++ {
			blocks[i] = make([]byte, 0)
		}
		for i, b := range bytes {
			idx := i % keySize
			blocks[idx] = append(blocks[idx], b)
		}
		// Decrypt each block using single cipher XOR.
		plaintexts := make([]string, len(blocks))
		for i, block := range blocks {
			plaintext, err := SingleByteXorCipher(hex.EncodeToString(block))
			if err != nil {
				return nil, err
			}
			plaintexts[i] = plaintext
		}
		// Combine the plaintexts.
		var sb strings.Builder
		var i int
		for {
			if i >= len(plaintexts[0]) {
				break
			}
			for x := 0; x < len(plaintexts); x++ {
				if i < len(plaintexts[x]) {
					sb.WriteByte(plaintexts[x][i])
				}
			}
			i++
		}
		candidatePlaintexts = append(candidatePlaintexts, sb.String())
	}

	return candidatePlaintexts, nil
}

// scoreKeySize computes a score for a key size based on the normalized edit
// distances between the first 10 blocks of encrypted data.
//
// This is a variation of the algorithms suggested in
// https://cryptopals.com/sets/1/challenges/6 that attributes the best score
// to the correct key for the sample input.
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
				return 0, err
			}
			distances += d
		}
	}
	avg := distances / 45 // C(10, 2)
	norm := avg / size
	return norm, nil
}

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
