//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what o02CaseConflictDiag — 단일 케이스(no conflict)/복수 케이스(conflict) 검증

package openapi

import (
	"strings"
	"testing"
)

func TestO02CaseConflictDiag(t *testing.T) {
	t.Run("single case variant returns false", func(t *testing.T) {
		cases := map[string]map[string]bool{
			"user_id": {"/users/{user_id}": true},
		}
		_, ok := o02CaseConflictDiag("user_id", cases)
		if ok {
			t.Error("expected false for single variant")
		}
	})

	t.Run("empty map returns false", func(t *testing.T) {
		_, ok := o02CaseConflictDiag("id", map[string]map[string]bool{})
		if ok {
			t.Error("expected false for empty map")
		}
	})

	t.Run("multiple case variants returns diagnostic", func(t *testing.T) {
		cases := map[string]map[string]bool{
			"userId":  {"/users/{userId}": true},
			"user_id": {"/accounts/{user_id}": true},
		}
		diag, ok := o02CaseConflictDiag("userid", cases)
		if !ok {
			t.Fatal("expected true for conflict")
		}
		if !strings.Contains(diag.Message, "O-2") {
			t.Errorf("Message missing O-2: %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "userId") {
			t.Errorf("Message missing variant userId: %s", diag.Message)
		}
		if !strings.Contains(diag.Message, "user_id") {
			t.Errorf("Message missing variant user_id: %s", diag.Message)
		}
	})
}
