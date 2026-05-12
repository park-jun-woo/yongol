//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=domain-security
//ff:what makeMultiDomainFS — multi-domain 테스트용 Fullstack 생성 (임시 OpenAPI 파일 기록)
package domain_security

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// makeMultiDomainFS creates a Fullstack configured for multi-domain testing.
// It writes minimal OpenAPI YAML files to a temp directory so that
// loadDomainOpenAPIDocs can parse them.
func makeMultiDomainFS(domains map[string]manifest.DomainConfig, opFiles map[string]string, pages []stml.PageSpec, policies []rego.Policy) *yongol.Fullstack {
	tmpDir, err := os.MkdirTemp("", "domain-security-test")
	if err != nil {
		panic(err)
	}
	// Write OpenAPI files.
	for relPath, content := range opFiles {
		absPath := filepath.Join(tmpDir, relPath)
		if err := os.MkdirAll(filepath.Dir(absPath), 0o755); err != nil {
			panic(err)
		}
		if err := os.WriteFile(absPath, []byte(content), 0o644); err != nil {
			panic(err)
		}
	}

	fs := &yongol.Fullstack{
		SpecsDir: tmpDir,
		Manifest: &manifest.ProjectConfig{
			Domains: domains,
		},
		STMLPages:      pages,
		ParsedPolicies: policies,
	}
	return fs
}
