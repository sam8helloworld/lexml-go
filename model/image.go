package model

import "encoding/xml"

type Image struct {
	XMLName  xml.Name `xml:"image"`
	Type     string   `xml:"type,attr"`
	Src      string   `xml:"src,attr"`
	MimeType string   `xml:"mime-type,attr"`
}
