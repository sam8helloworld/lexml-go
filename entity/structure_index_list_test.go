package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIndexListUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "indexlist内のinline_htmlをUnmarshalできる",
			input: `<indexlist><b>b</b></indexlist>`,
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
			name:  "indexlist内のinline_lexmlをUnmarshalできる",
			input: `<indexlist><alabel>alabel</alabel></indexlist>`,
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
			name:  "indexlist内のpcdataをUnmarshalできる",
			input: `<indexlist>pcdata</indexlist>`,
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
			name:  "indexlist内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<indexlist><b>b</b><alabel>alabel</alabel>pcdata</indexlist>`,
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
			var i IndexList
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

func TestIndexListUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got IndexList
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrIndexListUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
