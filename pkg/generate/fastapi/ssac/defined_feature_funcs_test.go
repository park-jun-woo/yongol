//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestDefinedFeatureFuncs — definedFeatureFuncs OperationID→snake_case 정의 집합 구성 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestDefinedFeatureFuncs(t *testing.T) {
	plans := []*ir.ServicePlan{
		{OperationID: "IssueToken"},
		{OperationID: "RefreshToken"},
	}
	got := definedFeatureFuncs(plans)
	if len(got) != 2 {
		t.Fatalf("expected 2 entries, got: %v", got)
	}
	if !got["issue_token"] || !got["refresh_token"] {
		t.Errorf("expected issue_token+refresh_token, got: %v", got)
	}

	if empty := definedFeatureFuncs(nil); len(empty) != 0 {
		t.Errorf("expected empty set for nil plans, got: %v", empty)
	}
}
