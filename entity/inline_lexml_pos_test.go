package entity

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPosUnmarshalXML_Success(t *testing.T) {
	input := `<pos>pos</pos>`
	want := SmallPos{
		XMLName: xml.Name{Local: "pos"},
		Value: PCDATA{
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
