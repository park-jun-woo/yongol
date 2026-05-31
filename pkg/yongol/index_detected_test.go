//ff:func feature=orchestrator type=test control=sequence
//ff:what TestFullstack Ground/SetGround/PresenceOf 및 indexDetected 검증
package yongol

import (
	"testing"
)

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
