//ff:func feature=validate type=test control=sequence topic=init-check
//ff:what TestINI01_YongolFileExists_NoFire

package initcheck

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestINI01_YongolFileExists_NoFire(t *testing.T) {
	tmp := t.TempDir()

	// Create .yongol file
	if err := os.WriteFile(filepath.Join(tmp, ".yongol"), []byte("hashes: {}"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := Run(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
