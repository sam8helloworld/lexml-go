package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAccentUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "accent内のinline_htmlをUnmarshalできる",
			input: `<accent><b>b</b></accent>`,
			want: []Entity{
				{
					Type: "b",
					Value: SmallB{
						XMLName: xml.Name{Local: "b"},
						Value:   "b",
					},
				},
			},
		},
		{
			name:  "accent内のinline_lexmlをUnmarshalできる",
			input: `<accent><alabel>alabel</alabel></accent>`,
			want: []Entity{
				{
					Type: "alabel",
					Value: ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value:   "alabel",
					},
				},
			},
		},
		{
			name:  "accent内のpcdataをUnmarshalできる",
			input: `<accent>pcdata</accent>`,
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
			name:  "accent内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<accent><b>b</b><alabel>alabel</alabel>pcdata</accent>`,
			want: []Entity{
				{
					Type: "b",
					Value: SmallB{
						XMLName: xml.Name{Local: "b"},
						Value:   "b",
					},
				},
				{
					Type: "alabel",
					Value: ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value:   "alabel",
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
			var a Accent
			err := xml.Unmarshal([]byte(tt.input), &a)
			got := a.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestAccentUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Accent
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrAccentUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
