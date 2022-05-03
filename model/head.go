package model

import "encoding/xml"

type Head struct {
	XMLName xml.Name `xml:"head"`
	// headwordとkeyを持つ
}
