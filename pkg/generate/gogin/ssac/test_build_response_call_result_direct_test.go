//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildResponse_CallResultUsesConverter — @call 결과 변수의 @response 도 converter 호출 검증

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestBuildResponse_CallResultUsesConverter verifies that @response <var>
// for a @call result variable routes through convert<Type>() just like
// DB model variables (Phase002 — BUG-051 fix).
func TestBuildResponse_CallResultUsesConverter(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetDashboard",
		SuccessStatus: 200,
		RespFields:    make(map[string]responseField),
		VarTypes:      map[string]string{"summary": "SummarizeResponse"},
	}
	seq := ssacparser.Sequence{
		Type:   "response",
		Target: "summary",
	}
	lines := g.buildResponse(seq)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "convertSummarizeResponse(summary)") {
		t.Fatalf("@call result must use converter, got:\n%s", body)
	}
	want := "return api.GetDashboard200JSONResponse(*converted), nil"
	if !strings.Contains(body, want) {
		t.Fatalf("expected converted deref %q, got:\n%s", want, body)
	}
}
