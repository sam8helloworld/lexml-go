package lexml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model"
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
	want := model.DicItem{
		XMLName: xml.Name{Local: "dic-item"},
		ID:      "sample01",
		Head: model.Head{
			XMLName: xml.Name{Local: "head"},
			HeadChildren: []model.HeadChild{
				{
					Type:  "headword",
					Value: model.Headword{XMLName: xml.Name{Local: "headword"}, Value: "plasma"},
				},
				{
					Type:  "key",
					Value: model.Key{XMLName: xml.Name{Local: "key"}, Value: "plasma"},
				},
				{
					Type:  "headword",
					Value: model.Headword{XMLName: xml.Name{Local: "headword"}, Value: "pl&aeacute;zm&schwa;"},
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
