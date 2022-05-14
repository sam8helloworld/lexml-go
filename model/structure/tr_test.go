package structure

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sam8helloworld/lexml-go/model"
)

func TestTrUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  TrChild
	}{
		{
			name:  "thをUnmarshalできる",
			input: `<th>th</th>`,
			want: TrChild{
				Type: "th",
				Value: Th{
					XMLName: xml.Name{Local: "th"},
					Value:   "th",
				},
			},
		},
		{
			name:  "tdをUnmarshalできる",
			input: `<td>td</td>`,
			want: TrChild{
				Type: "td",
				Value: Td{
					XMLName: xml.Name{Local: "td"},
					Value: model.InnerXML{
						Value: "td",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got TrChild
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

func TestTrUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got TrChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrTrChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
