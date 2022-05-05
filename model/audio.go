package model

import "encoding/xml"

type Audio struct {
	XMLName  xml.Name `xml:"audio"`
	Type     string   `xml:"type,attr"`
	Src      string   `xml:"src,attr"`
	MimeType string   `xml:"mime-type,attr"`
}
