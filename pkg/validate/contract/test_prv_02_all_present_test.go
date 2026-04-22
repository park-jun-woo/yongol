//ff:func feature=validate-contract type=test control=sequence
//ff:what TestPRV02AllPresent — 모든 외부 심볼이 SSOT에 있으면 PRV-02 ERROR 없음

package contract

import (
	"path/filepath"
	"testing"
)

func TestPRV02AllPresent(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\n\nimport \"billing\"\n\nfunc F() {\n"+
			"  server.Queries.UserFindByID(ctx, 1)\n"+
			"  billing.CheckCredits(req)\n"+
			"  _ = u.Email\n"+
			"}\n")
	fs := buildFSForPRV02()
	diags := prv02ExternalSymbolDrift(fs, []string{p})
	if len(diags) != 0 {
		t.Errorf("expected no diags when symbols match SSOT, got %+v", diags)
	}
}
