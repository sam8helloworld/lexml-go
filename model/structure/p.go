package structure

import "encoding/xml"

type P struct {
	Structure
	XMLName xml.Name `xml:"p"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
