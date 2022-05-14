package entity

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKeyUnmarshalXML_Success(t *testing.T) {
	input := `<key type="表記">keyのPCDATA</key>`
	want := Key{
		XMLName: xml.Name{Local: "key"},
		Type:    "表記",
		Value: PCDATA{
			Value: "keyのPCDATA",
		},
	}
	var got Key
	err := xml.Unmarshal([]byte(input), &got)
	if err != nil {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("got differs: (-got +want)\n%s", diff)
	}
}
