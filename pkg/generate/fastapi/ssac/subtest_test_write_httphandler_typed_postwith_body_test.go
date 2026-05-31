//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerTypedPOSTWithBody — POSTWithBody 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerTypedPOSTWithBody(t *testing.T) {

	plan := &ir.ServicePlan{
		OperationID: "CreateWorkflow",
		HTTPMethod:  "POST",
		TriggerKind: ir.TriggerHTTP,
		URLPath:     "/workflow",
		Feature:     "workflow",
		BodyFields: []ir.BodyFieldMeta{
			{Name: "title"},
		},
	}
	var b strings.Builder
	writeHTTPHandler(&b, plan)
	got := b.String()
	if !strings.Contains(got, "body: CreateWorkflowRequest") {
		t.Errorf("POST with body should have Pydantic model, got: %s", got)
	}

}
