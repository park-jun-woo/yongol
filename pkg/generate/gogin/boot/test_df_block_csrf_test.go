//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestBlockCsrf_BearerMode_Dormant — bearer 모드에서는 inert block 으로 나와야 한다

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_BearerMode_Dormant(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Auth:   &pmanifest.Auth{Mode: "bearer"},
			},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("bearer mode should yield inert block, got lines: %+v", block.Lines)
	}
	if block.Active == nil || block.Active(fs) {
		t.Fatalf("bearer mode should report Active()=false")
	}
}
