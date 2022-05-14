package entity

import "encoding/xml"

type SmallBR struct {
	InlineHTML
	XMLName xml.Name `xml:"br"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
}

type LargeBR struct {
	InlineHTML
	XMLName xml.Name `xml:"BR"`
	Type    string   `xml:"type,attr"`
	Class   string   `xml:"class,attr"`
}
