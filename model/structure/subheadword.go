package structure

import "encoding/xml"

type Subheadword struct {
	Structure
	XMLName   xml.Name `xml:"subheadword"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     string   `xml:",chardata"`
}
