package model

import "encoding/xml"

type pos struct {
	InlineLeXML
	XMLName xml.Name `xml:"pos"`
	Value   string   `xml:",chardata"`
}

type Pos struct {
	InlineLeXML
	XMLName xml.Name `xml:"Pos"`
	Value   string   `xml:",chardata"`
}
