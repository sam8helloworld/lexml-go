package entity

import "encoding/xml"

type Mlg struct {
	InlineLeXML
	XMLName xml.Name `xml:"mlg"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml; | mlgbr)*
}
