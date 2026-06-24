//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestEmitConvertFiles_ZeroCov — emitConvertFuncFile / emitConvertListFile 파일 기록
package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEmitConvertFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}

	if err := emitConvertFuncFile(dir, "example.com/app", "Widget", convertSchemaZeroCov(), nil, used, domainGen{}); err != nil {
		t.Fatalf("emitConvertFuncFile: %v", err)
	}
	if err := emitConvertListFile(dir, "example.com/app", "Widget", used, domainGen{}); err != nil {
		t.Fatalf("emitConvertListFile: %v", err)
	}

	got, _ := os.ReadDir(dir)
	if len(got) == 0 {
		t.Fatalf("expected converter files emitted, got none")
	}
	// at least one file references convertWidget
	found := false
	for _, e := range got {
		b, _ := os.ReadFile(filepath.Join(dir, e.Name()))
		if strings.Contains(string(b), "convertWidget") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a file containing convertWidget, files=%v", got)
	}
}
