package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKKaeriUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "kkaeri内のinline_htmlをUnmarshalできる",
			input: `<kkaeri><b>b</b></kkaeri>`,
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
			name:  "kkaeri内のinline_lexmlをUnmarshalできる",
			input: `<kkaeri><alabel>alabel</alabel></kkaeri>`,
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
			name:  "kkaeri内のpcdataをUnmarshalできる",
			input: `<kkaeri>pcdata</kkaeri>`,
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
			name:  "kkaeri内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<kkaeri><b>b</b><alabel>alabel</alabel>pcdata</kkaeri>`,
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
			var k KKaeri
			err := xml.Unmarshal([]byte(tt.input), &k)
			got := k.Value.StructuredValue()
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestKKaeriUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got KKaeri
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrKKaeriUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
