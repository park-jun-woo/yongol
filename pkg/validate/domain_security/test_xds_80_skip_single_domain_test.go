//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDS80_Skip_SingleDomain — 단일 도메인 프로젝트 시 skip

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDS80_Skip_SingleDomain(t *testing.T) {
	// No domains key → skip all rules.
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{},
	}
	diags := Run(fs)
	if len(diags) > 0 {
		t.Errorf("expected no diagnostics for single-domain project, got %v", diags)
	}
}
