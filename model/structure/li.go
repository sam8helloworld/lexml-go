package structure

import "encoding/xml"

type Li struct {
	Structure
	XMLName xml.Name `xml:"li"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
