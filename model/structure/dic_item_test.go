package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
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
					Value: model.InnerXML{
						Value: "意味",
					},
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
					Value: model.InnerXML{
						Value: "サブヘッドワード",
					},
				},
			},
		},
		{
			name:  "indexをUnmarshalできる",
			input: `<index></index>`,
			want: DicItemChild{
				Type: "index",
				Value: Index{
					XMLName: xml.Name{Local: "index"},
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
					Value:   pcdata.PCDATA{Value: "key"},
				},
			},
		},
		{
			name:  "columnをUnmarshalできる",
			input: `<column></column>`,
			want: DicItemChild{
				Type: "column",
				Value: Column{
					XMLName: xml.Name{Local: "column"},
				},
			},
		},
		{
			name:  "divをUnmarshalできる",
			input: `<div></div>`,
			want: DicItemChild{
				Type: "div",
				Value: Div{
					XMLName: xml.Name{Local: "div"},
				},
			},
		},
		{
			name:  "pをUnmarshalできる",
			input: `<p>p</p>`,
			want: DicItemChild{
				Type: "p",
				Value: P{
					XMLName: xml.Name{Local: "p"},
					Value: model.InnerXML{
						Value: "p",
					},
				},
			},
		},
		{
			name:  "imageをUnmarshalできる",
			input: `<image></image>`,
			want: DicItemChild{
				Type: "image",
				Value: Image{
					XMLName: xml.Name{Local: "image"},
				},
			},
		},
		{
			name:  "audioをUnmarshalできる",
			input: `<audio></audio>`,
			want: DicItemChild{
				Type: "audio",
				Value: Audio{
					XMLName: xml.Name{Local: "audio"},
				},
			},
		},
		{
			name:  "videoをUnmarshalできる",
			input: `<video></video>`,
			want: DicItemChild{
				Type: "video",
				Value: Video{
					XMLName: xml.Name{Local: "video"},
				},
			},
		},
		{
			name:  "tableをUnmarshalできる",
			input: `<table></table>`,
			want: DicItemChild{
				Type: "table",
				Value: Table{
					XMLName: xml.Name{Local: "table"},
				},
			},
		},
		{
			name:  "replaceをUnmarshalできる",
			input: `<replace></replace>`,
			want: DicItemChild{
				Type: "replace",
				Value: Replace{
					XMLName: xml.Name{Local: "replace"},
				},
			},
		},
		{
			name:  "ulをUnmarshalできる",
			input: `<ul></ul>`,
			want: DicItemChild{
				Type: "ul",
				Value: Ul{
					XMLName: xml.Name{Local: "ul"},
				},
			},
		},
		{
			name:  "dlをUnmarshalできる",
			input: `<dl></dl>`,
			want: DicItemChild{
				Type: "dl",
				Value: Dl{
					XMLName: xml.Name{Local: "dl"},
				},
			},
		},
		{
			name:  "memoをUnmarshalできる",
			input: `<memo>memo</memo>`,
			want: DicItemChild{
				Type: "memo",
				Value: Memo{
					XMLName: xml.Name{Local: "memo"},
					Value: model.InnerXML{
						Value: "memo",
					},
				},
			},
		},
		{
			name:  "dataをUnmarshalできる",
			input: `<data>data</data>`,
			want: DicItemChild{
				Type: "data",
				Value: Data{
					XMLName: xml.Name{Local: "data"},
					Value: model.InnerXML{
						Value: "data",
					},
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
					Value:   pcdata.PCDATA{Value: "snippet"},
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
