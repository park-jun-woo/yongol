//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldDDL — features.yaml tables로부터 테이블별 DDL 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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
		if err := scaffoldDDLTable(specsDir, dbDir, tableName, ff, tableFeatMap, systemPrompt, cfg, out); err != nil {
			return err
		}
	}

	return nil
}
