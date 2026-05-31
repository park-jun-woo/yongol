//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBaseCandidateBlocks — main.go 후보 블록 구성 + prepared backend 분기 검증
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBaseCandidateBlocks(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}

	t.Run("NoOptionalBackends", func(t *testing.T) {
		p := prepared.New(fs)
		blocks := baseCandidateBlocks(fs, p, "example.com/app")
		names := blockNames(blocks)
		// Always-present blocks.
		for _, n := range []string{"logger-init", "router", "gin-run"} {
			if !names[n] {
				t.Errorf("expected block %q present: %v", n, names)
			}
		}
		// Optional backends absent -> their init blocks must not appear.
		for _, n := range []string{"session-init", "cache-init", "file-init", "queue-init"} {
			if names[n] {
				t.Errorf("block %q should be absent when backend inactive", n)
			}
		}
	})

	t.Run("AllOptionalBackendsActive", func(t *testing.T) {
		p := prepared.New(fs)
		p.ActiveBackends.Session = &prepared.Session{Backend: "redis"}
		p.ActiveBackends.Cache = &prepared.Cache{Backend: "redis"}
		p.ActiveBackends.File = &prepared.File{Backend: "local"}
		p.ActiveBackends.Queue = &prepared.Queue{Backend: "memory"}

		blocks := baseCandidateBlocks(fs, p, "example.com/app")
		names := blockNames(blocks)
		for _, n := range []string{"session-init", "cache-init", "file-init", "queue-init"} {
			if !names[n] {
				t.Errorf("expected active backend block %q present: %v", n, names)
			}
		}
	})
}
