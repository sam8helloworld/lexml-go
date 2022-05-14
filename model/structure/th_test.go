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

func TestThUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []model.Entity
	}{
		{
			name:  "th内のinline_htmlをUnmarshalできる",
			input: `<th><b>b</b></th>`,
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
			name:  "th内のinline_lexmlをUnmarshalできる",
			input: `<th><alabel>alabel</alabel></th>`,
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
			name:  "th内のpcdataをUnmarshalできる",
			input: `<th>pcdata</th>`,
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
			name:  "th内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<th><b>b</b><alabel>alabel</alabel>pcdata</th>`,
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
			var th Th
			err := xml.Unmarshal([]byte(tt.input), &th)
			got := th.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestThUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Th
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrThUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
