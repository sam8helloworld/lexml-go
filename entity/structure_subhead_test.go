package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSubheadUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  SubheadChild
	}{
		{
			name:  "meaningをUnmarshalできる",
			input: `<meaning>意味</meaning>`,
			want: SubheadChild{
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
			name:  "exampleをUnmarshalできる",
			input: `<example>例</example>`,
			want: SubheadChild{
				Type: "example",
				Value: Example{
					XMLName: xml.Name{Local: "example"},
					Value:   "例",
				},
			},
		},
		{
			name:  "snippetをUnmarshalできる",
			input: `<snippet>snippet</snippet>`,
			want: SubheadChild{
				Type: "snippet",
				Value: Snippet{
					XMLName: xml.Name{Local: "snippet"},
					Value:   PCDATA{Value: "snippet"},
				},
			},
		},
		{
			name:  "columnをUnmarshalできる",
			input: `<column></column>`,
			want: SubheadChild{
				Type: "column",
				Value: Column{
					XMLName: xml.Name{Local: "column"},
				},
			},
		},
		{
			name:  "subheadwordをUnmarshalできる",
			input: `<subheadword>subheadword</subheadword>`,
			want: SubheadChild{
				Type: "subheadword",
				Value: Subheadword{
					XMLName: xml.Name{Local: "subheadword"},
					Value: InnerXML{
						Value: "subheadword",
					},
				},
			},
		},
		{
			name:  "keyをUnmarshalできる",
			input: `<key>key</key>`,
			want: SubheadChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   PCDATA{Value: "key"},
				},
			},
		},
		{
			name:  "divをUnmarshalできる",
			input: `<div></div>`,
			want: SubheadChild{
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
			var got SubheadChild
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

func TestSubheadUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got MeaningGroupChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrMeaningGroupChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
