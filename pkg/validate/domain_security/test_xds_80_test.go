//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what XDS-80 test — admin endpoint with security: [] triggers ERROR
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDS80_Positive_AdminPublicAccess(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers", Security: "[]"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XDS-80]") {
		t.Errorf("expected XDS-80 diagnostic, got %v", diags)
	}
	if diagLevel(diags, "[XDS-80]") != diagnostic.LevelError {
		t.Errorf("expected ERROR level for XDS-80, got %v", diagLevel(diags, "[XDS-80]"))
	}
}

func TestXDS80_Negative_AdminWithSecurity(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers", Security: "[{bearerAuth: []}]"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XDS-80]") {
		t.Errorf("unexpected XDS-80 diagnostic, got %v", diags)
	}
}

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
