package inline_lexml

var Tags = []string{
	"alabel", "glabel", "slabel", "pos", "POS", "hinshi", "gender", "GENDER", "pron", "hatsuon", "svoc", "ref", "REF", "xref", "XREF", "inflec", "lang", "spellout", "variant",
	"hidden", "light", "sc", "oblq", "audio", "video", "note", "accent", "ex", "etym", "ymd", "article", "ud", "hs", "pha", "ipa", "pinyin", "cn", "jp", "kr", "tw", "ggk", "kigo",
	"mlg", "gi", "gix", "fbox", "FBOX", "pro-n", "pro-v", "abbr", "fullform", "vnum", "primary", "kanbun",
}

type InlineLeXML interface{}

func IsInlineLeXML(t string) bool {
	for _, v := range Tags {
		if t == v {
			return true
		}
	}
	return false
}
