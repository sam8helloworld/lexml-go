package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExSrcAuthorUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "ex-src-author内のinline_htmlをUnmarshalできる",
			input: `<ex-src-author><b>b</b></ex-src-author>`,
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
			name:  "ex-src-author内のinline_lexmlをUnmarshalできる",
			input: `<ex-src-author><alabel>alabel</alabel></ex-src-author>`,
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
			name:  "ex-src-author内のpcdataをUnmarshalできる",
			input: `<ex-src-author>pcdata</ex-src-author>`,
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
			name:  "ex-src-author内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<ex-src-author><b>b</b><alabel>alabel</alabel>pcdata</ex-src-author>`,
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
			var e ExSrcAuthor
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

func TestExSrcAuthorUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got ExSrcAuthor
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrExSrcAuthorUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
