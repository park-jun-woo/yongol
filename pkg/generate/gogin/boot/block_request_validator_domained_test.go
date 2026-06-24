//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBlockRequestValidator_DomainModeSkips — 도메인 모드에서 전역 validator 블록 미방출 (BUG-142)

package boot

import "testing"

func TestBlockRequestValidator_DomainModeSkips(t *testing.T) {
	// BUG-142 — domain mode emits per-domain validators on route groups
	// (appendDomainHandler → group.Use), so the global block contributes nothing
	// (no global r.Use(validator)).
	block := blockRequestValidator(domainedFS(nil), "example.com/app")
	if len(block.Lines) != 0 || len(block.Imports) != 0 {
		t.Errorf("domain mode must emit no global validator, got lines=%v imports=%v",
			block.Lines, block.Imports)
	}
}
