//ff:func feature=manifest type=parser control=sequence
//ff:what ExtractResponseConstraints 가 FieldConstraint.Line 을 LineIndex 값으로 채우는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractResponseConstraints_FillsLine(t *testing.T) {
	path := writeFixture(t)
	idx, _ := BuildLineIndex(path)
	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		t.Fatalf("LoadFromFile: %v", err)
	}
	cs := ExtractResponseConstraints(doc, idx)
	login := cs["Login"]
	if login == nil {
		t.Fatalf("Login response constraints missing")
	}
	if got, want := login["access_token"].Line, 35; got != want {
		t.Errorf("login[access_token].Line = %d, want %d", got, want)
	}
}
