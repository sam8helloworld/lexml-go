package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCaptionUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "caption内のinline_htmlをUnmarshalできる",
			input: `<caption><b>b</b></caption>`,
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
			name:  "caption内のinline_lexmlをUnmarshalできる",
			input: `<caption><alabel>alabel</alabel></caption>`,
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
			name:  "caption内のpcdataをUnmarshalできる",
			input: `<caption>pcdata</caption>`,
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
			name:  "caption内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<caption><b>b</b><alabel>alabel</alabel>pcdata</caption>`,
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
			var c Caption
			err := xml.Unmarshal([]byte(tt.input), &c)
			got := c.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestCaptionUnmarshalXML_Caption(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Caption
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrCaptionUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
