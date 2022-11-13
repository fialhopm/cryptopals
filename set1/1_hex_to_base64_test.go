package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestHexToBase64(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := set1.HexToBase64(hex)
	if err != nil {
		t.Fatalf("HexToBase64: %v", err)
	}
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}
