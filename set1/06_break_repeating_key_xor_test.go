package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
	"github.com/fialhopm/cryptopals/testutil"
)

func TestBreakRepeatingKeyXor(t *testing.T) {
	data, err := testutil.ReadAndBase64Decode("1_06.txt")
	if err != nil {
		t.Fatalf("testutil.ReadAndBase64Decode: %v", err)
	}
	candidates, err := set1.BreakRepeatingKeyXor(data, 3)
	if err != nil {
		t.Fatalf("set1.BreakRepeatingKeyXor: %v", err)
	}

	// It's sufficient to assert only on the first line.
	want := "I'm back and I'm ringin' the bell"
	// For this sample input, the key size scoring function returns the correct
	// output as the top result.
	got := string(candidates[0][:len(want)])
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}

func TestHammingDistance(t *testing.T) {
	b1 := []byte("this is a test")
	b2 := []byte("wokka wokka!!!")
	want := 37
	got, err := set1.HammingDistance(b1, b2)
	if err != nil {
		t.Fatalf("set1.HammingDistace: %v", err)
	}
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}
