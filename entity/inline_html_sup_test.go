package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSupUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "sup内のinline_htmlをUnmarshalできる",
			input: `<sup><b>b</b></sup>`,
			want: []Entity{
				{
					Type: "b",
					Value: SmallB{
						XMLName: xml.Name{Local: "b"},
						Value: InnerXML{
							Value: "b",
						},
					},
				},
			},
		},
		{
			name:  "sup内のinline_lexmlをUnmarshalできる",
			input: `<sup><alabel>alabel</alabel></sup>`,
			want: []Entity{
				{
					Type: "alabel",
					Value: ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value: InnerXML{
							Value: "alabel",
						},
					},
				},
			},
		},
		{
			name:  "sup内のpcdataをUnmarshalできる",
			input: `<sup>pcdata</sup>`,
			want: []Entity{
				{
					Type: "pcdata",
					Value: PCDATA{
						Value: "pcdata",
					},
				},
			},
		},
		{
			name:  "sup内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<sup><b>b</b><alabel>alabel</alabel>pcdata</sup>`,
			want: []Entity{
				{
					Type: "b",
					Value: SmallB{
						XMLName: xml.Name{Local: "b"},
						Value: InnerXML{
							Value: "b",
						},
					},
				},
				{
					Type: "alabel",
					Value: ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value: InnerXML{
							Value: "alabel",
						},
					},
				},
				{
					Type: "pcdata",
					Value: PCDATA{
						Value: "pcdata",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var s SmallSup
			err := xml.Unmarshal([]byte(tt.input), &s)
			got := s.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestSupUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var gotSmallSup SmallSup
	err := xml.Unmarshal(input, &gotSmallSup)
	if !errors.Is(err, ErrSupUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	var gotLargeSup LargeSup
	err = xml.Unmarshal(input, &gotLargeSup)
	if !errors.Is(err, ErrSupUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
