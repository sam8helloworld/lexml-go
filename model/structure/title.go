package structure

import "encoding/xml"

type Title struct {
	Structure
	XMLName xml.Name `xml:"title"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
