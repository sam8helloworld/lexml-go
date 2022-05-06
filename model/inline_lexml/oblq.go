package model

import "encoding/xml"

type Oblq struct {
	InlineLeXML
	XMLName xml.Name `xml:"oblq"`
	Value   string   `xml:",chardata"`
}
