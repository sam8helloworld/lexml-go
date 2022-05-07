package structure

import "encoding/xml"

type Key struct {
	Structure
	XMLName xml.Name `xml:"key"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
