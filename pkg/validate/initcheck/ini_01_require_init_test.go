//ff:func feature=validate type=test control=sequence topic=init-check
//ff:what INI-01 — .yongol 존재 여부에 따른 진단 테스트

package initcheck

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestINI01_MissingYongolFile_Fires(t *testing.T) {
	tmp := t.TempDir()

	// No .yongol file in tmp
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := Run(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[INI-01]") {
		t.Errorf("want [INI-01] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "yongol init") {
		t.Errorf("want 'yongol init' in message, got %s", diags[0].Message)
	}
}

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

func TestINI01_EmptySpecsDir_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: ""}
	diags := Run(fs)
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
