package structure

import "encoding/xml"

type Caption struct {
	Structure
	XMLName xml.Name `xml:"caption"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
