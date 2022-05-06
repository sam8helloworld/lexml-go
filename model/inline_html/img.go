package model

import "encoding/xml"

type Img struct {
	InlineHTML
	XMLName  xml.Name `xml:"img"`
	Type     string   `xml:"type,attr"`
	Src      string   `xml:"src,attr"`
	Alt      string   `xml:"alt,attr"`
	Title    string   `xml:"title,attr"`
	MimeType string   `xml:"mime-type,attr"`
	Value    string   `xml:",chardata"`
}
