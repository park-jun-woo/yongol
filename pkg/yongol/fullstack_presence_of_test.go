//ff:func feature=orchestrator type=test control=sequence
//ff:what TestFullstack Ground/SetGround/PresenceOf 및 indexDetected 검증
package yongol

import (
	"testing"
)

func TestFullstackPresenceOf(t *testing.T) {
	// nil receiver and nil map both yield SSOTAbsent.
	var nilFS *Fullstack
	if got := nilFS.PresenceOf(KindOpenAPI); got != SSOTAbsent {
		t.Errorf("nil Fullstack PresenceOf = %v, want SSOTAbsent", got)
	}
	if got := (&Fullstack{}).PresenceOf(KindOpenAPI); got != SSOTAbsent {
		t.Errorf("nil Presences PresenceOf = %v, want SSOTAbsent", got)
	}

	fs := &Fullstack{Presences: map[SSOTKind]SSOTPresence{
		KindOpenAPI: SSOTPopulated,
		KindDDL:     SSOTDeclared,
	}}
	if got := fs.PresenceOf(KindOpenAPI); got != SSOTPopulated {
		t.Errorf("PresenceOf(OpenAPI) = %v, want SSOTPopulated", got)
	}
	if got := fs.PresenceOf(KindDDL); got != SSOTDeclared {
		t.Errorf("PresenceOf(DDL) = %v, want SSOTDeclared", got)
	}
	// Unregistered kind => Absent.
	if got := fs.PresenceOf(KindSSaC); got != SSOTAbsent {
		t.Errorf("PresenceOf(unregistered) = %v, want SSOTAbsent", got)
	}
}
