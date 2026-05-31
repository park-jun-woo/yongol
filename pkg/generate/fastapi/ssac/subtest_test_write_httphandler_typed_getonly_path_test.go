//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerTypedGETOnlyPath — GETOnlyPath 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerTypedGETOnlyPath(t *testing.T) {

	plan := &ir.ServicePlan{
		OperationID: "GetWorkflow",
		HTTPMethod:  "GET",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/workflow/:id",
		Feature:     "workflow",
		PathParams:  []string{"id"},
	}
	var b strings.Builder
	writeHTTPHandler(&b, plan)
	got := b.String()
	if !strings.Contains(got, "id: int,") {
		t.Errorf("GET should have typed path param, got: %s", got)
	}
	if strings.Contains(got, "body:") {
		t.Errorf("GET should not have body, got: %s", got)
	}
	if !strings.Contains(got, "Depends(get_current_user)") {
		t.Errorf("expected Depends(get_current_user), got: %s", got)
	}
	if !strings.Contains(got, "Depends(get_session)") {
		t.Errorf("expected Depends(get_session), got: %s", got)
	}

}
