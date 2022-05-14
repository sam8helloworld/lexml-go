package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrMemoUnknownElement = errors.New("memo have unknown children")

type Memo struct {
	Structure
	XMLName xml.Name `xml:"memo"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (m *Memo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var mm struct {
		Structure
		XMLName xml.Name `xml:"memo"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&mm, &start); err != nil {
		return errors.Wrap(ErrMemoUnknownElement, "failed to unmarshal data")
	}
	*m = Memo{
		XMLName: xml.Name{Local: "memo"},
		Type:    mm.Type,
		Value: model.InnerXML{
			Value: mm.Value,
		},
	}
	return nil
}
