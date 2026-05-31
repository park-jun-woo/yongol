//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what TestSSaCDDLHelpers — unit tests for the pure ssac_ddl helper functions
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestIsAuthRequiredTable(t *testing.T) {
	// Backend.Auth configured → refresh_tokens is auth-required.
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	fs.Manifest.Backend.Auth = &manifest.Auth{}
	if !isAuthRequiredTable(fs, "refresh_tokens") {
		t.Error("refresh_tokens should be auth-required when auth configured")
	}
	if isAuthRequiredTable(fs, "courses") {
		t.Error("courses should not be auth-required")
	}
	// no manifest → false.
	if isAuthRequiredTable(&yongol.Fullstack{}, "refresh_tokens") {
		t.Error("nil manifest → not auth-required")
	}
	// manifest without Backend.Auth → false.
	if isAuthRequiredTable(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}, "refresh_tokens") {
		t.Error("no Backend.Auth → not auth-required")
	}
}
