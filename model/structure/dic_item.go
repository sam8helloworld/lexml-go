package model

import (
	"encoding/xml"
	"errors"
)

var ErrDicItemChildrenUnknownElement = errors.New("dic-item have unknown children")

type DicItem struct {
	Structure
	XMLName         xml.Name `xml:"dic-item"`
	ID              string   `xml:"id,attr"`
	Type            string   `xml:"type,attr"`
	Rank            string   `xml:"rank,attr"`
	Level           string   `xml:"level,attr"`
	OrgID           string   `xml:"orgid,attr"`
	PID             string   `xml:"pid,attr"`
	SortKey         string   `xml:"sortkey,attr"`
	Head            Head
	DicItemChildren []DicItemChild `xml:",any"`
}

type DicItemChild struct {
	Type  string
	Value interface{}
}

func (di *DicItem) GetHead() Head {
	return di.Head
}

func (dic *DicItemChild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "meaning":
		var m Meaning
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		dic.Value = m
		dic.Type = start.Name.Local
	case "meaning-group":
		var mg MeaningGroup
		if err := d.DecodeElement(&mg, &start); err != nil {
			return err
		}
		dic.Value = mg
		dic.Type = start.Name.Local
	case "example":
		var e Example
		if err := d.DecodeElement(&e, &start); err != nil {
			return err
		}
		dic.Value = e
		dic.Type = start.Name.Local
	case "example-group":
		var eg ExampleGroup
		if err := d.DecodeElement(&eg, &start); err != nil {
			return err
		}
		dic.Value = eg
		dic.Type = start.Name.Local
	case "subhead":
		var s Subhead
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		dic.Value = s
		dic.Type = start.Name.Local
	case "subheadword":
		var s Subheadword
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		dic.Value = s
		dic.Type = start.Name.Local
	case "index":
		var i Index
		if err := d.DecodeElement(&i, &start); err != nil {
			return err
		}
		dic.Value = i
		dic.Type = start.Name.Local
	case "key":
		var k Key
		if err := d.DecodeElement(&k, &start); err != nil {
			return err
		}
		dic.Value = k
		dic.Type = start.Name.Local
	case "column":
		// TODO: column
	case "div":
		// TODO: dic
	case "p":
		// TODO: p
	case "image":
		// TODO: image
	case "audio":
		// TODO: audio
	case "video":
		// TODO: video
	case "table":
		// TODO: table
	case "replace":
		// TODO: replace
	case "ul":
		// TODO: ul
	case "dl":
		// TODO: dl
	case "memo":
		// TODO: memo
	case "data":
		// TODO: data
	case "snippet":
		var s Snippet
		if err := d.DecodeElement(&s, &start); err != nil {
			return err
		}
		dic.Value = s
		dic.Type = start.Name.Local
	default:
		return ErrDicItemChildrenUnknownElement
	}
	return nil
}
