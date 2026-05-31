//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what TestBlockOtelInit_Disabled — tracing 미활성 시 inert block

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockOtelInit_Disabled(t *testing.T) {
	// No tracing block → inert (no imports, no lines). Ensures non-tracing
	// projects do not pay OTel dependency / runtime cost.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockOtelInit(fs, "example.com/zenflow")
	if len(block.Lines) != 0 || len(block.Imports) != 0 {
		t.Fatalf("expected inert block when tracing disabled, got Lines=%d Imports=%d", len(block.Lines), len(block.Imports))
	}
}
