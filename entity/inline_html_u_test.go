package entity

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Entity
	}{
		{
			name:  "u内のinline_htmlをUnmarshalできる",
			input: `<u><b>b</b></u>`,
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
			name:  "u内のinline_lexmlをUnmarshalできる",
			input: `<u><alabel>alabel</alabel></u>`,
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
			name:  "u内のpcdataをUnmarshalできる",
			input: `<u>pcdata</u>`,
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
			name:  "u内のinline_html, inline_lexml, pcdataを同時にUnmarshalできる",
			input: `<u><b>b</b><alabel>alabel</alabel>pcdata</u>`,
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
			var u SmallU
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

func TestUUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var gotSmallU SmallU
	err := xml.Unmarshal(input, &gotSmallU)
	if !errors.Is(err, ErrUUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	var gotLargeU LargeU
	err = xml.Unmarshal(input, &gotLargeU)
	if !errors.Is(err, ErrUUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
