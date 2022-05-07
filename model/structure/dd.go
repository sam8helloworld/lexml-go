package structure

import "encoding/xml"

type Dd struct {
	Structure
	XMLName xml.Name `xml:"dd"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
