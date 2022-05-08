package lexml

import (
	"bytes"
	"encoding/xml"

	"github.com/sam8helloworld/lexml-go/model/structure"
)

func UnMarshal(xmldoc []byte) (*structure.DicItem, error) {
	di := structure.DicItem{}
	dec := xml.NewDecoder(bytes.NewReader(xmldoc))
	dec.Strict = false
	err := dec.Decode(&di)
	if err != nil {
		return nil, err
	}
	return &di, nil
}
