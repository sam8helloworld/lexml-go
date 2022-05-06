package model

import "encoding/xml"

type Td struct {
	Structure
	XMLName xml.Name `xml:"td"`
	ColSpan string   `xml:"colspan,attr"`
	RowSpan string   `xml:"rowspan,attr"`
	Value   string   `xml:",chardata"`
}
