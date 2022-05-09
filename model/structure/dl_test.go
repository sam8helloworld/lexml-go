package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model"
)

func TestDlUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  DlChild
	}{
		{
			name:  "dtをUnmarshalできる",
			input: `<dt>dt</dt>`,
			want: DlChild{
				Type: "dt",
				Value: Dt{
					XMLName: xml.Name{Local: "dt"},
					Value:   "dt",
				},
			},
		},
		{
			name:  "ddをUnmarshalできる",
			input: `<dd>dd</dd>`,
			want: DlChild{
				Type: "dd",
				Value: Dd{
					XMLName: xml.Name{Local: "dd"},
					Value: model.InnerXML{
						Value: "dd",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got DlChild
			err := xml.Unmarshal([]byte(tt.input), &got)
			if err != nil {
				t.Fatalf("failed to execute Unmarshal %#v", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestDlUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got DlChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrDlChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
