package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "i内のinline_htmlをUnmarshalできる",
			input: `<i><b>b</b></i>`,
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
			name:  "i内のinline_lexmlをUnmarshalできる",
			input: `<i><alabel>alabel</alabel></i>`,
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
			name:  "i内のpcdataをUnmarshalできる",
			input: `<i>pcdata</i>`,
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
			name:  "i内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<i><b>b</b><alabel>alabel</alabel>pcdata</i>`,
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
			var i SmallI
			err := xml.Unmarshal([]byte(tt.input), &i)
			got := i.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestIUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var gotSmallI SmallI
	err := xml.Unmarshal(input, &gotSmallI)
	if !errors.Is(err, ErrIUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	var gotLargeI LargeI
	err = xml.Unmarshal(input, &gotLargeI)
	if !errors.Is(err, ErrIUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
