package model

import "encoding/xml"

type Key struct {
	XMLName xml.Name `xml:"key"`
	KeyType string   `xml:"type,attr"`
	Value   string   `xml:",innerxml"`
}
