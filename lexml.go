package lexml

import (
	"bytes"
	"encoding/xml"

	"github.com/sam8helloworld/lexml-go/entity"
)

func UnMarshal(xmldoc []byte) (*entity.DicItem, error) {
	di := entity.DicItem{}
	dec := xml.NewDecoder(bytes.NewReader(xmldoc))
	dec.Strict = false
	err := dec.Decode(&di)
	if err != nil {
		return nil, err
	}
	return &di, nil
}
