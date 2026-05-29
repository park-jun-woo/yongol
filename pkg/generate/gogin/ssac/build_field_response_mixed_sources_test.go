//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_MixedSources — DB 모델 + @call 결과 모두 converter 경유

package ssac

import (
	"strings"
	"testing"
)

// TestBuildFieldResponse_MixedSources verifies that a @response with
// both a DB model field and a @call result field emits converter calls
// for both (Phase002 — BUG-051 fix).
func TestBuildFieldResponse_MixedSources(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetDashboard",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"summary":  {JSONName: "summary", GoName: "Summary", RefType: "SummarizeResponse", IsRequired: true},
			"workflow": {JSONName: "workflow", GoName: "Workflow", RefType: "Workflow", IsRequired: true},
		},
	}
	fields := map[string]string{
		"summary":  "summary",
		"workflow": "wf",
	}
	lines, _ := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	// DB model field "workflow" must go through converter
	if !strings.Contains(body, "convertWorkflow(wf)") {
		t.Fatalf("DB model field must use converter, got:\n%s", body)
	}

	// @call result field "summary" must also go through converter now
	if !strings.Contains(body, "convertSummarizeResponse(summary)") {
		t.Fatalf("@call result field must use converter, got:\n%s", body)
	}
}
