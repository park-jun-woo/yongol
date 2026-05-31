//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestExtractBodyFormats_NoBody_ZeroCov — request body 없음 early-return
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractBodyFormats_NoBody_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("X")
	op := openapi3.NewOperation()
	g.extractBodyFormats(op) // op.RequestBody == nil → returns
	if len(g.BodyFormats) != 0 {
		t.Errorf("expected empty BodyFormats, got %v", g.BodyFormats)
	}
}
