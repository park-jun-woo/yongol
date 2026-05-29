//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCheckOnePreservedSymbols — preserved 파일 1건 외부 심볼 drift Diagnostic 목록 검증

package contract

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckOnePreservedSymbols(t *testing.T) {
	dir := t.TempDir()
	expQueries := map[string]bool{"FindByID": true}
	expCalls := map[string]bool{}
	expFields := map[string]bool{canonicalFieldKey("email"): true}

	t.Run("no external symbols → nil", func(t *testing.T) {
		p := filepath.Join(dir, "empty.go")
		writePreserved(t, p, "package service\nfunc F() { _ = 1 }\n")
		if d := checkOnePreservedSymbols(p, expQueries, expCalls, expFields); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("drifted DDL field → PRV-02", func(t *testing.T) {
		p := filepath.Join(dir, "drift.go")
		writePreserved(t, p, "package service\nfunc F() { _ = u.DeletedAt }\n")
		diags := checkOnePreservedSymbols(p, expQueries, expCalls, expFields)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "DeletedAt") {
				found = true
			}
		}
		if !found {
			t.Errorf("expected DeletedAt drift diag, got %+v", diags)
		}
	})
}
