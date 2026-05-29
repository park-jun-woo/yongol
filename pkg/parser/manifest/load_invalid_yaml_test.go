//ff:func feature=manifest type=parser control=sequence
//ff:what 잘못된 YAML 파싱 시 diagnostic 반환 검증
package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(":\ninvalid: [yaml"), 0644)

	_, diags := Load(dir)
	if len(diags) == 0 {
		t.Fatal("expected diagnostics for invalid YAML")
	}
}
