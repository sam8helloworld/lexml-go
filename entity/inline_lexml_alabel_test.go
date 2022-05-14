package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAlabelUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "alabel内のinline_htmlをUnmarshalできる",
			input: `<alabel><b>b</b></alabel>`,
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
			name:  "alabel内のinline_lexmlをUnmarshalできる",
			input: `<alabel><alabel>alabel</alabel></alabel>`,
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
			name:  "alabel内のpcdataをUnmarshalできる",
			input: `<alabel>pcdata</alabel>`,
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
			name:  "alabel内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<alabel><b>b</b><alabel>alabel</alabel>pcdata</alabel>`,
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
			var a ALabel
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

func TestAlabelUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got ALabel
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrAlabelUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
