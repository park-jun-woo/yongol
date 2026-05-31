//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestExtractRespFields_NoResp_ZeroCov — 매칭 응답 없음 early-return
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRespFields_NoResp_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("X")
	op := openapi3.NewOperation()
	op.Responses = openapi3.NewResponses()
	g.extractRespFields(op) // no 200 → returns
	if len(g.RespFields) != 0 {
		t.Errorf("expected empty RespFields, got %v", g.RespFields)
	}
}
