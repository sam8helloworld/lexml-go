package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUdUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "ud内のinline_htmlをUnmarshalできる",
			input: `<ud><b>b</b></ud>`,
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
			name:  "ud内のinline_lexmlをUnmarshalできる",
			input: `<ud><alabel>alabel</alabel></ud>`,
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
			name:  "ud内のpcdataをUnmarshalできる",
			input: `<ud>pcdata</ud>`,
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
			name:  "ud内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<ud><b>b</b><alabel>alabel</alabel>pcdata</ud>`,
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
			var u Ud
			err := xml.Unmarshal([]byte(tt.input), &u)
			got := u.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestUdUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Ud
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrUdUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
