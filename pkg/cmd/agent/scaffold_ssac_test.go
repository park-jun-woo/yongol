//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestScaffoldSSaC — features 없음 0,nil / 기존파일 skip(count 0) / 미지원 backend LLM 에러 분기 검증

package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSSaCNoFeatures(t *testing.T) {
	var out bytes.Buffer
	n, err := scaffoldSSaC(t.TempDir(), &features.FeaturesFile{}, "", nil, Config{}, &out)
	if err != nil || n != 0 {
		t.Fatalf("no features → %d, %v; want 0, nil", n, err)
	}
}

func TestScaffoldSSaCSkipExisting(t *testing.T) {
	dir := t.TempDir()
	svcDir := filepath.Join(dir, "service", "auth")
	if err := os.MkdirAll(svcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(svcDir, "Login.ssac"), []byte("func Login() {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "Login", Path: "/auth/login"}}}
	var out bytes.Buffer
	n, err := scaffoldSSaC(dir, ff, "", nil, Config{}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 0 {
		t.Errorf("expected count 0 (skipped), got %d", n)
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}

func TestScaffoldSSaCLLMError(t *testing.T) {
	// A missing SSaC file + unsupported backend makes scaffoldSSaCFeature's LLM
	// call fail, propagated out of scaffoldSSaC.
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "Login", Path: "/auth/login"}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if _, err := scaffoldSSaC(t.TempDir(), ff, "", nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldSSaCFeature")
	}
}
