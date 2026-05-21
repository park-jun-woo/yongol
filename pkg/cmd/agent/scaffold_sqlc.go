//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldSQLc — DDL 기반으로 테이블별 sqlc 쿼리 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// scaffoldSQLc generates sqlc query files for each table defined in features.yaml.
// Must be called after scaffoldDDL so that DDL files exist.
// Existing files are skipped to protect user modifications.
func scaffoldSQLc(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) error {
	if len(ff.Tables) == 0 {
		return nil
	}

	queriesDir := filepath.Join(specsDir, "db", "queries")
	if err := os.MkdirAll(queriesDir, 0755); err != nil {
		return fmt.Errorf("create db/queries dir: %w", err)
	}

	// Load sqlc docs for system prompt
	systemDoc, err := docs.FS.ReadFile("sqlc.md")
	if err != nil {
		return fmt.Errorf("read sqlc.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	// Build feature lookup by table name
	tableFeatMap := buildTableFeatureMap(ff)

	sorted := topoSortTables(ff.Tables)

	for _, tableName := range sorted {
		outPath := filepath.Join(queriesDir, tableName+".sql")
		if _, err := os.Stat(outPath); err == nil {
			fmt.Fprintf(out, "  scaffold sqlc: skipped %s (exists)\n", tableName+".sql")
			continue
		}

		// Read the DDL file for this table
		ddlPath := filepath.Join(specsDir, "db", tableName+".sql")
		ddlContent, err := os.ReadFile(ddlPath)
		if err != nil {
			fmt.Fprintf(out, "  scaffold sqlc: skipped %s (DDL not found: %v)\n", tableName+".sql", err)
			continue
		}

		feats := tableFeatMap[tableName]
		userPrompt := buildSQLcUserPrompt(tableName, string(ddlContent), feats)

		// Dynamic num_ctx: (len(system+user) / 4) * 1.5 + 2048
		numCtx := int(float64(len(systemPrompt)+len(userPrompt))/4*1.5) + 2048

		reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
		if err != nil {
			return fmt.Errorf("scaffold sqlc %s: %w", tableName, err)
		}

		content := stripCodeBlock(reply)
		if content == "" {
			return fmt.Errorf("scaffold sqlc %s: empty LLM response", tableName)
		}

		if err := os.WriteFile(outPath, []byte(content+"\n"), 0644); err != nil {
			return fmt.Errorf("scaffold sqlc %s: write: %w", tableName, err)
		}

		fmt.Fprintf(out, "  scaffold sqlc: created %s\n", tableName+".sql")
	}

	return nil
}

// buildSQLcUserPrompt builds the user prompt for generating sqlc queries.
func buildSQLcUserPrompt(tableName, ddlContent string, feats []features.Feature) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Table: %s\n\n", tableName)
	fmt.Fprintf(&b, "DDL:\n%s\n\n", ddlContent)

	if len(feats) > 0 {
		b.WriteString("Related features with cardinality hints:\n")
		for _, f := range feats {
			hint := cardinalityHint(f.Op, f.Path)
			fmt.Fprintf(&b, "  - %s %s: %s (cardinality: %s)\n", f.Op, f.Path, f.Desc, hint)
		}
		b.WriteByte('\n')
	}

	b.WriteString("Generate sqlc-compatible SQL queries for this table.")
	b.WriteString("\nOutput ONLY the SQL. No explanations. No markdown fences.")

	return b.String()
}

// cardinalityHint infers the sqlc cardinality annotation from the operation and path.
func cardinalityHint(op, path string) string {
	// Extract HTTP method from operationId naming convention
	opLower := strings.ToLower(op)

	switch {
	case strings.HasPrefix(opLower, "get") || strings.HasPrefix(opLower, "list"):
		// GET /{id} -> :one, GET / -> :many
		if strings.Contains(path, "{") {
			return ":one"
		}
		return ":many"
	case strings.HasPrefix(opLower, "create") || strings.HasPrefix(opLower, "post"):
		return ":one RETURNING"
	case strings.HasPrefix(opLower, "update") || strings.HasPrefix(opLower, "put"):
		return ":one RETURNING"
	case strings.HasPrefix(opLower, "delete") || strings.HasPrefix(opLower, "remove"):
		return ":exec"
	default:
		return ":one"
	}
}
