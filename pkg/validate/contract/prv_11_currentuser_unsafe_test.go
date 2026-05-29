//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV11CurrentUserUnsafe — 단일 대입 currentUser 단언은 ERROR

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV11CurrentUserUnsafe(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"context\"\n\n"+
			"func h(ctx context.Context) {\n"+
			"\tcu := ctx.Value(\"currentUser\").(*model.UserClaim)\n"+
			"\t_ = cu\n}\n")
	diags := prv11PreservedCurrentUserAssertion([]string{p})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
}
