//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldDDL — features.yaml tables로부터 테이블별 DDL 자동 생성

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

// scaffoldDDL generates DDL files for each table defined in features.yaml.
// Tables are processed in topological order (parents first via belongs_to).
// Existing files are skipped to protect user modifications.
func scaffoldDDL(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) error {
	if len(ff.Tables) == 0 {
		return nil
	}

	dbDir := filepath.Join(specsDir, "db")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("create db dir: %w", err)
	}

	// Load DDL docs for system prompt
	systemDoc, err := docs.FS.ReadFile("ddl.md")
	if err != nil {
		return fmt.Errorf("read ddl.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	// Build feature lookup by table name
	tableFeatMap := buildTableFeatureMap(ff)

	// Process tables in topological order
	sorted := topoSortTables(ff.Tables)

	for _, tableName := range sorted {
		outPath := filepath.Join(dbDir, tableName+".sql")
		if _, err := os.Stat(outPath); err == nil {
			fmt.Fprintf(out, "  scaffold ddl: skipped %s (exists)\n", tableName+".sql")
			continue
		}

		td := ff.Tables[tableName]
		userPrompt := buildDDLUserPrompt(tableName, td, tableFeatMap[tableName])

		numCtx := len(systemPrompt) + len(userPrompt) + 2048

		reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
		if err != nil {
			return fmt.Errorf("scaffold ddl %s: %w", tableName, err)
		}

		content := stripCodeBlock(reply)
		if content == "" {
			return fmt.Errorf("scaffold ddl %s: empty LLM response", tableName)
		}

		if err := os.WriteFile(outPath, []byte(content+"\n"), 0644); err != nil {
			return fmt.Errorf("scaffold ddl %s: write: %w", tableName, err)
		}

		fmt.Fprintf(out, "  scaffold ddl: created %s\n", tableName+".sql")
	}

	return nil
}

// buildDDLUserPrompt builds the user prompt for generating a DDL file.
func buildDDLUserPrompt(tableName string, td features.TableDef, feats []features.Feature) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Table: %s\n", tableName)

	if len(td.BelongsTo) > 0 {
		fmt.Fprintf(&b, "belongs_to: %s\n", strings.Join(td.BelongsTo, ", "))
	}
	if len(td.HasMany) > 0 {
		fmt.Fprintf(&b, "has_many: %s\n", strings.Join(td.HasMany, ", "))
	}
	if len(td.States) > 0 {
		fmt.Fprintf(&b, "states: %s\n", strings.Join(td.States, ", "))
	}

	if len(feats) > 0 {
		b.WriteString("\nRelated features:\n")
		for _, f := range feats {
			fmt.Fprintf(&b, "  - %s %s: %s\n", f.Op, f.Path, f.Desc)
		}
	}

	b.WriteString("\nGenerate a PostgreSQL CREATE TABLE statement for this table.")
	b.WriteString("\nOutput ONLY the SQL. No explanations. No markdown fences.")

	return b.String()
}

// buildTableFeatureMap groups features by their table field.
func buildTableFeatureMap(ff *features.FeaturesFile) map[string][]features.Feature {
	m := make(map[string][]features.Feature)
	for _, f := range ff.Features {
		if f.Table != "" {
			m[f.Table] = append(m[f.Table], f)
		}
	}
	return m
}
