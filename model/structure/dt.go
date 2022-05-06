package model

import "encoding/xml"

type Dt struct {
	Structure
	XMLName xml.Name `xml:"dt"`
	SubID   string   `xml:"subid,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}