package model

import "encoding/xml"

type Data struct {
	Structure
	XMLName xml.Name `xml:"data"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}