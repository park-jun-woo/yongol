//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestBlockCsrf_NoAuthBlock_Dormant — auth 블록 자체 없을 때 inert block

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockCsrf_NoAuthBlock_Dormant(t *testing.T) {
	a := prepared.Auth{}
	block := blockCsrf(nil, a, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Fatalf("missing auth should yield inert block, got lines: %+v", block.Lines)
	}
	if block.Active == nil || block.Active(&yongol.Fullstack{}) {
		t.Fatalf("missing auth should report Active()=false via csrfAlwaysInactive")
	}
}
