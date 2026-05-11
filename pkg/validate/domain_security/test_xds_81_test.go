//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what XDS-81 test — internal endpoint with security declaration triggers WARNING
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDS81_Positive_InternalWithSecurity(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"internal": {OpenAPI: "api/internal.yaml"},
	}
	opFiles := map[string]string{
		"api/internal.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/internal/sync": {"post": {ID: "SyncData", Security: "[{bearerAuth: []}]"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XDS-81]") {
		t.Errorf("expected XDS-81 diagnostic, got %v", diags)
	}
	if diagLevel(diags, "[XDS-81]") != diagnostic.LevelWarning {
		t.Errorf("expected WARNING level for XDS-81, got %v", diagLevel(diags, "[XDS-81]"))
	}
}

func TestXDS81_Negative_InternalNoSecurity(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"internal": {OpenAPI: "api/internal.yaml"},
	}
	opFiles := map[string]string{
		"api/internal.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/internal/sync": {"post": {ID: "SyncData"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XDS-81]") {
		t.Errorf("unexpected XDS-81 diagnostic, got %v", diags)
	}
}
