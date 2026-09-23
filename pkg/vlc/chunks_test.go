package vlc_test

import (
	"Archiver/pkg/vlc"
	"reflect"
	"testing"
)

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		str  vlc.BinaryChunks
		want string
	}{
		{
			name: "base name",
			str:  vlc.BinaryChunks{"00101111", "10000000"},
			want: "0010111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := tt.str
			got := bc.Join()
			if got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want vlc.BinaryChunks
	}{
		{
			name: "test binaryChunks",
			data: []byte{20, 30, 60, 18},
			want: vlc.BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vlc.NewBinChunks(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
