package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDivUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  DivChild
	}{
		{
			name:  "titleをUnmarshalできる",
			input: `<title>タイトル</title>`,
			want: DivChild{
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
			want: DivChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   "key",
				},
			},
		},
		{
			name:  "meaningをUnmarshalできる",
			input: `<meaning>意味</meaning>`,
			want: DivChild{
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
			want: DivChild{
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
			want: DivChild{
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
			var got DivChild
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

func TestDivUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got DivChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrDivChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
