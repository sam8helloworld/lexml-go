package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKOkuriUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "kokuri内のinline_htmlをUnmarshalできる",
			input: `<kokuri><b>b</b></kokuri>`,
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
			name:  "kokuri内のinline_lexmlをUnmarshalできる",
			input: `<kokuri><alabel>alabel</alabel></kokuri>`,
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
			name:  "kokuri内のpcdataをUnmarshalできる",
			input: `<kokuri>pcdata</kokuri>`,
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
			name:  "kokuri内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<kokuri><b>b</b><alabel>alabel</alabel>pcdata</kokuri>`,
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
			var k KOkuri
			err := xml.Unmarshal([]byte(tt.input), &k)
			got := k.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestKOkuriUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got KOkuri
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrKOkuriUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
