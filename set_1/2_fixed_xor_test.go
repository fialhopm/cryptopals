package set1

import "testing"

func TestFixedXor(t *testing.T) {
	const (
		s1   = "1c0111001f010100061a024b53535009181c"
		s2   = "686974207468652062756c6c277320657965"
		want = "746865206b696420646f6e277420706c6179"
	)
	got, err := FixedXor(s1, s2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}
