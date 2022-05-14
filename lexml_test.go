package lexml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/entity"
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
	want := entity.DicItem{
		XMLName: xml.Name{Local: "dic-item"},
		ID:      "sample01",
		Head: entity.Head{
			XMLName: xml.Name{Local: "head"},
			HeadChildren: []entity.HeadChild{
				{
					Type: "headword",
					Value: entity.Headword{
						XMLName: xml.Name{Local: "headword"},
						Value: entity.InnerXML{
							Value: "plasma",
						},
					},
				},
				{
					Type:  "key",
					Value: entity.Key{XMLName: xml.Name{Local: "key"}, Value: entity.PCDATA{Value: "plasma"}},
				},
				{
					Type: "headword",
					Value: entity.Headword{
						XMLName: xml.Name{Local: "headword"},
						Value: entity.InnerXML{
							Value: "pl&aeacute;zm&schwa;",
						},
					},
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
