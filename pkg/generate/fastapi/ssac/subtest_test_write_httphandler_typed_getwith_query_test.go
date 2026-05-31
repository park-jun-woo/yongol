//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerTypedGETWithQuery — GETWithQuery 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerTypedGETWithQuery(t *testing.T) {

	plan := &ir.ServicePlan{
		OperationID: "ListWorkflows",
		HTTPMethod:  "GET",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/workflow",
		Feature:     "workflow",
		QueryParams: []ir.QueryParamMeta{
			{Name: "per_page", Type: "integer"},
			{Name: "cursor", Type: "string", Required: false},
		},
	}
	var b strings.Builder
	writeHTTPHandler(&b, plan)
	got := b.String()
	if !strings.Contains(got, "per_page: int | None = None") {
		t.Errorf("expected optional int query param, got: %s", got)
	}
	if !strings.Contains(got, "cursor: str | None = None") {
		t.Errorf("expected optional str query param, got: %s", got)
	}

}
