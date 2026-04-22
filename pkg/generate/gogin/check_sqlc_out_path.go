//ff:func feature=gen-gogin type=validator control=sequence
//ff:what checkSqlcOutPath — sqlc.yaml out 경로가 <artifacts>/backend/internal/db 인지 검증

package gogin

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// checkSqlcOutPath verifies that every `sql[].gen.go.out` in specsDir/db/sqlc.yaml
// resolves to the same absolute path as artifactsDir/backend/internal/db.
// sqlc resolves relative `out:` paths against the config-file directory (specs/db),
// so running `yongol generate` with an artifactsDir that differs from what
// sqlc.yaml expects would write sqlc output to the wrong location and leave
// subsequent generator stages referencing types that were never produced.
//
// Returning an error instead of silently rewriting the user's sqlc.yaml keeps
// the failure mode explicit: the user fixes their config once, rather than
// chasing intermittent "undefined: X" build errors.
func checkSqlcOutPath(specsDir, artifactsDir string) error {
	dbDir := filepath.Join(specsDir, "db")
	yamlPath := filepath.Join(dbDir, "sqlc.yaml")

	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", yamlPath, err)
	}

	var cfg struct {
		SQL []struct {
			Gen struct {
				Go struct {
					Out string `yaml:"out"`
				} `yaml:"go"`
			} `yaml:"gen"`
		} `yaml:"sql"`
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("parse %s: %w", yamlPath, err)
	}

	absArtifacts, err := filepath.Abs(artifactsDir)
	if err != nil {
		return fmt.Errorf("resolve artifacts path: %w", err)
	}
	absDB, err := filepath.Abs(dbDir)
	if err != nil {
		return fmt.Errorf("resolve db dir: %w", err)
	}
	expected := filepath.Join(absArtifacts, "backend", "internal", "db")

	if len(cfg.SQL) == 0 {
		return fmt.Errorf("%s: no sql entries found", yamlPath)
	}

	for i, entry := range cfg.SQL {
		out := entry.Gen.Go.Out
		if out == "" {
			return fmt.Errorf("%s: sql[%d].gen.go.out is empty", yamlPath, i)
		}
		resolved := out
		if !filepath.IsAbs(resolved) {
			resolved = filepath.Join(absDB, out)
		}
		resolved = filepath.Clean(resolved)
		if resolved != expected {
			return fmt.Errorf(
				"%s: sql[%d].gen.go.out = %q resolves to %s but expected %s\n"+
					"→ yongol generate writes backend artifacts under <artifacts>/backend/internal/db.\n"+
					"  Update sqlc.yaml `out:` to match the artifacts path you pass to `yongol generate`.",
				yamlPath, i, out, resolved, expected,
			)
		}
	}
	return nil
}
