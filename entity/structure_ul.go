package entity

import "encoding/xml"

type Ul struct {
	Structure
	XMLName xml.Name `xml:"ul"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   []Li     `xml:"li"` // (li)+
}
