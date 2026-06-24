//ff:func feature=gen-gogin type=util control=sequence
//ff:what generateAPICode — IsDomained 분기로 단일/도메인별 oapi-codegen+splitter 디스패치
package gogin

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/splitter"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateAPICode runs oapi-codegen (strict-server) and the splitter for the
// API layer. In domain mode it delegates to generateAPIPerDomain (one package
// per domain). In single-site mode it produces the unchanged
// backend/internal/api package with -package "api".
func generateAPICode(fs *yongol.Fullstack, artifactsDir string) error {
	if fs.IsDomained() {
		return generateAPIPerDomain(fs, artifactsDir)
	}
	oapiPath := filepath.Join(fs.SpecsDir, "api", "openapi.yaml")
	apiDir := filepath.Join(artifactsDir, "backend", "internal", "api")
	if err := generateOpenAPIGoGin(oapiPath, apiDir, "api"); err != nil {
		return fmt.Errorf("oapi-codegen strict-server: %w", err)
	}
	if err := splitter.SplitDirectory(apiDir, splitter.ToolOAPICodegen); err != nil {
		return fmt.Errorf("split oapi-codegen output: %w", err)
	}
	return nil
}
