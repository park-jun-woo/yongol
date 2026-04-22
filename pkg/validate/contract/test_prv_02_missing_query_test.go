//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestPRV02MissingQuery — 사라진 sqlc 쿼리 참조에 대한 PRV-02 ERROR 발행

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV02MissingQuery(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\nfunc F() {\n"+
			"  server.Queries.UserDeleteByID(ctx, 1)\n"+
			"}\n")
	fs := buildFSForPRV02()
	diags := prv02ExternalSymbolDrift(fs, []string{p})
	if len(diags) == 0 {
		t.Fatalf("expected PRV-02 diag for missing query")
	}
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "UserDeleteByID") && strings.Contains(d.Message, "[PRV-02]") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected missing-query PRV-02 diag, got %+v", diags)
	}
}
