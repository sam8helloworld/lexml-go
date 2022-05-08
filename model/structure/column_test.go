package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestColumnUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ColumnChild
	}{
		{
			name:  "titleをUnmarshalできる",
			input: `<title>タイトル</title>`,
			want: ColumnChild{
				Type: "title",
				Value: Title{
					XMLName: xml.Name{Local: "title"},
					Value:   "タイトル",
				},
			},
		},
		{
			name:  "keyをUnmarshalできる",
			input: `<key>key</key>`,
			want: ColumnChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   pcdata.PCDATA{Value: "key"},
				},
			},
		},
		{
			name:  "meaningをUnmarshalできる",
			input: `<meaning>意味</meaning>`,
			want: ColumnChild{
				Type: "meaning",
				Value: Meaning{
					XMLName: xml.Name{Local: "meaning"},
					Value:   "意味",
				},
			},
		},
		{
			name:  "exampleをUnmarshalできる",
			input: `<example>例</example>`,
			want: ColumnChild{
				Type: "example",
				Value: Example{
					XMLName: xml.Name{Local: "example"},
					Value:   "例",
				},
			},
		},
		{
			name:  "subheadをUnmarshalできる",
			input: `<subhead></subhead>`,
			want: ColumnChild{
				Type: "subhead",
				Value: Subhead{
					XMLName: xml.Name{Local: "subhead"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got ColumnChild
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

func TestColumnUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got ColumnChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrColumnChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
