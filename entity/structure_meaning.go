package entity

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

var ErrMeaningUnknownElement = errors.New("meaning have unknown children")

type Meaning struct {
	Structure
	XMLName xml.Name `xml:"meaning"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Level   string   `xml:"level,attr"`
	No      string   `xml:"no,attr"`
	Value   InnerXML
}

func (m *Meaning) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var meaning struct {
		Structure
		XMLName xml.Name `xml:"meaning"`
		SubID   string   `xml:"subid,attr"`
		Type    string   `xml:"type,attr"`
		Level   string   `xml:"level,attr"`
		No      string   `xml:"no,attr"`
		Value   string   `xml:",innerxml"`
	}
	if err := d.DecodeElement(&meaning, &start); err != nil {
		return errors.Wrap(ErrMeaningUnknownElement, "failed to unmarshal meaning")
	}
	*m = Meaning{
		XMLName: xml.Name{Local: "meaning"},
		SubID:   meaning.SubID,
		Type:    meaning.Type,
		Level:   meaning.Level,
		No:      meaning.No,
		Value: InnerXML{
			Value: meaning.Value,
		},
	}
	return nil
}
