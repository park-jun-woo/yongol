//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what test helpers — domain_security 검증 테스트용 fixture 생성 유틸
package domain_security

import (
	"os"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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

// cleanupFS removes the temporary directory used for testing.
func cleanupFS(fs *yongol.Fullstack) {
	os.RemoveAll(fs.SpecsDir)
}

// minimalOpenAPI returns minimal valid OpenAPI YAML with given operations.
// ops is a map of path -> method -> operationId.
func minimalOpenAPI(ops map[string]map[string]opDef) string {
	yaml := "openapi: '3.0.0'\ninfo:\n  title: test\n  version: '1.0'\npaths:\n"
	for path, methods := range ops {
		yaml += "  " + path + ":\n"
		for method, def := range methods {
			yaml += "    " + method + ":\n"
			yaml += "      operationId: " + def.ID + "\n"
			yaml += "      responses:\n        '200':\n          description: ok\n"
			if def.Security != "" {
				yaml += "      security: " + def.Security + "\n"
			}
		}
	}
	return yaml
}

// opDef defines an operation for test YAML generation.
type opDef struct {
	ID       string
	Security string // e.g. "[]" or "[{bearerAuth: []}]"
}

// hasDiag returns true if any diagnostic message contains the given rule prefix.
func hasDiag(diags []diagnostic.Diagnostic, prefix string) bool {
	for _, d := range diags {
		if len(d.Message) >= len(prefix) && d.Message[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

// countDiag returns the count of diagnostics whose message starts with the prefix.
func countDiag(diags []diagnostic.Diagnostic, prefix string) int {
	n := 0
	for _, d := range diags {
		if len(d.Message) >= len(prefix) && d.Message[:len(prefix)] == prefix {
			n++
		}
	}
	return n
}

// diagLevel returns the level of the first diagnostic matching the prefix.
func diagLevel(diags []diagnostic.Diagnostic, prefix string) diagnostic.Level {
	for _, d := range diags {
		if len(d.Message) >= len(prefix) && d.Message[:len(prefix)] == prefix {
			return d.Level
		}
	}
	return ""
}

// makeDocInline creates an openapi3.T using kin-openapi for in-memory tests.
func makeDocInline(paths map[string]*openapi3.PathItem) *openapi3.T {
	p := &openapi3.Paths{}
	for path, item := range paths {
		p.Set(path, item)
	}
	return &openapi3.T{Paths: p}
}
