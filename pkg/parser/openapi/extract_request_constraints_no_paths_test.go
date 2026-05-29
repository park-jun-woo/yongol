//ff:func feature=openapi-parse type=test control=sequence
//ff:what ExtractRequestConstraints — Paths 가 없는 doc 에서도 빈 map 반환

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestConstraints_NoPaths(t *testing.T) {
	doc := &openapi3.T{}
	cs := ExtractRequestConstraints(doc, nil)
	if len(cs) != 0 {
		t.Errorf("got %d entries, want 0", len(cs))
	}
}
