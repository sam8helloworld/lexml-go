package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrExSrcAuthorUnknownElement = errors.New("ex-src-author have unknown children")

type ExSrcAuthor struct {
	InlineLeXML
	XMLName xml.Name `xml:"ex-src-author"`
	Org     string   `xml:"org,attr"`
	Value   InnerXML
}

func (e *ExSrcAuthor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ee struct {
		InlineLeXML
		XMLName xml.Name `xml:"ex-src-author"`
		Org     string   `xml:"org,attr"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&ee, &start); err != nil {
		return errors.Wrap(ErrExSrcAuthorUnknownElement, "failed to unmarshal data")
	}
	*e = ExSrcAuthor{
		XMLName: xml.Name{Local: "ex-src-author"},
		Org:     ee.Org,
		Value: InnerXML{
			Value: ee.Value,
		},
	}
	return nil
}
