//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what p01Parse — Rego parse 검증 (이미 파싱됨/빈 디렉토리/유효 파일) 검증
package rego

import (
	"os"
	"path/filepath"
	"testing"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestP01Parse_Unit(t *testing.T) {
	t.Run("already parsed returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ParsedPolicies: []regoparser.Policy{
				{File: "auth.rego"},
			},
		}
		diags := p01Parse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty policy dir returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		policyDir := filepath.Join(tmp, "policy")
		os.MkdirAll(policyDir, 0o755)
		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := p01Parse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("invalid rego file returns P-1 diagnostics", func(t *testing.T) {
		tmp := t.TempDir()
		policyDir := filepath.Join(tmp, "policy")
		os.MkdirAll(policyDir, 0o755)
		os.WriteFile(filepath.Join(policyDir, "bad.rego"), []byte("this is not valid rego {{{"), 0o644)
		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := p01Parse(fs)
		if len(diags) == 0 {
			t.Fatal("expected diagnostics for invalid rego")
		}
		for _, d := range diags {
			if len(d.Message) < 4 || d.Message[:4] != "[P-1" {
				t.Errorf("expected [P-1] prefix, got %s", d.Message)
			}
		}
	})

	t.Run("valid rego file returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		policyDir := filepath.Join(tmp, "policy")
		os.MkdirAll(policyDir, 0o755)
		os.WriteFile(filepath.Join(policyDir, "auth.rego"), []byte("package authz\n\ndefault allow = false\n"), 0o644)
		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := p01Parse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
