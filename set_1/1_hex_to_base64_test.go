package set1

import "testing"

func TestHexToBase64(t *testing.T) {
	const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const want = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := HexToBase64(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}
