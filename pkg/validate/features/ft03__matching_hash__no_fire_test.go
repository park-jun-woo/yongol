//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what TestFT03_MatchingHash_NoFire

package features

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
