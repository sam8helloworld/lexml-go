package model

import "encoding/xml"

type Example struct {
	Structure
	XMLName   xml.Name `xml:"example"`
	SubID     string   `xml:"subid,attr"`
	Type      string   `xml:"type,attr"`
	Delimiter string   `xml:"delimiter,attr"`
	Value     string   `xml:",chardata"`
	// TODO: ex-body, ex-trans, ex-src, ex-miscを追加
}
