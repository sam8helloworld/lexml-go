package model

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExampleGroupUnmarshalXML_Success(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ExampleGroupChild
	}{
		{
			name:  "exampleをUnmarshalできる",
			input: `<example>例</example>`,
			want: ExampleGroupChild{
				Type: "example",
				Value: Example{
					XMLName: xml.Name{Local: "example"},
					Value:   "例",
				},
			},
		},
		{
			name:  "keyをUnmarshalできる",
			input: `<key>key</key>`,
			want: ExampleGroupChild{
				Type: "key",
				Value: Key{
					XMLName: xml.Name{Local: "key"},
					Value:   "key",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got ExampleGroupChild
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

func TestExampleGroupUnmarshalXML_Fail(t *testing.T) {
	input := []byte(`<unknown>plasma</unknown>`)
	var got ExampleGroupChild
	err := xml.Unmarshal(input, &got)
	if !errors.Is(err, ErrExampleGroupChildrenUnknownElement) {
		t.Fatalf("failed to execute Unmarshal %#v", err)
	}
}
