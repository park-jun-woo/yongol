//ff:func feature=generate type=command control=sequence
//ff:what Generate — Fullstack + backend/frontend 타겟 지정으로 코드 산출물 생성 (migration 단계 포함)
package generate

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/hurl_mirror"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces backend + frontend artifacts from a parsed Fullstack
// according to the requested targets. The migration step runs before
// backend codegen so that schema changes are captured independently of
// Go code preservation.
func Generate(fs *yongol.Fullstack, artifactsDir string, backend BackendType, frontend FrontendType, hooks ...GenerateOption) error {
	cfg := &generateConfig{}
	applyGenerateOptions(cfg, hooks)
	if err := runMigrationStep(fs, artifactsDir, cfg); err != nil {
		return err
	}
	if err := runBackend(fs, artifactsDir, backend); err != nil {
		return fmt.Errorf("backend: %w", err)
	}
	if err := runFrontend(fs, artifactsDir, frontend); err != nil {
		return fmt.Errorf("frontend: %w", err)
	}
	if fs != nil {
		if _, err := hurl_mirror.MirrorSpecsTests(fs.SpecsDir, artifactsDir); err != nil {
			return fmt.Errorf("hurl mirror: %w", err)
		}
	}
	if err := copyOPARego(fs, artifactsDir); err != nil {
		return fmt.Errorf("opa rego: %w", err)
	}
	return nil
}
