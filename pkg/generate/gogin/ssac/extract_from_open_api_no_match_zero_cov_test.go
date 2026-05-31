//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestExtractFromOpenAPI_NoMatch_ZeroCov — operationId 미매칭 / Paths nil
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractFromOpenAPI_NoMatch_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("Nope")
	doc := docZeroCov("GetWidget")
	g.extractFromOpenAPI(doc, "DoesNotExist")
	if len(g.PathParams) != 0 {
		t.Errorf("expected no params on mismatch, got %v", g.PathParams)
	}

	// Nil paths — early return.
	g2 := newMethodGenZeroCov("X")
	g2.extractFromOpenAPI(&openapi3.T{}, "anything")
	if len(g2.PathParams) != 0 {
		t.Errorf("expected no-op for nil paths")
	}
}
