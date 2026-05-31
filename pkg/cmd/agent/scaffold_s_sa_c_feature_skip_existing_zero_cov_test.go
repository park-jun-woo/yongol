//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldByName_ZeroCov — scaffoldOpenAPI / scaffoldSSaCFeature 직접 호출 (LLM 미사용 분기)
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSSaCFeature_SkipExisting_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	feat := features.Feature{Op: "Login", Path: "/auth/login"}
	domain := domainFromPath(feat.Path)
	svcDir := filepath.Join(dir, "service", domain)
	if err := os.MkdirAll(svcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(svcDir, "Login.ssac"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	// Existing file → skip branch (false, nil); LLM not called.
	created, err := scaffoldSSaCFeature(dir, feat, "", "", Config{}, &out)
	if err != nil {
		t.Fatalf("scaffoldSSaCFeature skip: %v", err)
	}
	if created {
		t.Errorf("expected created=false for existing file")
	}
}
