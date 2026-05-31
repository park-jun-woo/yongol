//ff:func feature=gen-ir type=test control=sequence
//ff:what isCRUDSeq/parseDigits/splitModelMethod/planNeedsTransaction/csrfIsActive 순수 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestCsrfIsActive(t *testing.T) {
	// csrf not required -> false
	if csrfIsActive(&prepared.State{Auth: prepared.Auth{CsrfRequired: false}}) {
		t.Errorf("not required should be false")
	}
	// required but auth absent -> false
	if csrfIsActive(&prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: false}}) {
		t.Errorf("auth absent should be false")
	}
	// required, present, raw csrf nil -> default true
	ps := &prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: true, Raw: &manifest.Auth{}}}
	if !csrfIsActive(ps) {
		t.Errorf("nil csrf with cookie/hybrid should default true")
	}
	// required, present, explicit csrf.Enabled=false -> false
	ps2 := &prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: true,
		Raw: &manifest.Auth{Csrf: &manifest.CsrfConfig{Enabled: false}}}}
	if csrfIsActive(ps2) {
		t.Errorf("explicit disabled csrf should be false")
	}
}
