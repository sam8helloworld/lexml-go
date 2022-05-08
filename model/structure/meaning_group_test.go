package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestMeaningGroupUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  MeaningGroupChild
	}{
		{
			name:  "meaningをUnmarshalできる",
			input: `<meaning>意味</meaning>`,
			want: MeaningGroupChild{
				Type: "meaning",
				Value: Meaning{
					XMLName: xml.Name{Local: "meaning"},
					Value:   "意味",
				},
			},
		},
		{
			name:  "exampleをUnmarshalできる",
			input: `<example>例</example>`,
			want: MeaningGroupChild{
				Type: "example",
				Value: Example{
					XMLName: xml.Name{Local: "example"},
					Value:   "例",
				},
			},
		},
		{
			name:  "subheadをUnmarshalできる",
			input: `<subhead></subhead>`,
			want: MeaningGroupChild{
				Type: "subhead",
				Value: Subhead{
					XMLName: xml.Name{Local: "subhead"},
				},
			},
		},
		{
			name:  "columnをUnmarshalできる",
			input: `<column></column>`,
			want: MeaningGroupChild{
				Type: "column",
				Value: Column{
					XMLName: xml.Name{Local: "column"},
				},
			},
		},
		{
			name:  "keyをUnmarshalできる",
			input: `<key>key</key>`,
			want: MeaningGroupChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   pcdata.PCDATA{Value: "key"},
				},
			},
		},
		{
			name:  "divをUnmarshalできる",
			input: `<div></div>`,
			want: MeaningGroupChild{
				Type: "div",
				Value: Div{
					XMLName: xml.Name{Local: "div"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got MeaningGroupChild
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

func TestMeaningGroupUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got MeaningGroupChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrMeaningGroupChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
