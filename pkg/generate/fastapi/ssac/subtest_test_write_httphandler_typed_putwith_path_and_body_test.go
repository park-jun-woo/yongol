//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerTypedPUTWithPathAndBody — PUTWithPathAndBody 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerTypedPUTWithPathAndBody(t *testing.T) {

	plan := &ir.ServicePlan{
		OperationID: "UpdateWorkflow",
		HTTPMethod:  "PUT",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/workflow/:id",
		Feature:     "workflow",
		PathParams:  []string{"id"},
		BodyFields: []ir.BodyFieldMeta{
			{Name: "title"},
			{Name: "status"},
		},
	}
	var b strings.Builder
	writeHTTPHandler(&b, plan)
	got := b.String()
	if !strings.Contains(got, "id: int,") {
		t.Errorf("PUT should have path param, got: %s", got)
	}
	if !strings.Contains(got, "body: UpdateWorkflowRequest") {
		t.Errorf("PUT should have body, got: %s", got)
	}

}
