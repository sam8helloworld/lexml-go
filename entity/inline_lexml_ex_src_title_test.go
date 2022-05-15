package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExSrcTitleUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "ex-src-title内のinline_htmlをUnmarshalできる",
			input: `<ex-src-title><b>b</b></ex-src-title>`,
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
			name:  "ex-src-title内のinline_lexmlをUnmarshalできる",
			input: `<ex-src-title><alabel>alabel</alabel></ex-src-title>`,
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
			name:  "ex-src-title内のpcdataをUnmarshalできる",
			input: `<ex-src-title>pcdata</ex-src-title>`,
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
			name:  "ex-src-title内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<ex-src-title><b>b</b><alabel>alabel</alabel>pcdata</ex-src-title>`,
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
			var e ExSrcTitle
			err := xml.Unmarshal([]byte(tt.input), &e)
			got := e.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestExSrcTitleUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got ExSrcTitle
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrExSrcTitleUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
