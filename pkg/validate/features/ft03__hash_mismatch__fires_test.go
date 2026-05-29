//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT03_HashMismatch_Fires

package features

import (
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
