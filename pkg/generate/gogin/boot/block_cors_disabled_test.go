//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockCORS — gin-contrib/cors 미들웨어 등록 (manifest + env 기반)
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockCORS_Disabled(t *testing.T) {
	block := blockCORS(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}})
	if block.Name != "cors" {
		t.Errorf("name = %q, want cors", block.Name)
	}
	if len(block.Lines) != 0 || len(block.Imports) != 0 || len(block.Funcs) != 0 {
		t.Fatalf("disabled CORS must be inert, got %+v", block)
	}
}
