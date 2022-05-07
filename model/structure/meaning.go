package structure

import "encoding/xml"

type Meaning struct {
	Structure
	XMLName xml.Name `xml:"meaning"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Level   string   `xml:"level,attr"`
	No      string   `xml:"no,attr"`
	Value   string   `xml:",chardata"`
}
