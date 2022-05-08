package inline_lexml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestPosUnmarshalXML_Success(t *testing.T) {
	input := `<pos>pos</pos>`
	want := SmallPos{
		XMLName: xml.Name{Local: "pos"},
		Value: pcdata.PCDATA{
			Value: "pos",
		},
	}
	var got SmallPos
	err := xml.Unmarshal([]byte(input), &got)
	if err != nil {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
