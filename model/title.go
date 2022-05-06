package model

import "encoding/xml"

type Title struct {
	XMLName xml.Name `xml:"title"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
