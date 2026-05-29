//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_CallResultUsesConverter — @call 결과 변수가 $ref 필드에 매핑되면 converter 호출

package ssac

import (
	"strings"
	"testing"
)

// TestBuildFieldResponse_CallResultUsesConverter verifies that when a
// @response field maps to a @call result variable and the field schema
// is a $ref type, the converter call is emitted (Phase002 — BUG-051 fix).
func TestBuildFieldResponse_CallResultUsesConverter(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetDashboard",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"summary": {JSONName: "summary", GoName: "Summary", RefType: "SummarizeResponse", IsRequired: true},
		},
	}
	fields := map[string]string{
		"summary": "summary",
	}
	lines, _ := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "convertSummarizeResponse(summary)") {
		t.Fatalf("@call result field must use converter, got:\n%s", body)
	}
	if !strings.Contains(body, "summaryConverted") {
		t.Fatalf("expected hoisted converter local var, got:\n%s", body)
	}
}
