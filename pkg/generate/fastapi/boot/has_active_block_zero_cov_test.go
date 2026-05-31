//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasActiveBlock_ZeroCov(t *testing.T) {
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true},
		{Name: "auth", Active: false},
	}}
	if !hasActiveBlock(plan, "cors") {
		t.Error("expected cors active")
	}
	if hasActiveBlock(plan, "auth") {
		t.Error("auth should not be active")
	}
}
