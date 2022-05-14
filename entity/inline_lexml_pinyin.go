package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrPinyinUnknownElement = errors.New("pinyin have unknown children")

type Pinyin struct {
	InlineLeXML
	XMLName xml.Name `xml:"pinyin"`
	Value   InnerXML
}

func (p *Pinyin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pp struct {
		InlineLeXML
		XMLName xml.Name `xml:"pinyin"`
		Value   string   `xml:",innerxml"` // (#PCDATA | %inline.html; | %inline.lexml;)*
	}
	if err := d.DecodeElement(&pp, &start); err != nil {
		return errors.Wrap(ErrPinyinUnknownElement, "failed to unmarshal data")
	}
	*p = Pinyin{
		XMLName: xml.Name{Local: "pinyin"},
		Value: InnerXML{
			Value: pp.Value,
		},
	}
	return nil
}
