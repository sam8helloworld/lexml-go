package entity

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSnippetUnmarshalXML_Success(t *testing.T) {
	input := `<snippet type="表記">snippetのPCDATA</snippet>`
	want := Snippet{
		XMLName: xml.Name{Local: "snippet"},
		Type:    "表記",
		Value: PCDATA{
			Value: "snippetのPCDATA",
		},
	}
	var got Snippet
	err := xml.Unmarshal([]byte(input), &got)
	if err != nil {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
