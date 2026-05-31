//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestHasActiveBlock_ZeroCov — 활성 블록 존재 여부
package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasActiveBlock_ZeroCov(t *testing.T) {
	plan := &ir.BootPlan{ActiveBlocks: []ir.BootBlock{
		{Name: "cors", Active: true},
		{Name: "session", Active: false},
	}}
	if !hasActiveBlock(plan, "cors") {
		t.Error("expected cors active")
	}
	if hasActiveBlock(plan, "session") {
		t.Error("session should not be active")
	}
	if hasActiveBlock(plan, "missing") {
		t.Error("missing should not be active")
	}
}
