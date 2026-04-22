//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestPRV02MissingField — 사라진 DDL 컬럼 필드 참조에 대한 PRV-02 ERROR 발행

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV02MissingField(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p,
		"package service\nfunc F() {\n"+
			"  _ = u.DeletedAt\n"+
			"}\n")
	fs := buildFSForPRV02()
	diags := prv02ExternalSymbolDrift(fs, []string{p})
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "DeletedAt") && strings.Contains(d.Message, "[PRV-02]") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected missing-field PRV-02 diag, got %+v", diags)
	}
}
