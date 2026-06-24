//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what generateAPIPerDomain — 도메인마다 oapi-codegen+splitter 를 api_<domain> 패키지로 실행
package gogin

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/splitter"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateAPIPerDomain runs oapi-codegen + the splitter once per manifest
// domain. Each domain's disk spec (cfg.OpenAPI under fs.SpecsDir) is rendered
// into its own backend/internal/api_<domain> package. oapi-codegen reads the
// spec from disk via the path, so DomainView (in-memory fs.OpenAPIDoc) is not
// involved here.
func generateAPIPerDomain(fs *yongol.Fullstack, artifactsDir string) error {
	for name, cfg := range fs.Manifest.Domains {
		oapiPath := filepath.Join(fs.SpecsDir, cfg.OpenAPI)
		pkgName := "api_" + sanitizeDomainName(name)
		outDir := filepath.Join(artifactsDir, "backend", "internal", pkgName)
		if err := generateOpenAPIGoGin(oapiPath, outDir, pkgName); err != nil {
			return fmt.Errorf("oapi-codegen strict-server (%s): %w", name, err)
		}
		if err := splitter.SplitDirectory(outDir, splitter.ToolOAPICodegen); err != nil {
			return fmt.Errorf("split oapi-codegen output (%s): %w", name, err)
		}
	}
	return nil
}
