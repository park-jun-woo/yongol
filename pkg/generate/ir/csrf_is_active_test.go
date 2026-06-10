//ff:func feature=gen-ir type=test control=sequence
//ff:what TestCsrfIsActive — auth 선언 기반 CSRF 방출 게이트 검증 (BUG-116: bearer 포함)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestCsrfIsActive(t *testing.T) {
	// auth absent -> false (no session to protect)
	if csrfIsActive(&prepared.State{Auth: prepared.Auth{Present: false}}) {
		t.Errorf("auth absent should be false")
	}
	// present but raw nil -> false
	if csrfIsActive(&prepared.State{Auth: prepared.Auth{Present: true, Raw: nil}}) {
		t.Errorf("nil raw should be false")
	}
	// present, raw csrf nil -> default true (cookie/hybrid build)
	ps := &prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: true, Raw: &manifest.Auth{}}}
	if !csrfIsActive(ps) {
		t.Errorf("nil csrf with auth present should default true")
	}
	// BUG-116 / Phase-B1 — bearer build (CsrfRequired=false) still emits CSRF
	// because BACKEND_AUTH_MODE can reach cookie/hybrid at runtime.
	bearer := &prepared.State{Auth: prepared.Auth{CsrfRequired: false, Present: true, Mode: "bearer", Raw: &manifest.Auth{}}}
	if !csrfIsActive(bearer) {
		t.Errorf("bearer build with auth present should emit CSRF (runtime-reachable)")
	}
	// present, explicit csrf.Enabled=false -> false (operator opt-out)
	ps2 := &prepared.State{Auth: prepared.Auth{CsrfRequired: true, Present: true,
		Raw: &manifest.Auth{Csrf: &manifest.CsrfConfig{Enabled: false}}}}
	if csrfIsActive(ps2) {
		t.Errorf("explicit disabled csrf should be false")
	}
}
