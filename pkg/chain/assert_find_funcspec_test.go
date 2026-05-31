//ff:func feature=chain type=test-helper control=sequence
//ff:what assertFindFuncSpec — findFuncSpecLink 결과(ok/kind/file/summary/line) 검증 헬퍼
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// assertFindFuncSpec invokes findFuncSpecLink for tc and asserts the result.
func assertFindFuncSpec(t *testing.T, tc findFuncSpecCase, specs []funcspec.FuncSpec, specsDir string) {
	t.Helper()
	callRef := tc.pkg + "." + tc.funcName
	link, ok := findFuncSpecLink(callRef, tc.pkg, tc.funcName, specs, specsDir)
	if ok != tc.wantOK {
		t.Fatalf("ok: got %v, want %v", ok, tc.wantOK)
	}
	if !tc.wantOK {
		return
	}
	if link.Kind != "FuncSpec" {
		t.Errorf("kind: got %q, want FuncSpec", link.Kind)
	}
	if link.File != tc.wantFile {
		t.Errorf("file: got %q, want %q", link.File, tc.wantFile)
	}
	if link.Summary != tc.wantSumm {
		t.Errorf("summary: got %q, want %q", link.Summary, tc.wantSumm)
	}
	if tc.wantLineP && link.Line <= 0 {
		t.Errorf("line: got %d, want > 0", link.Line)
	}
}
