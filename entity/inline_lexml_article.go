package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrArticleUnknownElement = errors.New("article have unknown children")

type Article struct {
	InlineLeXML
	XMLName xml.Name `xml:"article"`
	Value   InnerXML
}

func (a *Article) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var aa struct {
		InlineLeXML
		XMLName xml.Name `xml:"article"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&aa, &start); err != nil {
		return errors.Wrap(ErrArticleUnknownElement, "failed to unmarshal data")
	}
	*a = Article{
		XMLName: xml.Name{Local: "article"},
		Value: InnerXML{
			Value: aa.Value,
		},
	}
	return nil
}
