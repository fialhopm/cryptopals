package set1_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestFixedXor(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got, err := set1.FixedXor(hex1, hex2)
	if err != nil {
		t.Fatalf("FixedXor: %v", err)
	}
	if want != got {
		t.Fatalf("want %#v got %#v", want, got)
	}
}
