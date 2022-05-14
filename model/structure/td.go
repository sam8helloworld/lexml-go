package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrTdUnknownElement = errors.New("td have unknown children")

type Td struct {
	Structure
	XMLName xml.Name `xml:"td"`
	ColSpan string   `xml:"colspan,attr"`
	RowSpan string   `xml:"rowspan,attr"`
	Value   model.InnerXML
}

func (t *Td) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var tt struct {
		Structure
		XMLName xml.Name `xml:"td"`
		ColSpan string   `xml:"colspan,attr"`
		RowSpan string   `xml:"rowspan,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&tt, &start); err != nil {
		return errors.Wrap(ErrTdUnknownElement, "failed to unmarshal data")
	}
	*t = Td{
		XMLName: xml.Name{Local: "td"},
		ColSpan: tt.ColSpan,
		RowSpan: tt.RowSpan,
		Value: model.InnerXML{
			Value: tt.Value,
		},
	}
	return nil
}
