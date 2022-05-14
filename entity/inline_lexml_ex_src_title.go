package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrExSrcTitleUnknownElement = errors.New("ex-src-title have unknown children")

type ExSrcTitle struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src-title"`
	Org     string   `xml:"org,attr"`
	Value   InnerXML
}

func (e *ExSrcTitle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"ex-src-title"`
		Org     string   `xml:"org,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrExSrcTitleUnknownElement, "failed to unmarshal data")
	}
	*e = ExSrcTitle{
		XMLName: xml.Name{Local: "ex-src-author"},
		Org:     ee.Org,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
