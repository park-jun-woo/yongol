//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPRV11CurrentUserCommaOK — comma-ok 형식 단언은 PRV-11 면제

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV11CurrentUserCommaOK(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"context\"\n\n"+
			"func h(ctx context.Context) {\n"+
			"\tcu, ok := ctx.Value(\"currentUser\").(*model.CurrentUser)\n"+
			"\t_ = cu; _ = ok\n}\n")
	diags := prv11PreservedCurrentUserAssertion([]string{p})
	if len(diags) != 0 {
		t.Fatalf("comma-ok form should be safe, got %+v", diags)
	}
}
