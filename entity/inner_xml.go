package entity

import (
	"encoding/xml"
	"regexp"
	"strings"
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
	for _, t := range append(InlineHtmlTags, InlineLeXMLTags...) {
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
			model.Value = PCDATA{
				Value: entity,
			}
			models = append(models, model)
			continue
		}
		model.Type = submatch[1]
		if IsInlineHTML(model.Type) {
			var ent SmallB
			xml.Unmarshal([]byte(entity), &ent)
			model.Value = ent
		}
		if IsInlineLeXML(model.Type) {
			var ent ALabel
			xml.Unmarshal([]byte(entity), &ent)
			model.Value = ent
		}
		models = append(models, model)
	}
	return models
}
