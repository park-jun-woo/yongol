//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_MixedSources — DB 모델 + @call 결과 혼재 시 올바른 분기

package ssac

import (
	"strings"
	"testing"
)

// TestBuildFieldResponse_MixedSources verifies that a @response with
// both a DB model field (needs converter) and a @call result field
// (direct assign) emits the correct code for each.
func TestBuildFieldResponse_MixedSources(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetDashboard",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"summary":  {JSONName: "summary", GoName: "Summary", RefType: "SummarizeResponse", IsRequired: true},
			"workflow": {JSONName: "workflow", GoName: "Workflow", RefType: "Workflow", IsRequired: true},
		},
		CallResultVars: map[string]bool{"summary": true},
	}
	fields := map[string]string{
		"summary":  "summary",
		"workflow": "wf",
	}
	lines := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	// DB model field "workflow" must go through converter
	if !strings.Contains(body, "convertWorkflow(wf)") {
		t.Fatalf("DB model field must use converter, got:\n%s", body)
	}

	// @call result field "summary" must NOT go through converter
	if strings.Contains(body, "convertSummarizeResponse") {
		t.Fatalf("@call result field must NOT use converter, got:\n%s", body)
	}
	if !strings.Contains(body, "Summary: summary,") {
		t.Fatalf("expected direct assignment for summary, got:\n%s", body)
	}
}
