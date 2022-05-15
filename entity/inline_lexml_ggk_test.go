package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGgkUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "ggk内のinline_htmlをUnmarshalできる",
			input: `<ggk><b>b</b></ggk>`,
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
			name:  "ggk内のinline_lexmlをUnmarshalできる",
			input: `<ggk><alabel>alabel</alabel></ggk>`,
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
			name:  "ggk内のpcdataをUnmarshalできる",
			input: `<ggk>pcdata</ggk>`,
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
			name:  "ggk内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<ggk><b>b</b><alabel>alabel</alabel>pcdata</ggk>`,
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
			var g Ggk
			err := xml.Unmarshal([]byte(tt.input), &g)
			got := g.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestGgkUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Ggk
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrGgkUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
