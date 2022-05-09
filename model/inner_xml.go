package model

import (
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/sam8helloworld/lexml-go/model/inline_html"
	"github.com/sam8helloworld/lexml-go/model/inline_lexml"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

type InnerXML struct {
	Value string
}

type Entity struct {
	Type  string
	Value interface{}
}

func (i *InnerXML) StructuredValue() []Entity {
	tags := []string{}
	for _, t := range append(inline_html.Tags, inline_lexml.Tags...) {
		tg := "<" + t + ">.*</" + t + ">"
		tags = append(tags, tg)
	}
	tagsRegStr := "(" + strings.Join(tags, "|") + ")"
	reg := regexp.MustCompile(tagsRegStr)
	// FIXME: 流石に独自に@@@を区切り文字として入れるのは良くない
	rep := reg.ReplaceAllString(i.Value, "@@@$1@@@")
	entities := strings.Split(rep, "@@@")
	var models []Entity
	for _, entity := range entities {
		// FIXME: entitiesにから文字列の要素が入ってしまう
		if entity == "" {
			continue
		}
		reg := regexp.MustCompile(`<(.+)>(.+)</(.+)>`)
		submatch := reg.FindStringSubmatch(entity)
		var model Entity
		// タグがない時
		if len(submatch) == 0 {
			model.Type = "pcdata"
			model.Value = pcdata.PCDATA{
				Value: entity,
			}
			models = append(models, model)
			continue
		}
		model.Type = submatch[1]
		if inline_html.IsInlineHTML(model.Type) {
			var ent inline_html.SmallB
			xml.Unmarshal([]byte(entity), &ent)
			model.Value = ent
		}
		if inline_lexml.IsInlineLeXML(model.Type) {
			var ent inline_lexml.ALabel
			xml.Unmarshal([]byte(entity), &ent)
			model.Value = ent
		}
		models = append(models, model)
	}
	return models
}
