package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDdUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "dd内のinline_htmlをUnmarshalできる",
			input: `<dd><b>b</b></dd>`,
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
			name:  "dd内のinline_lexmlをUnmarshalできる",
			input: `<dd><alabel>alabel</alabel></dd>`,
			want: []Entity{
				{
					Type: "alabel",
					Value: ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value:   "alabel",
					},
				},
			},
		},
		{
			name:  "dd内のpcdataをUnmarshalできる",
			input: `<dd>pcdata</dd>`,
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
			name:  "dd内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<dd><b>b</b><alabel>alabel</alabel>pcdata</dd>`,
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
						Value:   "alabel",
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
			var d Dd
			err := xml.Unmarshal([]byte(tt.input), &d)
			got := d.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestDdUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Dd
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrDdUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
