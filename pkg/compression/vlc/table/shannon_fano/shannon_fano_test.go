package shannonfano

import (
	"reflect"
	"testing"
)

func Test_bestDeviderPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one element",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "three elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bestDeviderPosition(tt.codes)
			if got != tt.want {
				t.Errorf("bestDeviderPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assingCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "assignCode test",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 1, Bits: 2, Size: 2},
				{Quantity: 1, Bits: 3, Size: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)
			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got %v want %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "build test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char: 'a',
					Quantity: 1,
					Bits: 3,
					Size: 2,
				},
				'b': code{
					Char: 'b',
					Quantity: 3,
					Size: 1,
					Bits: 0,
				},
				'c': code{
					Char: 'c',
					Quantity: 2,
					Bits: 2,
					Size: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := build(newCharStart(tt.text))

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
