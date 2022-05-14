package entity

import (
	"encoding/xml"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type InnerXML struct {
	Value string
}

type Entity struct {
	Type  string
	Value interface{}
}

// FIXME: ロジックが適当
func (i *InnerXML) StructuredValue() []Entity {
	node, _ := html.Parse(strings.NewReader(i.Value))
	var nodes []Entity
	c := node.FirstChild
	for c != nil {
		if c.Type == html.TextNode {
			nodes = append(nodes, Entity{
				Type: "pcdata",
				Value: PCDATA{
					Value: c.Data,
				},
			})
		}
		if c.Type == html.ElementNode {
			if !IsInlineHTML(c.Data) && !IsInlineLeXML(c.Data) {
				if c.FirstChild == nil {
					c = c.NextSibling
				} else {
					c = c.FirstChild
				}
				continue
			}
			if c.Data == "b" {
				nodes = append(nodes, Entity{
					Type: c.Data,
					Value: SmallB{
						XMLName: xml.Name{Local: "b"},
						Value:   htmlquery.OutputHTML(c, false),
					},
				})
			}
			if c.Data == "alabel" {
				nodes = append(nodes, Entity{
					Type: c.Data,
					Value: ALabel{
						XMLName: xml.Name{Local: "alabel"},
						Value: InnerXML{
							Value: htmlquery.OutputHTML(c, false),
						},
					},
				})
			}
			c = c.NextSibling
			continue
		}
		c = c.NextSibling
	}
	return nodes
}
