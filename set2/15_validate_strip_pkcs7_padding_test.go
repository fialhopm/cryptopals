package set2_test

import (
	"testing"

	"github.com/fialhopm/cryptopals/set2"
	"github.com/fialhopm/cryptopals/util"
)

func TestStripPkcs7Padding(t *testing.T) {
	tests := []struct {
		input   []byte
		want    []byte
		wantErr error
	}{
		{
			input:   []byte("ICE ICE BABY\x04\x04\x04\x04"),
			want:    []byte("ICE ICE BABY"),
			wantErr: nil,
		},
		{
			input:   []byte("ICE ICE BABY\x05\x05\x05\x05"),
			want:    nil,
			wantErr: set2.ErrInvalidPadding,
		},
		{
			input:   []byte("ICE ICE BABY\x01\x02\x03\x04"),
			want:    nil,
			wantErr: set2.ErrInvalidPadding,
		},
	}
	for _, test := range tests {
		got, err := set2.StripPkcs7Padding(test.input)
		if !util.IsEqual(test.want, got) {
			t.Errorf("want %#v got %#v", test.want, got)
		}
		if test.wantErr != nil && err != nil && test.wantErr.Error() != err.Error() {
			t.Errorf("want err %#v got err %#v", test.wantErr, err)
		}
		if (test.wantErr == nil && err != nil) || (test.wantErr != nil && err == nil) {
			t.Errorf("want err %#v got err %#v", test.wantErr, err)
		}
	}
}
