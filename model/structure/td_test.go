package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model"
	"github.com/sam8helloworld/lexml-go/model/inline_html"
	"github.com/sam8helloworld/lexml-go/model/inline_lexml"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestTdUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []model.Entity
	}{
		{
			name:  "td内のinline_htmlをUnmarshalできる",
			input: `<td><b>b</b></td>`,
			want: []model.Entity{
				{
					Type: "b",
					Value: inline_html.SmallB{
						XMLName: xml.Name{Local: "b"},
						Value:   "b",
					},
				},
			},
		},
		{
			name:  "td内のinline_lexmlをUnmarshalできる",
			input: `<td><alabel>alabel</alabel></td>`,
			want: []model.Entity{
				{
					Type: "alabel",
					Value: inline_lexml.ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value:   "alabel",
					},
				},
			},
		},
		{
			name:  "td内のpcdataをUnmarshalできる",
			input: `<td>pcdata</td>`,
			want: []model.Entity{
				{
					Type: "pcdata",
					Value: pcdata.PCDATA{
						Value: "pcdata",
					},
				},
			},
		},
		{
			name:  "td内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<td><b>b</b><alabel>alabel</alabel>pcdata</td>`,
			want: []model.Entity{
				{
					Type: "b",
					Value: inline_html.SmallB{
						XMLName: xml.Name{Local: "b"},
						Value:   "b",
					},
				},
				{
					Type: "alabel",
					Value: inline_lexml.ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value:   "alabel",
					},
				},
				{
					Type: "pcdata",
					Value: pcdata.PCDATA{
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
			var td Td
			err := xml.Unmarshal([]byte(tt.input), &td)
			got := td.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestTdUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Td
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrTdUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
