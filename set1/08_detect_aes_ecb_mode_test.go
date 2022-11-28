package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDetectAesEcbMode(t *testing.T) {
	buffers, err := testutil.ReadAndHexDecode("1_08.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndHexDecode: %v", err)
	}
	const blockSize = 16
	results := make([]int, 0)
	for i, buffer := range buffers {
		if set1.DetectAesEcbMode(buffer, blockSize) {
			results = append(results, i)
		}
	}
	if len(results) != 1 {
		t.Errorf("want 1 result, got %d", len(results))
	}
	want := 132
	got := results[0]
	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
