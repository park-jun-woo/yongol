//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestExtractFromOpenAPI_ZeroCov — operationId 매칭 후 path/query/body/resp 메타 적재
package ssac

import (
	"testing"
)

func TestExtractFromOpenAPI_ZeroCov(t *testing.T) {
	g := newMethodGenZeroCov("GetWidget")
	doc := docZeroCov("GetWidget")
	g.extractFromOpenAPI(doc, "GetWidget")

	if !g.PathParams["id"] {
		t.Errorf("expected path param id, got %v", g.PathParams)
	}
	if _, ok := g.QueryParams["q"]; !ok {
		t.Errorf("expected query param q, got %v", g.QueryParams)
	}
	if g.BodyFormats["email"] != "email" {
		t.Errorf("expected email format, got %v", g.BodyFormats)
	}
	if g.BodyFormats["plan"] != "enum" {
		t.Errorf("expected plan enum, got %v", g.BodyFormats)
	}
	if !g.BodyJSONBFields["meta"] {
		t.Errorf("expected meta marked JSONB, got %v", g.BodyJSONBFields)
	}
	if !g.BodyRequiredFields["email"] {
		t.Errorf("expected email required, got %v", g.BodyRequiredFields)
	}
	wf, ok := g.RespFields["widget"]
	if !ok || wf.RefType != "Widget" {
		t.Errorf("expected widget RefType Widget, got %+v (ok=%v)", wf, ok)
	}
	if g.Method != "GET" {
		t.Errorf("Method = %q, want GET", g.Method)
	}
}
