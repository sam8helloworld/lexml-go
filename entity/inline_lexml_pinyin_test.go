package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPinyinUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "pinyin内のinline_htmlをUnmarshalできる",
			input: `<pinyin><b>b</b></pinyin>`,
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
			name:  "pinyin内のinline_lexmlをUnmarshalできる",
			input: `<pinyin><alabel>alabel</alabel></pinyin>`,
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
			name:  "pinyin内のpcdataをUnmarshalできる",
			input: `<pinyin>pcdata</pinyin>`,
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
			name:  "pinyin内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<pinyin><b>b</b><alabel>alabel</alabel>pcdata</pinyin>`,
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
			var p Pinyin
			err := xml.Unmarshal([]byte(tt.input), &p)
			got := p.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestPinyinUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got Pinyin
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrPinyinUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
