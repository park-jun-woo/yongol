//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestBlockCsrf_NoAuthBlock_Dormant — auth 블록 자체 없을 때 inert block

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_NoAuthBlock_Dormant(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("missing auth should yield inert block, got lines: %+v", block.Lines)
	}
	if block.Active == nil || block.Active(fs) {
		t.Fatalf("missing auth should report Active()=false")
	}
}
