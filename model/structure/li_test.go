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

func TestLiUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []model.Entity
	}{
		{
			name:  "li内のinline_htmlをUnmarshalできる",
			input: `<li><b>b</b></li>`,
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
			name:  "li内のinline_lexmlをUnmarshalできる",
			input: `<li><alabel>alabel</alabel></li>`,
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
			name:  "li内のpcdataをUnmarshalできる",
			input: `<li>pcdata</li>`,
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
			name:  "li内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<li><b>b</b><alabel>alabel</alabel>pcdata</li>`,
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
			var li Li
			err := xml.Unmarshal([]byte(tt.input), &li)
			got := li.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestLiUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Li
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrLiUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
