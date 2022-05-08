package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTableUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  TableChild
	}{
		{
			name:  "captionをUnmarshalできる",
			input: `<caption>caption</caption>`,
			want: TableChild{
				Type: "caption",
				Value: Caption{
					XMLName: xml.Name{Local: "caption"},
					Value:   "caption",
				},
			},
		},
		{
			name:  "trをUnmarshalできる",
			input: `<tr></tr>`,
			want: TableChild{
				Type: "tr",
				Value: Tr{
					XMLName: xml.Name{Local: "tr"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got TableChild
			err := xml.Unmarshal([]byte(tt.input), &got)
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestTableUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got TableChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrTableChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
