package inline_lexml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model/pcdata"
)

func TestGenderUnmarshalXML_Success(t *testing.T) {
	input := `<gender>gender</gender>`
	want := SmallGender{
		XMLName: xml.Name{Local: "gender"},
		Value: pcdata.PCDATA{
			Value: "gender",
		},
	}
	var got SmallGender
	err := xml.Unmarshal([]byte(input), &got)
	if err != nil {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
