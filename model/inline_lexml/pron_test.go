package inline_lexml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestPronUnmarshalXML_Success(t *testing.T) {
	input := `<pron>pron</pron>`
	want := Pron{
		XMLName: xml.Name{Local: "pron"},
		Value: pcdata.PCDATA{
			Value: "pron",
		},
	}
	var got Pron
	err := xml.Unmarshal([]byte(input), &got)
	if err != nil {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
