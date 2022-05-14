package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIndexUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  IndexChild
	}{
		{
			name:  "meaningをUnmarshalできる",
			input: `<meaning>意味</meaning>`,
			want: IndexChild{
				Type: "meaning",
				Value: Meaning{
					XMLName: xml.Name{Local: "meaning"},
					Value: InnerXML{
						Value: "意味",
					},
				},
			},
		},
		{
			name:  "indexlistをUnmarshalできる",
			input: `<indexlist>1</indexlist>`,
			want: IndexChild{
				Type: "indexlist",
				Value: IndexList{
					XMLName: xml.Name{Local: "indexlist"},
					Value: InnerXML{
						Value: "1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got IndexChild
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

func TestIndexUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got IndexChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrIndexChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
