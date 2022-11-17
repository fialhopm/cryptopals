package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDetectAesEcbMode(t *testing.T) {
	data, err := testutil.ReadAndHexDecode("1_08.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndHexDecode: %v", err)
	}
	candidates, err := set1.DetectAesEcbMode(data)
	if err != nil {
		t.Fatalf("set1.DetectAesEcbMode: %v", err)
	}
	if len(candidates) != 1 {
		t.Errorf("want 1 candidate, got %d", len(candidates))
	}
	want := 132
	got := candidates[0]
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
