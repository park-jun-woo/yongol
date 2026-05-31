//ff:func feature=cli-init type=test control=sequence
//ff:what TestGenerateRegoFromFeatures — allow rule stub 생성 / write 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestGenerateRegoFromFeatures_Success(t *testing.T) {
	target := t.TempDir()
	if err := os.MkdirAll(filepath.Join(target, "specs", "policy"), 0o755); err != nil {
		t.Fatal(err)
	}
	feats := []features.Feature{{Op: "CreateTask", Path: "POST /tasks"}}
	if err := generateRegoFromFeatures(target, feats); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(target, "specs", "policy", "authz.rego"))
	if err != nil {
		t.Fatal(err)
	}
	s := string(got)
	if !strings.Contains(s, "package authz") || !strings.Contains(s, "default allow := false") {
		t.Errorf("header missing: %q", s)
	}
	if !strings.Contains(s, `input.action == "CreateTask"`) || !strings.Contains(s, `input.resource == "task"`) {
		t.Errorf("allow rule missing: %q", s)
	}
}
