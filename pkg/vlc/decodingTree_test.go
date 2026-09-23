package vlc_test

import (
	"Archiver/pkg/vlc"
	"reflect"
	"testing"
)

func TestEncodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   vlc.EncodingTable
		want vlc.DecodingTree
	}{
		{
			name: "base name",

			et: vlc.EncodingTable{'a': "11", 'b': "1001", 'z': "0101"},

			want: vlc.DecodingTree{

				Zero: &vlc.DecodingTree{

					One: &vlc.DecodingTree{

						Zero: &vlc.DecodingTree{

							One: &vlc.DecodingTree{
								Value: "z",
							},
						},
					},
				},

				One: &vlc.DecodingTree{
					Zero: &vlc.DecodingTree{
						Zero: &vlc.DecodingTree{
							One: &vlc.DecodingTree{
								Value: "b",
							},
						},
					},

					One: &vlc.DecodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.et.DecodingTree()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
