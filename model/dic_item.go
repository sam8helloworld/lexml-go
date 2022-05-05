package model

import "encoding/xml"

type DicItem struct {
	XMLName xml.Name `xml:"dic-item"`
	ID      string   `xml:"id,attr"`
	Type    string   `xml:"type,attr"`
	Rank    string   `xml:"rank,attr"`
	Level   string   `xml:"level,attr"`
	OrgID   string   `xml:"orgid,attr"`
	PID     string   `xml:"pid,attr"`
	SortKey string   `xml:"sortkey,attr"`
	Head    Head
}

func (di *DicItem) GetHead() Head {
	return di.Head
}
