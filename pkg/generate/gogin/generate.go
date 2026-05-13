//ff:func feature=gen-gogin type=command control=sequence
//ff:what Generate — Fullstack에서 Go+Gin 백엔드 코드 생성
package gogin

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/filefunc"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/auth"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/boot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/infra"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/middleware"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/sqlcpost"
	ssacgen "github.com/park-jun-woo/yongol/pkg/generate/gogin/ssac"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/state"
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/generate/splitter"
)

// Generate produces Go+Gin backend artifacts from a parsed Fullstack.
// Pipeline: oapi-codegen(strict-server) → sqlc → auth → state →
// boot → ssac → funcCopy → gomod.
// SQL queries (including pagination) are user-authored — yongol validate checks correctness.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	p := prepared.New(fs)
	if err := generateOpenAPIGoGin(fs.SpecsDir, artifactsDir); err != nil {
		return fmt.Errorf("oapi-codegen strict-server: %w", err)
	}
	apiDir := filepath.Join(artifactsDir, "backend", "internal", "api")
	if err := splitter.SplitDirectory(apiDir, splitter.ToolOAPICodegen); err != nil {
		return fmt.Errorf("split oapi-codegen output: %w", err)
	}
	if err := checkSqlcOutPath(fs.SpecsDir, artifactsDir); err != nil {
		return fmt.Errorf("sqlc.yaml validation: %w", err)
	}
	if err := generateSQLCGoGin(fs.SpecsDir); err != nil {
		return fmt.Errorf("sqlc: %w", err)
	}
	dbDir := filepath.Join(artifactsDir, "backend", "internal", "db")
	if err := splitter.SplitDirectory(dbDir, splitter.ToolSQLC); err != nil {
		return fmt.Errorf("split sqlc output: %w", err)
	}
	if err := sqlcpost.Generate(fs, artifactsDir); err != nil {
		return fmt.Errorf("sqlc post (log masking): %w", err)
	}
	// Phase002 (ssac/purify) — emit postgres adapters for every ssac package
	// that declares an interface.yaml with at least one active port under the
	// current manifest. The generated adapters live under
	// `backend/internal/infra/<pkg>/postgres.go` and are wired by boot/
	// block_infra_wiring.go below.
	if err := infra.EmitAll(fs, artifactsDir, fs.Manifest.Backend.Module); err != nil {
		return fmt.Errorf("infra postgres adapters: %w", err)
	}
	if err := auth.Generate(fs, p, artifactsDir); err != nil {
		return fmt.Errorf("auth: %w", err)
	}
	if err := state.Generate(fs, artifactsDir); err != nil {
		return fmt.Errorf("statemachine: %w", err)
	}
	if err := middleware.Generate(fs, p, artifactsDir); err != nil {
		return fmt.Errorf("middleware: %w", err)
	}
	if err := boot.Generate(fs, p, artifactsDir); err != nil {
		return fmt.Errorf("main.go: %w", err)
	}
	if err := ssacgen.Generate(fs, artifactsDir); err != nil {
		return fmt.Errorf("ssac service: %w", err)
	}
	if err := copyFuncSpecs(fs.SpecsDir, artifactsDir); err != nil {
		return fmt.Errorf("func copy: %w", err)
	}
	if err := generateGoMod(fs, fs.Manifest.Backend.Module, artifactsDir); err != nil {
		return fmt.Errorf("go.mod: %w", err)
	}
	backendDir := filepath.Join(artifactsDir, "backend")
	if err := injectFFChecked(backendDir, fs.SpecsDir); err != nil {
		return fmt.Errorf("inject //ff:checked: %w", err)
	}
	if err := filefunc.Generate(fs, artifactsDir); err != nil {
		return fmt.Errorf("filefunc codebook: %w", err)
	}
	if err := generateFFIgnore(artifactsDir); err != nil {
		return fmt.Errorf("ffignore: %w", err)
	}
	return nil
}
