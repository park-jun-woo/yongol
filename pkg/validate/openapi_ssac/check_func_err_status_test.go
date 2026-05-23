//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what checkFuncErrStatus — status 정의됨/미정의됨/비guard 시퀀스 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCheckFuncErrStatus(t *testing.T) {
	t.Run("empty sequences returns nil", func(t *testing.T) {
		resps := openapi3.NewResponses()
		op := &openapi3.Operation{Responses: resps}
		diags := checkFuncErrStatus("test.ssac", "getUser", nil, op, nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("non-guard sequence skipped", func(t *testing.T) {
		resps := openapi3.NewResponses()
		op := &openapi3.Operation{Responses: resps}
		seqs := []ssac.Sequence{
			{Type: "response"},
		}
		diags := checkFuncErrStatus("test.ssac", "getUser", seqs, op, nil)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("status defined in OpenAPI passes", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		op := &openapi3.Operation{Responses: resps}
		seqs := []ssac.Sequence{
			{Type: "empty", ErrStatus: 404},
		}
		diags := checkFuncErrStatus("test.ssac", "getUser", seqs, op, nil)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("status not in OpenAPI raises diagnostic", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		op := &openapi3.Operation{Responses: resps}
		seqs := []ssac.Sequence{
			{Type: "empty", ErrStatus: 404},
		}
		diags := checkFuncErrStatus("test.ssac", "getUser", seqs, op, nil)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-21") {
			t.Errorf("Message missing XOS-21: %s", diags[0].Message)
		}
	})
}
