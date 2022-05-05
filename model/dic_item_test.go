package model

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDicItemUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  DicItemChild
	}{
		{
			name:  "meaningをUnmarshalできる",
			input: `<meaning>意味</meaning>`,
			want: DicItemChild{
				Type: "meaning",
				Value: Meaning{
					XMLName: xml.Name{Local: "meaning"},
					Value:   "意味",
				},
			},
		},
		{
			name:  "meaning-groupをUnmarshalできる",
			input: `<meaning-group></meaning-group>`,
			want: DicItemChild{
				Type: "meaning-group",
				Value: MeaningGroup{
					XMLName: xml.Name{Local: "meaning-group"},
				},
			},
		},
		{
			name:  "exampleをUnmarshalできる",
			input: `<example>例</example>`,
			want: DicItemChild{
				Type: "example",
				Value: Example{
					XMLName: xml.Name{Local: "example"},
					Value:   "例",
				},
			},
		},
		{
			name:  "example-groupをUnmarshalできる",
			input: `<example-group></example-group>`,
			want: DicItemChild{
				Type: "example-group",
				Value: ExampleGroup{
					XMLName: xml.Name{Local: "example-group"},
				},
			},
		},
		{
			name:  "subheadをUnmarshalできる",
			input: `<subhead></subhead>`,
			want: DicItemChild{
				Type: "subhead",
				Value: Subhead{
					XMLName: xml.Name{Local: "subhead"},
				},
			},
		},
		{
			name:  "subheadwordをUnmarshalできる",
			input: `<subheadword>サブヘッドワード</subheadword>`,
			want: DicItemChild{
				Type: "subheadword",
				Value: Subheadword{
					XMLName: xml.Name{Local: "subheadword"},
					Value:   "サブヘッドワード",
				},
			},
		},
		{
			name:  "keyをUnmarshalできる",
			input: `<key>key</key>`,
			want: DicItemChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   "key",
				},
			},
		},
		{
			name:  "snippetをUnmarshalできる",
			input: `<snippet>snippet</snippet>`,
			want: DicItemChild{
				Type: "snippet",
				Value: Snippet{
					XMLName: xml.Name{Local: "snippet"},
					Value:   "snippet",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got DicItemChild
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

func TestDicItemUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got DicItemChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrDicItemChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
