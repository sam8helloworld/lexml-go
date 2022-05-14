package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExBodyUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "ex-body内のinline_htmlをUnmarshalできる",
			input: `<ex-body><b>b</b></ex-body>`,
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
			name:  "ex-body内のinline_lexmlをUnmarshalできる",
			input: `<ex-body><alabel>alabel</alabel></ex-body>`,
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
			name:  "ex-body内のpcdataをUnmarshalできる",
			input: `<ex-body>pcdata</ex-body>`,
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
			name:  "ex-body内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<ex-body><b>b</b><alabel>alabel</alabel>pcdata</ex-body>`,
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
			var e ExBody
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

func TestExBodyUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got ExBody
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrExBodyUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
