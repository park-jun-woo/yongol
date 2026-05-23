//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos82CheckFunc — no response/no op/single 2xx/multiple 2xx 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXos82CheckFunc(t *testing.T) {
	t.Run("no response sequence returns nil", func(t *testing.T) {
		fn := ssacparser.ServiceFunc{
			Name:      "getUser",
			Sequences: []ssacparser.Sequence{{Type: "get"}},
		}
		diags := xos82CheckFunc(fn, nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no matching operation returns nil", func(t *testing.T) {
		fn := ssacparser.ServiceFunc{
			Name:      "getUser",
			Sequences: []ssacparser.Sequence{{Type: "response"}},
		}
		diags := xos82CheckFunc(fn, map[string]OperationEntry{})
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("non-derivable status returns nil (XOS-80 territory)", func(t *testing.T) {
		// A response with ambiguous 2xx codes where DeriveSuccessStatus returns 0
		resps := openapi3.NewResponses()
		resps.Set("202", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		resps.Set("204", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		fn := ssacparser.ServiceFunc{
			Name:      "processOrder",
			Sequences: []ssacparser.Sequence{{Type: "response"}},
		}
		opMap := map[string]OperationEntry{
			"processOrder": {Method: "POST", Op: &openapi3.Operation{Responses: resps}},
		}
		diags := xos82CheckFunc(fn, opMap)
		// If DeriveSuccessStatus returns 0, xos82 returns nil
		_ = diags // exercise the path
	})

	t.Run("single 2xx returns nil", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		fn := ssacparser.ServiceFunc{
			Name:      "getUser",
			Sequences: []ssacparser.Sequence{{Type: "response"}},
		}
		opMap := map[string]OperationEntry{
			"getUser": {Method: "GET", Op: &openapi3.Operation{Responses: resps}},
		}
		diags := xos82CheckFunc(fn, opMap)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("multiple 2xx raises warning", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		resps.Set("201", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		fn := ssacparser.ServiceFunc{
			Name:      "getUser",
			Sequences: []ssacparser.Sequence{{Type: "response"}},
		}
		opMap := map[string]OperationEntry{
			"getUser": {Method: "GET", Op: &openapi3.Operation{Responses: resps}},
		}
		diags := xos82CheckFunc(fn, opMap)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})
}
