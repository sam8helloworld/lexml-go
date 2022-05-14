package structure

import (
	"encoding/xml"

	"github.com/pkg/errors"
	"github.com/sam8helloworld/lexml-go/model"
)

var ErrCaptionUnknownElement = errors.New("caption have unknown children")

type Caption struct {
	Structure
	XMLName xml.Name `xml:"caption"`
	Type    string   `xml:"type,attr"`
	Value   model.InnerXML
}

func (c *Caption) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var caption struct {
		Structure
		XMLName xml.Name `xml:"caption"`
		Type    string   `xml:"type,attr"`
		Value   string   `xml:",innerxml"`
	}
	if err := d.DecodeElement(&caption, &start); err != nil {
		return errors.Wrap(ErrCaptionUnknownElement, "failed to unmarshal caption")
	}
	*c = Caption{
		XMLName: xml.Name{Local: "caption"},
		Type:    caption.Type,
		Value: model.InnerXML{
			Value: caption.Value,
		},
	}
	return nil
}
