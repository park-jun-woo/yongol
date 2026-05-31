//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerTypedDELETENoBody — DELETENoBody 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerTypedDELETENoBody(t *testing.T) {

	plan := &ir.ServicePlan{
		OperationID: "DeleteWorkflow",
		HTTPMethod:  "DELETE",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/workflow/:id",
		Feature:     "workflow",
		PathParams:  []string{"id"},
	}
	var b strings.Builder
	writeHTTPHandler(&b, plan)
	got := b.String()
	if !strings.Contains(got, "id: int,") {
		t.Errorf("DELETE should have path param, got: %s", got)
	}
	if strings.Contains(got, "body:") {
		t.Errorf("DELETE should not have body, got: %s", got)
	}

}
