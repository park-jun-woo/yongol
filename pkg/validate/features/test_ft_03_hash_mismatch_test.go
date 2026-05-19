//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-03 — features.yaml 해시 불일치 시 에러 진단 테스트

package features

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT03_HashMismatch_Fires(t *testing.T) {
	tmp := t.TempDir()

	// Write features.yaml
	featContent := []byte("features:\n  - op: CreateWorkflow\n    path: POST /workflows\n    desc: Create\n")
	if err := os.WriteFile(filepath.Join(tmp, "features.yaml"), featContent, 0o644); err != nil {
		t.Fatal(err)
	}

	// Write .yongol with wrong hash
	yongolContent := "hashes:\n  features.yaml: sha256:0000000000000000000000000000000000000000000000000000000000000000\n"
	if err := os.WriteFile(filepath.Join(tmp, ".yongol"), []byte(yongolContent), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := ft03HashMismatch(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-03]") {
		t.Errorf("want [FT-03] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "modified after baseline") {
		t.Errorf("want 'modified after baseline' message, got %s", diags[0].Message)
	}
}

func TestFT03_MissingYongolFile_Fires(t *testing.T) {
	tmp := t.TempDir()

	// Write features.yaml without .yongol
	featContent := []byte("features:\n  - op: CreateWorkflow\n    path: POST /workflows\n    desc: Create\n")
	if err := os.WriteFile(filepath.Join(tmp, "features.yaml"), featContent, 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := ft03HashMismatch(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-03]") {
		t.Errorf("want [FT-03] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, ".yongol not found") {
		t.Errorf("want '.yongol not found' message, got %s", diags[0].Message)
	}
}

func TestFT03_MatchingHash_NoFire(t *testing.T) {
	tmp := t.TempDir()

	// Write features.yaml
	featContent := []byte("features:\n  - op: CreateWorkflow\n    path: POST /workflows\n    desc: Create\n")
	if err := os.WriteFile(filepath.Join(tmp, "features.yaml"), featContent, 0o644); err != nil {
		t.Fatal(err)
	}

	// Write .yongol with correct hash
	hash := sha256.Sum256(featContent)
	yongolContent := fmt.Sprintf("hashes:\n  features.yaml: sha256:%x\n", hash)
	if err := os.WriteFile(filepath.Join(tmp, ".yongol"), []byte(yongolContent), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := ft03HashMismatch(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}

func TestFT03_NoFeaturesYAML_NoFire(t *testing.T) {
	tmp := t.TempDir()

	// No features.yaml at all
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := ft03HashMismatch(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}

func TestFT03_EmptySpecsDir_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: ""}
	diags := ft03HashMismatch(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
