//ff:func feature=agent type=test control=iteration dimension=3
//ff:what TestExtractOperationIDs — 진단 메시지에서 고유 operationId 추출(중복 제거) 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestExtractOperationIDs(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Message: `operationId "CreateUser" mismatch`},
		{Message: `SSaC func GetUser missing`},
		{Message: `operationId: CreateUser duplicate`}, // dup, should not repeat
		{Message: `Missing: ListUsers, DeleteUser`},
	}
	got := extractOperationIDs(diags)

	want := map[string]bool{
		"CreateUser": false,
		"GetUser":    false,
		"ListUsers":  false,
		"DeleteUser": false,
	}
	for _, id := range got {
		if _, ok := want[id]; !ok {
			t.Errorf("unexpected id %q in %v", id, got)
			continue
		}
		want[id] = true
	}
	for id, seen := range want {
		if !seen {
			t.Errorf("missing expected id %q in %v", id, got)
		}
	}
	// Uniqueness: CreateUser appears twice but must occur once.
	count := 0
	for _, id := range got {
		if id == "CreateUser" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("CreateUser appeared %d times, want 1", count)
	}
}
