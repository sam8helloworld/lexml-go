package inline_lexml

import "encoding/xml"

type Gi struct {
	InlineLeXML
	XMLName xml.Name `xml:"gi"`
	Set     string   `xml:"set,attr"`
	Name    string   `xml:"name,attr"`
	Alt     string   `xml:"alt,attr"`
	Value   string   `xml:",chardata"`
}
