package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEmUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "em内のinline_htmlをUnmarshalできる",
			input: `<em><b>b</b></em>`,
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
			name:  "em内のinline_lexmlをUnmarshalできる",
			input: `<em><alabel>alabel</alabel></em>`,
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
			name:  "em内のpcdataをUnmarshalできる",
			input: `<em>pcdata</em>`,
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
			name:  "em内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<em><b>b</b><alabel>alabel</alabel>pcdata</em>`,
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
			var e SmallEM
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

func TestEmUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var gotSmallEm SmallEM
	err := xml.Unmarshal(input, &gotSmallEm)
	if !errors.Is(err, ErrEmUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	var gotLargeEm LargeEM
	err = xml.Unmarshal(input, &gotLargeEm)
	if !errors.Is(err, ErrEmUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
