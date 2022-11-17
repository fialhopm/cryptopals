package set1_test

import (
	"encoding/hex"
	"testing"

	"github.com/fialhopm/cryptopals/set1"
)

func TestFixedXor(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	var (
		bytes1, bytes2 []byte
		err            error
	)
	bytes1, err = hex.DecodeString(hex1)
	if err != nil {
		t.Fatalf("hex.DecodeString: %v", err)
	}
	bytes2, err = hex.DecodeString(hex2)
	if err != nil {
		t.Fatalf("hex.DecodeString: %v", err)
	}
	result, err := set1.FixedXor(bytes1, bytes2)
	if err != nil {
		t.Fatalf("set1.FixedXor: %v", err)
	}

	want := "746865206b696420646f6e277420706c6179"
	got := hex.EncodeToString(result)
	if want != got {
		t.Errorf("want %#v got %#v", want, got)
	}
}
