//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectActiveBlocks_CacheAbsent — cache 비활성 시 cache-init 블록 미포함

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCollectActiveBlocks_CacheAbsent pins the same nil-deref regression
// class as BUG-008 for the cache subsystem.
func TestCollectActiveBlocks_CacheAbsent(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	p := prepared.New(fs)
	blocks := collectActiveBlocks(fs, p, "example.com/app")
	for _, b := range blocks {
		if b.Name == "cache-init" {
			t.Fatalf("cache-init must not appear when cache is inactive")
		}
	}
}
