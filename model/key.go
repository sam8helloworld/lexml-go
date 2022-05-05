package model

import "encoding/xml"

type Key struct {
	XMLName xml.Name `xml:"key"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",innerxml"`
}
