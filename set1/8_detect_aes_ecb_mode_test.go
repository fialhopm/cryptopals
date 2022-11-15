package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestDetectAesEcbMode(t *testing.T) {
	data, err := testutil.ReadAndHexDecode("1_8.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndHexDecode: %v", err)
	}
	got, err := set1.DetectAesEcbMode(data)
	if err != nil {
		t.Fatalf("set1.DetectAesEcbMode: %v", err)
	}
	want := 132
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}
