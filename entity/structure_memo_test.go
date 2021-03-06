package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMemoUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "memo内のinline_htmlをUnmarshalできる",
			input: `<memo><b>b</b></memo>`,
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
			name:  "memo内のinline_lexmlをUnmarshalできる",
			input: `<memo><alabel>alabel</alabel></memo>`,
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
			name:  "memo内のpcdataをUnmarshalできる",
			input: `<memo>pcdata</memo>`,
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
			name:  "memo内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<memo><b>b</b><alabel>alabel</alabel>pcdata</memo>`,
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
			var m Memo
			err := xml.Unmarshal([]byte(tt.input), &m)
			got := m.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestMemoUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Memo
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrMemoUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
