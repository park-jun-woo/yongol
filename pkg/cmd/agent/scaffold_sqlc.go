//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldSQLc — DDL 기반으로 테이블별 sqlc 쿼리 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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

	systemDoc, err := docs.FS.ReadFile("sqlc.md")
	if err != nil {
		return fmt.Errorf("read sqlc.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	tableFeatMap := buildTableFeatureMap(ff)
	sorted := topoSortTables(ff.Tables)

	for _, tableName := range sorted {
		if err := scaffoldSQLcTable(specsDir, queriesDir, tableName, systemPrompt, tableFeatMap, cfg, out); err != nil {
			return err
		}
	}

	return nil
}
