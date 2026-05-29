//ff:func feature=gen-gogin type=test control=sequence topic=publish
//ff:what TestBuildPublish_MapValue — @publish payload에 mapValue 적용하여 request.id → request.Id 변환 검증

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildPublish_MapValue confirms that buildPublish applies mapValue to
// payload values so that request.<field> expressions undergo PascalCase
// conversion via mapRequestValue. Without mapValue, "request.id" would be
// emitted verbatim, causing a compile error against oapi-codegen structs
// that use "request.Id" (BUG-054).
func TestBuildPublish_MapValue(t *testing.T) {
	g := &methodGen{
		FuncName:   "ExecuteWorkflow",
		PathParams: map[string]bool{"id": true},
	}
	seq := ssacparser.Sequence{
		Type:  "publish",
		Model: "workflow.execute",
		Inputs: map[string]string{
			"WorkflowID": "request.id",
		},
	}
	lines, _ := g.buildPublish(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "request.Id") {
		t.Fatalf("expected request.id to be mapped to request.Id, got:\n%s", body)
	}
	if strings.Contains(body, "request.id") {
		t.Fatalf("request.id should not appear verbatim in output, got:\n%s", body)
	}
}
