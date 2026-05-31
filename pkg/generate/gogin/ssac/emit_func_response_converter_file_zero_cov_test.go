//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitFuncResponseConverterFile_ZeroCov — 단일 func 응답 변환 파일 emit
package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEmitFuncResponseConverterFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}
	info := funcRespInfo{PkgAlias: "dashboard", ImportPath: "example.com/app/internal/dashboard"}
	if err := emitFuncResponseConverterFile(dir, "example.com/app", "SummarizeResponse", convertSchemaZeroCov(), info, nil, used); err != nil {
		t.Fatalf("emitFuncResponseConverterFile: %v", err)
	}
	got, _ := os.ReadDir(dir)
	if len(got) == 0 {
		t.Fatalf("expected a file emitted")
	}
	b, _ := os.ReadFile(filepath.Join(dir, got[0].Name()))
	if !strings.Contains(string(b), "convertSummarizeResponse") {
		t.Errorf("expected convertSummarizeResponse in:\n%s", string(b))
	}
}
