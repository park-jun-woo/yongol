//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseDesignIfPresent — DESIGN.md 미존재(return) + 존재 시 DesignSpec 설정
package yongol

import (
	"testing"
)

func TestParseDesignIfPresent_Missing(t *testing.T) {
	root := t.TempDir() // no frontend/DESIGN.md
	fs := &Fullstack{Presences: map[SSOTKind]SSOTPresence{}}
	parseDesignIfPresent(fs, root)
	if fs.DesignSpec != nil {
		t.Fatalf("expected no DesignSpec when DESIGN.md missing")
	}
	if _, ok := fs.Presences[KindDesign]; ok {
		t.Fatalf("expected KindDesign not marked present")
	}
}
