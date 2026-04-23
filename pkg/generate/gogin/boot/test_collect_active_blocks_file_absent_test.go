//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectActiveBlocks_FileAbsent — file 비활성 시 file-init 블록 미포함

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCollectActiveBlocks_FileAbsent pins the same nil-deref regression
// class as BUG-008 for the file subsystem.
func TestCollectActiveBlocks_FileAbsent(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	p := prepared.New(fs)
	blocks := collectActiveBlocks(fs, p, "example.com/app")
	for _, b := range blocks {
		if b.Name == "file-init" {
			t.Fatalf("file-init must not appear when file is inactive")
		}
	}
}
