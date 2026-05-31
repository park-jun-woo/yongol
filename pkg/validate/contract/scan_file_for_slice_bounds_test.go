//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanFileForSliceBounds — FuncDecl 단위 slice[0] 가드 누락 수집 검증
package contract

import (
	"path/filepath"
	"testing"
)

func TestScanFileForSliceBounds(t *testing.T) {
	dir := t.TempDir()

	t.Run("unguarded x[0] flagged", func(t *testing.T) {
		p := filepath.Join(dir, "s.go")
		writePreserved(t, p,
			"package service\nfunc F(xs []int) int { return xs[0] }\n")
		if d := scanFileForSliceBounds(p); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("guarded per-function scope", func(t *testing.T) {
		// Guard in F must not suppress the unguarded access in G.
		p := filepath.Join(dir, "multi.go")
		writePreserved(t, p,
			"package service\nfunc F(xs []int) int {\n  if len(xs) == 0 { return 0 }\n  return xs[0]\n}\n"+
				"func G(ys []int) int { return ys[0] }\n")
		if d := scanFileForSliceBounds(p); len(d) != 1 {
			t.Fatalf("expected 1 diag (G only), got %+v", d)
		}
	})
}
