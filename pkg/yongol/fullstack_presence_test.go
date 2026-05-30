//ff:func feature=orchestrator type=test control=sequence
//ff:what TestFullstack Ground/SetGround/PresenceOf 및 indexDetected 검증

package yongol

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestFullstackGroundSetGet(t *testing.T) {
	fs := &Fullstack{}
	if fs.Ground() != nil {
		t.Error("expected nil Ground before SetGround")
	}
	g := &rule.Ground{}
	fs.SetGround(g)
	if fs.Ground() != g {
		t.Error("Ground() did not return the bound rule.Ground")
	}
}

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

func TestIndexDetected(t *testing.T) {
	fs := &Fullstack{Presences: map[SSOTKind]SSOTPresence{}}
	detected := []DetectedSSOT{
		{Kind: KindOpenAPI, Path: "/a/openapi.yaml", Presence: SSOTPopulated},
		{Kind: KindDDL, Path: "/a/ddl", Presence: SSOTDeclared},
	}
	has := indexDetected(detected, fs)

	if len(has) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(has))
	}
	if has[KindOpenAPI].Path != "/a/openapi.yaml" {
		t.Errorf("OpenAPI path = %q", has[KindOpenAPI].Path)
	}
	// Presences must be mirrored into fs.
	if fs.Presences[KindOpenAPI] != SSOTPopulated {
		t.Errorf("fs.Presences[OpenAPI] = %v, want SSOTPopulated", fs.Presences[KindOpenAPI])
	}
	if fs.Presences[KindDDL] != SSOTDeclared {
		t.Errorf("fs.Presences[DDL] = %v, want SSOTDeclared", fs.Presences[KindDDL])
	}
}
