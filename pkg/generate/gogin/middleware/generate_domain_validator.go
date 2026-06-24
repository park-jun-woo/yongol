//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateDomainValidators — 도메인별 openapi_<ident>.yaml 복사 + request_validator_<ident>.go 방출

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateDomainValidators emits, for each manifest domain (BUG-142), that
// domain's request-validator: it copies the domain's OpenAPI spec
// (filepath.Join(fs.SpecsDir, cfg.OpenAPI), matching parse_domains_if_present.go:36)
// to internal/middleware/openapi_<ident>.yaml for go:embed and writes
// request_validator_<ident>.go exposing RequestValidator<Title>(). The
// single-site request_validator.go / openapi.yaml are deliberately NOT written
// in domain mode — each domain validates only against its own contract on its
// own route group (boot.appendDomainHandler wires group.Use). Domain-suffixed
// filenames and identifiers avoid collisions in the shared middleware package.
func generateDomainValidators(fs *yongol.Fullstack, mwDir string) error {
	for _, name := range fs.DomainNames() {
		cfg := fs.Manifest.Domains[name]
		ident := domainIdent(name)
		src := filepath.Join(fs.SpecsDir, cfg.OpenAPI)
		dst := filepath.Join(mwDir, "openapi_"+ident+".yaml")
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("copy openapi %s: %w", name, err)
		}
		goFile := filepath.Join(mwDir, "request_validator_"+ident+".go")
		if err := os.WriteFile(goFile, []byte(domainRequestValidatorSource(name)), 0o644); err != nil {
			return fmt.Errorf("write validator %s: %w", name, err)
		}
	}
	return nil
}
