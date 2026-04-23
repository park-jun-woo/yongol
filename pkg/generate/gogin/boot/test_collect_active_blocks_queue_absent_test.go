//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectActiveBlocks_QueueAbsent — queue 비활성 시 queue-init 블록 미포함

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCollectActiveBlocks_QueueAbsent pins the same nil-deref regression
// class as BUG-008 for the queue subsystem.
func TestCollectActiveBlocks_QueueAbsent(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	p := prepared.New(fs)
	blocks := collectActiveBlocks(fs, p, "example.com/app")
	for _, b := range blocks {
		if b.Name == "queue-init" {
			t.Fatalf("queue-init must not appear when queue is inactive")
		}
	}
}
