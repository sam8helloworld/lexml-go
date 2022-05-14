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

func TestSubheadwordUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []model.Entity
	}{
		{
			name:  "subheadword内のinline_htmlをUnmarshalできる",
			input: `<subheadword><b>b</b></subheadword>`,
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
			name:  "subheadword内のinline_lexmlをUnmarshalできる",
			input: `<subheadword><alabel>alabel</alabel></subheadword>`,
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
			name:  "subheadword内のpcdataをUnmarshalできる",
			input: `<subheadword>pcdata</subheadword>`,
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
			name:  "subheadword内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<subheadword><b>b</b><alabel>alabel</alabel>pcdata</subheadword>`,
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
			var s Subheadword
			err := xml.Unmarshal([]byte(tt.input), &s)
			got := s.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestSubheadwordUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Subheadword
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrSubheadwordUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
