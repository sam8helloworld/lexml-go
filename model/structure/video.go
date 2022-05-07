package structure

import "encoding/xml"

type Video struct {
	Structure
	XMLName  xml.Name `xml:"video"`
	Type     string   `xml:"type,attr"`
	Src      string   `xml:"src,attr"`
	MimeType string   `xml:"mime-type,attr"`
}
