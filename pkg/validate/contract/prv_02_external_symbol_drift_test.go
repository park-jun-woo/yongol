//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestPRV02ExternalSymbolDrift — preserved 파일 목록의 외부 심볼 drift 오케스트레이션 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPRV02ExternalSymbolDrift(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "activate_workflow.go")
	writePreserved(t, p, "package service\nfunc F() { _ = u.DeletedAt }\n")
	diags := prv02ExternalSymbolDrift(buildFSForPRV02(), []string{p})
	found := false
	for _, d := range diags {
		if strings.Contains(d.Message, "[PRV-02]") && strings.Contains(d.Message, "DeletedAt") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected PRV-02 drift diag, got %+v", diags)
	}

	t.Run("empty paths → no diags", func(t *testing.T) {
		if d := prv02ExternalSymbolDrift(buildFSForPRV02(), nil); len(d) != 0 {
			t.Errorf("expected no diags, got %+v", d)
		}
	})
}
