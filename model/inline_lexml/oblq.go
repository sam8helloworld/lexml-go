package inline_lexml

import "encoding/xml"

type Oblq struct {
	InlineLeXML
	XMLName xml.Name `xml:"oblq"`
	Value   string   `xml:",chardata"` // (#PCDATA | %inline.html; | %inline.lexml;)*
}
