//ff:func feature=manifest type=test control=sequence
//ff:what ExtractRequestConstraints / ExtractResponseConstraints 가 nil doc 에서도 빈 map 을 반환하는지 검증 (BuildLineIndex nil-safety 포함)

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestConstraints_NilDoc(t *testing.T) {
	cs := ExtractRequestConstraints(nil, nil)
	if cs == nil {
		t.Fatal("nil doc should return empty non-nil map")
	}
	if len(cs) != 0 {
		t.Errorf("got %d entries, want 0", len(cs))
	}
}

func TestExtractResponseConstraints_NilDoc(t *testing.T) {
	cs := ExtractResponseConstraints(nil, nil)
	if cs == nil {
		t.Fatal("nil doc should return empty non-nil map")
	}
	if len(cs) != 0 {
		t.Errorf("got %d entries, want 0", len(cs))
	}
}

func TestExtractRequestConstraints_NoPaths(t *testing.T) {
	doc := &openapi3.T{}
	cs := ExtractRequestConstraints(doc, nil)
	if len(cs) != 0 {
		t.Errorf("got %d entries, want 0", len(cs))
	}
}
