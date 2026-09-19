package vlc_test

import (
	"Archiver/pkg/vlc"
	"reflect"
	"testing"
)

func TestNewhexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want vlc.HexChunks
	}{
		{
			name: "base name",
			str: "20 30 3C 18",
			want: vlc.HexChunks{"20", "30", "3C", "18"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vlc.NewhexChunks(tt.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewhexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		text vlc.HexChunk
		want vlc.BinaryChunk
	}{
		{
			name: "base name",
			text: vlc.HexChunk("2F"),
			want: vlc.BinaryChunk("00101111"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := tt.text
			got := hc.ToBinary()
			if got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		str  vlc.HexChunks
		want vlc.BinaryChunks
	}{
		{
			name: "base name",
			str: vlc.HexChunks{"2F", "80"},
			want: vlc.BinaryChunks{"00101111", "10000000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := tt.str
			got := hc.ToBinary()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		str vlc.BinaryChunks
		want string
	}{
		{
			name: "base name",
			str: vlc.BinaryChunks{"00101111", "10000000"},
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
