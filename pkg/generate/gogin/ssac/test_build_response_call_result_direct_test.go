//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildResponse_CallResultDirect — @call 결과 변수의 @response 직접 대입 (converter 미사용)

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildResponse_CallResultDirect verifies that @response <var> for a
// @call result variable emits a direct cast without convert<Type>() (BUG-050).
func TestBuildResponse_CallResultDirect(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetDashboard",
		SuccessStatus: 200,
		RespFields:    make(map[string]responseField),
		VarTypes:      map[string]string{"summary": "SummarizeResponse"},
		CallResultVars: map[string]bool{"summary": true},
	}
	seq := ssacparser.Sequence{
		Type:   "response",
		Target: "summary",
	}
	lines := g.buildResponse(seq)
	body := strings.Join(lines, "\n")

	if strings.Contains(body, "convert") {
		t.Fatalf("@call result must NOT use converter, got:\n%s", body)
	}
	want := "return api.GetDashboard200JSONResponse(summary), nil"
	if !strings.Contains(body, want) {
		t.Fatalf("expected direct cast %q, got:\n%s", want, body)
	}
}
