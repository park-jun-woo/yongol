//ff:func feature=manifest type=parser control=sequence
//ff:what ExtractRequestConstraints 가 nil LineIndex 에서도 panic 없이 동작하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestConstraints_NilLineIndex(t *testing.T) {
	path := writeFixture(t)
	doc, err := openapi3.NewLoader().LoadFromFile(path)
	if err != nil {
		t.Fatalf("LoadFromFile: %v", err)
	}
	// nil LineIndex 도 panic 없이 동작해야 한다 (Line = 0 으로 남음).
	cs := ExtractRequestConstraints(doc, nil)
	if cs["Login"] == nil {
		t.Fatalf("Login missing")
	}
	if got := cs["Login"]["email"].Line; got != 0 {
		t.Errorf("nil-lines email.Line = %d, want 0", got)
	}
}
