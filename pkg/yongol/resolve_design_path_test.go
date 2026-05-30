//ff:func feature=orchestrator type=test control=sequence
//ff:what TestResolveDesignPath — manifest.frontend.design 우선 / convention fallback 분기 검증

package yongol

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestResolveDesignPath_ManifestOverride(t *testing.T) {
	fs := &Fullstack{Manifest: &manifest.ProjectConfig{}}
	fs.Manifest.Frontend.Design = "ui/DESIGN.md"
	got := resolveDesignPath(fs, "/root")
	want := filepath.Join("/root", "ui/DESIGN.md")
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestResolveDesignPath_Convention(t *testing.T) {
	// nil manifest → convention fallback.
	got := resolveDesignPath(&Fullstack{}, "/root")
	want := filepath.Join("/root", "frontend", "DESIGN.md")
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}

	// manifest present but Design empty → convention fallback.
	fs := &Fullstack{Manifest: &manifest.ProjectConfig{}}
	if got := resolveDesignPath(fs, "/root"); got != want {
		t.Fatalf("empty design: got %q, want %q", got, want)
	}
}
