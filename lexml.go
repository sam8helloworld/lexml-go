package lexml

import (
	"bytes"
	"encoding/xml"

	"github.com/sam8helloworld/lexml-go/model"
)

func UnMarshal(xmldoc []byte) (*model.DicItem, error) {
	di := model.DicItem{}
	dec := xml.NewDecoder(bytes.NewReader(xmldoc))
	dec.Strict = false
	err := dec.Decode(&di)
	if err != nil {
		return nil, err
	}
	return &di, nil
}
