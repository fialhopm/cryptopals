package set1

import "testing"

func TestSingleByteXorCipher(t *testing.T) {
	cipher := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"
	got, err := SingleByteXorCipher(cipher)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}
