//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.addParam 단위 테스트 (path → PathParams, query → QueryParams, 그 외 무시)
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMethodGenAddParam(t *testing.T) {
	t.Run("nil param ignored", func(t *testing.T) {
		g := newParamGen()
		g.addParam(nil, "Op")
		g.addParam(&openapi3.ParameterRef{}, "Op")
		if len(g.PathParams) != 0 || len(g.QueryParams) != 0 {
			t.Errorf("nil params should be ignored")
		}
	})
	t.Run("path param registered", func(t *testing.T) {
		g := newParamGen()
		g.addParam(paramRef("id", "path", true, "integer"), "GetX")
		if !g.PathParams["id"] {
			t.Errorf("path param id not registered: %v", g.PathParams)
		}
	})
	t.Run("query param registered with metadata", func(t *testing.T) {
		g := newParamGen()
		g.addParam(paramRef("limit", "query", true, "integer"), "ListX")
		qp, ok := g.QueryParams["limit"]
		if !ok {
			t.Fatalf("query param not registered: %v", g.QueryParams)
		}
		if !qp.IsRequired {
			t.Errorf("expected required query param")
		}
		if qp.GoType != "integer" {
			t.Errorf("GoType = %q, want integer", qp.GoType)
		}
	})
	t.Run("header param ignored", func(t *testing.T) {
		g := newParamGen()
		g.addParam(paramRef("X-Trace", "header", false, "string"), "Op")
		if len(g.PathParams) != 0 || len(g.QueryParams) != 0 {
			t.Errorf("header param should not be registered")
		}
	})
}
