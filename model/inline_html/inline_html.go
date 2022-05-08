package inline_html

var Tags = []string{
	"b", "B", "br", "BR", "em", "EM", "i", "I", "nobr", "span", "sub", "SUB", "sup", "SUP", "u", "U", "ruby", "a", "img", "big", "small",
}

type InlineHTML interface{}

func IsInlineHTML(t string) bool {
	for _, v := range Tags {
		if t == v {
			return true
		}
	}
	return false
}
