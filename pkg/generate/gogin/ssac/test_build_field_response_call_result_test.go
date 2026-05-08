//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_CallResultSkipsConverter — @call 결과 변수가 $ref 필드에 매핑되면 converter 건너뜀

package ssac

import (
	"strings"
	"testing"
)

// TestBuildFieldResponse_CallResultSkipsConverter verifies that when a
// @response field maps to a @call result variable and the field schema
// is a $ref type, the converter call is skipped and the variable is
// assigned directly (BUG-050).
func TestBuildFieldResponse_CallResultSkipsConverter(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetDashboard",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"summary": {JSONName: "summary", GoName: "Summary", RefType: "SummarizeResponse", IsRequired: true},
		},
		CallResultVars: map[string]bool{"summary": true},
	}
	fields := map[string]string{
		"summary": "summary",
	}
	lines := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	if strings.Contains(body, "convert") {
		t.Fatalf("@call result field must NOT use converter, got:\n%s", body)
	}
	if !strings.Contains(body, "Summary: summary,") {
		t.Fatalf("expected direct assignment, got:\n%s", body)
	}
}
