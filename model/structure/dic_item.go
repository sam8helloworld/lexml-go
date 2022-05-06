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
		var c Column
		if err := d.DecodeElement(&c, &start); err != nil {
			return err
		}
		dic.Value = c
		dic.Type = start.Name.Local
	case "div":
		var div Div
		if err := d.DecodeElement(&div, &start); err != nil {
			return err
		}
		dic.Value = div
		dic.Type = start.Name.Local
	case "p":
		var p P
		if err := d.DecodeElement(&p, &start); err != nil {
			return err
		}
		dic.Value = p
		dic.Type = start.Name.Local
	case "image":
		var i Image
		if err := d.DecodeElement(&i, &start); err != nil {
			return err
		}
		dic.Value = i
		dic.Type = start.Name.Local
	case "audio":
		var a Audio
		if err := d.DecodeElement(&a, &start); err != nil {
			return err
		}
		dic.Value = a
		dic.Type = start.Name.Local
	case "video":
		var v Video
		if err := d.DecodeElement(&v, &start); err != nil {
			return err
		}
		dic.Value = v
		dic.Type = start.Name.Local
	case "table":
		var t Table
		if err := d.DecodeElement(&t, &start); err != nil {
			return err
		}
		dic.Value = t
		dic.Type = start.Name.Local
	case "replace":
		var r Replace
		if err := d.DecodeElement(&r, &start); err != nil {
			return err
		}
		dic.Value = r
		dic.Type = start.Name.Local
	case "ul":
		var u Ul
		if err := d.DecodeElement(&u, &start); err != nil {
			return err
		}
		dic.Value = u
		dic.Type = start.Name.Local
	case "dl":
		var dl Dl
		if err := d.DecodeElement(&dl, &start); err != nil {
			return err
		}
		dic.Value = dl
		dic.Type = start.Name.Local
	case "memo":
		var m Memo
		if err := d.DecodeElement(&m, &start); err != nil {
			return err
		}
		dic.Value = m
		dic.Type = start.Name.Local
	case "data":
		var data Data
		if err := d.DecodeElement(&data, &start); err != nil {
			return err
		}
		dic.Value = data
		dic.Type = start.Name.Local
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
