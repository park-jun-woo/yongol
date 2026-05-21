//ff:func feature=agent type=command control=sequence
//ff:what scaffold — SSOT scaffold 오케스트레이션 (DDL + sqlc + OpenAPI + SSaC + Rego + stateDiagram)

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// scaffold orchestrates SSOT file generation from features.yaml.
// Generates in order: DDL → sqlc → OpenAPI → SSaC → Rego → stateDiagram.
// Existing files are never overwritten (user modifications are protected).
func scaffold(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) error {
	if ff == nil || len(ff.Tables) == 0 {
		fmt.Fprintf(out, "\nyongol agent: scaffold skipped (no tables in features.yaml)\n")
		return nil
	}

	nTables := len(ff.Tables)
	fmt.Fprintf(out, "\nyongol agent: scaffold — %d tables\n", nTables)

	// Phase 1: DDL
	if err := scaffoldDDL(specsDir, ff, llmFn, cfg, out); err != nil {
		return fmt.Errorf("scaffold DDL: %w", err)
	}

	// Phase 2: sqlc queries (depends on DDL files)
	if err := scaffoldSQLc(specsDir, ff, llmFn, cfg, out); err != nil {
		return fmt.Errorf("scaffold sqlc: %w", err)
	}

	// Phase 3: OpenAPI (depends on DDL files)
	nPaths, err := scaffoldOpenAPI(specsDir, ff, llmFn, cfg, out)
	if err != nil {
		return fmt.Errorf("scaffold OpenAPI: %w", err)
	}

	// Read generated openapi.yaml for SSaC prompts
	openapiContent := ""
	openapiPath := filepath.Join(specsDir, "api", "openapi.yaml")
	if data, err := os.ReadFile(openapiPath); err == nil {
		openapiContent = string(data)
	}

	// Phase 4: SSaC (depends on DDL, sqlc, OpenAPI)
	nSSaC, err := scaffoldSSaC(specsDir, ff, openapiContent, llmFn, cfg, out)
	if err != nil {
		return fmt.Errorf("scaffold SSaC: %w", err)
	}

	// Phase 5: Rego (depends on features list)
	if err := scaffoldRego(specsDir, ff, llmFn, cfg, out); err != nil {
		return fmt.Errorf("scaffold Rego: %w", err)
	}

	// Phase 6: stateDiagram (depends on tables with states)
	nStates, err := scaffoldStateMachine(specsDir, ff, llmFn, cfg, out)
	if err != nil {
		return fmt.Errorf("scaffold stateDiagram: %w", err)
	}

	// Count sqlc query files
	nQueries := 0
	queriesDir := filepath.Join(specsDir, "db", "queries")
	if entries, err := os.ReadDir(queriesDir); err == nil {
		for _, e := range entries {
			if !e.IsDir() {
				nQueries++
			}
		}
	}

	fmt.Fprintf(out, "yongol agent: scaffold complete (%d tables, %d queries, %d paths, %d ssac, 1 rego, %d states)\n",
		nTables, nQueries, nPaths, nSSaC, nStates)
	return nil
}
