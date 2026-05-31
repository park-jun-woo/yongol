//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldRego — 기존파일 skip / non-public 없음 skip / non-public 존재+미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldRegoSkipExisting(t *testing.T) {
	dir := t.TempDir()
	policyDir := filepath.Join(dir, "policy")
	if err := os.MkdirAll(policyDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(policyDir, "authz.rego"), []byte("package authz\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := scaffoldRego(dir, &features.FeaturesFile{}, nil, Config{}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skipped (exists)") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}
