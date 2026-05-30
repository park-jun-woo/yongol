//ff:func feature=agent type=test control=iteration dimension=1
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

func TestScaffoldRegoNoNonPublic(t *testing.T) {
	// Only public features → no non-public → skip.
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "Ping", Public: true}}}
	var out bytes.Buffer
	if err := scaffoldRego(t.TempDir(), ff, nil, Config{}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "no non-public") {
		t.Errorf("expected no-non-public message, got: %q", out.String())
	}
}

func TestScaffoldRegoLLMError(t *testing.T) {
	// A non-public feature triggers the batch LLM call, which fails for the
	// unsupported backend.
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "CreateUser", Public: false}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffoldRego(t.TempDir(), ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from rego batch")
	}
}

func TestScaffoldRegoMkdirError(t *testing.T) {
	// Creating "policy" as a regular file makes os.MkdirAll(policyDir) fail.
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "policy"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "CreateUser", Public: false}}}
	var out bytes.Buffer
	if err := scaffoldRego(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when policy is a file")
	}
}
