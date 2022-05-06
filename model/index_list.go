package model

import "encoding/xml"

type IndexList struct {
	XMLName xml.Name `xml:"indexlist"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}
