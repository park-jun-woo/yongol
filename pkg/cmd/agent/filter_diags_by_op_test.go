//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestFilterDiagsByOp — op 포함 진단만 필터, 없으면 원본 반환 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestFilterDiagsByOp(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Message: "CreateUser mismatch"},
		{Message: "ListOrders missing"},
		{Message: "CreateUser dup"},
	}
	got := filterDiagsByOp(diags, "CreateUser")
	if len(got) != 2 {
		t.Errorf("filterDiagsByOp = %d diags, want 2", len(got))
	}
	// No matches -> return original slice unchanged.
	got = filterDiagsByOp(diags, "Nonexistent")
	if len(got) != 3 {
		t.Errorf("no-match fallback = %d, want 3 (original)", len(got))
	}
}
