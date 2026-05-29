//ff:func feature=manifest type=parser control=sequence
//ff:what ExtractRequestConstraints 가 FieldConstraint.Line 을 LineIndex 값으로 채우는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestConstraints_FillsLine(t *testing.T) {
	path := writeFixture(t)
	idx, err := BuildLineIndex(path)
	if err != nil {
		t.Fatalf("BuildLineIndex: %v", err)
	}
	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		t.Fatalf("LoadFromFile: %v", err)
	}
	cs := ExtractRequestConstraints(doc, idx)
	login := cs["Login"]
	if login == nil {
		t.Fatalf("Login request constraints missing")
	}
	if got, want := login["email"].Line, 25; got != want {
		t.Errorf("login[email].Line = %d, want %d", got, want)
	}
	if got, want := login["password"].Line, 26; got != want {
		t.Errorf("login[password].Line = %d, want %d", got, want)
	}
}
