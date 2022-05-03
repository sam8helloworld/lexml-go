package lexml

import (
	"encoding/xml"

	"github.com/sam8helloworld/lexml-go/model"
)

func UnMarshal(xmldoc []byte) (*model.DicItem, error) {
	di := model.DicItem{}
	err := xml.Unmarshal(xmldoc, &di)
	if err != nil {
		return nil, err
	}
	return &di, nil
}
