//ff:func feature=generate type=command control=sequence
//ff:what Generate — Fullstack + backend/frontend 타겟 지정으로 코드 산출물 생성 (migration 단계 포함)
package generate

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/hurl"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// MigrationHook configures the DDL migration step. Version is embedded
// in emitted headers. Logger receives human-readable status lines (ok to
// pass io.Discard / nil).
type MigrationHook struct {
	Version string
	Logger  io.Writer
}

// Generate produces backend + frontend artifacts from a parsed Fullstack
// according to the requested targets. The migration step runs before
// backend codegen so that schema changes are captured independently of
// Go code preservation.
func Generate(fs *yongol.Fullstack, artifactsDir string, backend BackendType, frontend FrontendType, hooks ...GenerateOption) error {
	cfg := &generateConfig{}
	for _, h := range hooks {
		h(cfg)
	}
	// Migration step (non-fatal when the project has no DDL).
	if fs.SpecsDir != "" {
		diags, err := runMigration(fs.SpecsDir, artifactsDir, cfg.migration)
		if err != nil {
			return fmt.Errorf("migration: %w", err)
		}
		// MIG ERRORs are already surfaced through validate; generate
		// would be refused upstream. Log warnings so users see them.
		if cfg.migration.Logger != nil {
			for _, d := range diags {
				if d.Level == diagnostic.LevelWarning {
					fmt.Fprintf(cfg.migration.Logger, "[migration] WARNING %s\n", d.Message)
				}
			}
		}
	}
	if err := runBackend(fs, artifactsDir, backend); err != nil {
		return fmt.Errorf("backend: %w", err)
	}
	if err := runFrontend(fs, artifactsDir, frontend); err != nil {
		return fmt.Errorf("frontend: %w", err)
	}
	if err := hurl.Generate(fs, artifactsDir); err != nil {
		return fmt.Errorf("hurl: %w", err)
	}
	if err := copyOPARego(fs, artifactsDir); err != nil {
		return fmt.Errorf("opa rego: %w", err)
	}
	return nil
}

type generateConfig struct {
	migration MigrationHook
}

// GenerateOption tunes Generate without breaking the existing signature.
type GenerateOption func(*generateConfig)

// WithMigration attaches the DDL migration step configuration.
func WithMigration(h MigrationHook) GenerateOption {
	return func(c *generateConfig) { c.migration = h }
}

func runMigration(specsDir, artifactsDir string, h MigrationHook) ([]diagnostic.Diagnostic, error) {
	opt := migration.Options{
		YongolVersion: h.Version,
	}
	res, diags, err := migration.Generate(specsDir, artifactsDir, opt)
	if err != nil {
		return diags, err
	}
	if h.Logger != nil {
		switch res.Mode {
		case migration.ModeInitial:
			fmt.Fprintf(h.Logger, "[migration] mode=initial file=%s ops=%d\n",
				res.MigrationFile, res.OpsCount)
		case migration.ModeIncremental:
			fmt.Fprintf(h.Logger, "[migration] mode=incremental file=%s ops=%d\n",
				res.MigrationFile, res.OpsCount)
			for _, op := range res.Operations {
				fmt.Fprintf(h.Logger, "[migration]   * %s\n", op.Description())
			}
		case migration.ModeNoop:
			fmt.Fprintln(h.Logger, "[migration] mode=noop (no schema changes)")
		}
	}
	return diags, nil
}
