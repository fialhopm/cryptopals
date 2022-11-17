package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestHexToBase64(t *testing.T) {
	data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	got, err := set1.HexToBase64(data)
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if err != nil {
		t.Fatalf("set1.HexToBase64: %v", err)
	}
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
