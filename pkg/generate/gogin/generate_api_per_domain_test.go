//ff:func feature=gen-gogin type=test control=sequence
//ff:what generateAPIPerDomain — 도메인별 api_<domain> outDir/pkg 파생 + 에러 래핑 검증

package gogin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestGenerateAPIPerDomain exercises the per-domain path/package derivation
// without requiring oapi-codegen to succeed: generateOpenAPIGoGin MkdirAll's
// the outDir before invoking the (missing-spec) codegen, so the
// backend/internal/api_<sanitized> directory is created and we can assert the
// "api_" + sanitizeDomainName(name) convention. The codegen exec then fails,
// and the loop must wrap the error with the domain name.
func TestGenerateAPIPerDomain(t *testing.T) {
	specs := t.TempDir()
	arts := t.TempDir()
	fs := &yongol.Fullstack{
		SpecsDir: specs,
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				// hyphen + uppercase exercise sanitizeDomainName.
				"My-Admin": {OpenAPI: "admin/openapi.yaml"},
			},
		},
	}

	err := generateAPIPerDomain(fs, arts)
	if err == nil {
		t.Fatalf("expected error (missing spec / oapi-codegen), got nil")
	}
	if !strings.Contains(err.Error(), "My-Admin") {
		t.Errorf("error should name the domain: %v", err)
	}

	wantDir := filepath.Join(arts, "backend", "internal", "api_my_admin")
	if info, statErr := os.Stat(wantDir); statErr != nil || !info.IsDir() {
		t.Errorf("expected per-domain outDir %q created (api_+sanitize derivation); stat err=%v", wantDir, statErr)
	}
}
