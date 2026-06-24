//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=domain-security
//ff:what makeMultiDomainFS — multi-domain 테스트용 Fullstack 생성 (임시 OpenAPI 파일 기록 + DomainOpenAPIDocs 사전 파싱)
package domain_security

import (
	"os"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// makeMultiDomainFS creates a Fullstack configured for multi-domain testing.
// It writes minimal OpenAPI YAML files to a temp directory and pre-parses each
// into fs.DomainOpenAPIDocs (mirroring ParseAll's Phase004 domain loop), which
// is what loadDomainOpenAPIDocs now sources from instead of re-parsing on every
// rule call.
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

	// Pre-parse each domain's OpenAPI into DomainOpenAPIDocs (Phase004).
	domainDocs := make(map[string]*openapi3.T)
	for name, cfg := range domains {
		if cfg.OpenAPI == "" {
			continue
		}
		doc, err := openapi3.NewLoader().LoadFromFile(filepath.Join(tmpDir, cfg.OpenAPI))
		if err != nil {
			continue
		}
		domainDocs[name] = doc
	}

	fs := &yongol.Fullstack{
		SpecsDir: tmpDir,
		Manifest: &manifest.ProjectConfig{
			Domains: domains,
		},
		STMLPages:         pages,
		ParsedPolicies:    policies,
		DomainOpenAPIDocs: domainDocs,
	}
	return fs
}
