//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildResponse_DBModelUnchanged — DB 모델 변수의 @response 는 기존대로 converter 경유

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildResponse_DBModelUnchanged verifies that @response <var> for a
// DB model variable (not a @call result) still routes through
// convert<Model>() — regression guard for BUG-050 fix.
func TestBuildResponse_DBModelUnchanged(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetWorkflow",
		SuccessStatus: 200,
		RespFields:    make(map[string]responseField),
		VarTypes:      map[string]string{"workflow": "Workflow"},
	}
	seq := ssacparser.Sequence{
		Type:   "response",
		Target: "workflow",
	}
	lines, _ := g.buildResponse(seq)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "convertWorkflow(workflow)") {
		t.Fatalf("DB model must use converter, got:\n%s", body)
	}
}
