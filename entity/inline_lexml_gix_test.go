package entity

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGixUnmarshalXML_Success(t *testing.T) {
	input := `<gix>gix</gix>`
	want := Gix{
		XMLName: xml.Name{Local: "gix"},
		Value: PCDATA{
			Value: "gix",
		},
	}
	var got Gix
	err := xml.Unmarshal([]byte(input), &got)
	if err != nil {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
