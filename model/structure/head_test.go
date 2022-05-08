package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestHeadUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  HeadChild
	}{
		{
			name:  "headwordをUnmarshalできる",
			input: `<headword>plasma</headword>`,
			want: HeadChild{
				Type: "headword",
				Value: Headword{
					XMLName: xml.Name{Local: "headword"},
					Value:   "plasma",
				},
			},
		},
		{
			name:  "keyをUnmarshalできる",
			input: `<key>key</key>`,
			want: HeadChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   pcdata.PCDATA{Value: "key"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got HeadChild
			err := xml.Unmarshal([]byte(tt.input), &got)
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestHeadUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got HeadChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrHeadChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
