//ff:func feature=ssac-parse type=parser control=sequence
//ff:what parseTestFile 헬퍼 — 소스 문자열을 임시 파일에 쓰고 ParseFile 호출

package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func parseTestFile(t *testing.T, src string) []ServiceFunc {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.go")
	if err := os.WriteFile(path, []byte(src), 0644); err != nil {
		t.Fatal(err)
	}
	sfs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %v", diags[0].Message)
	}
	if len(sfs) == 0 {
		t.Fatal("expected at least 1 ServiceFunc")
	}
	return sfs
}
