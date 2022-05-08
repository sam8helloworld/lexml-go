package lexml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
	"github.com/sam8helloworld/lexml-go/model/structure"
)

func TestSuccessUnmarshalDicItem(t *testing.T) {
	input := []byte(`
	<dic-item id="sample01">
	<head>
	<headword>plasma</headword>
	<key>plasma</key>
	<headword>pl&aeacute;zm&schwa;</headword>
	</head>
	</dic-item>`)
	want := structure.DicItem{
		XMLName: xml.Name{Local: "dic-item"},
		ID:      "sample01",
		Head: structure.Head{
			XMLName: xml.Name{Local: "head"},
			HeadChildren: []structure.HeadChild{
				{
					Type:  "headword",
					Value: structure.Headword{XMLName: xml.Name{Local: "headword"}, Value: "plasma"},
				},
				{
					Type:  "key",
					Value: structure.Key{XMLName: xml.Name{Local: "key"}, Value: pcdata.PCDATA{Value: "plasma"}},
				},
				{
					Type:  "headword",
					Value: structure.Headword{XMLName: xml.Name{Local: "headword"}, Value: "pl&aeacute;zm&schwa;"},
				},
			},
		},
	}
	got, err := UnMarshal(input)
	if err != nil {
		t.Fatalf("failed to execute UnMarshal: %v", err)
	}
	if diff := cmp.Diff(*got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
