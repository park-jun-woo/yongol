//ff:func feature=agent type=helper control=sequence
//ff:what scaffoldDDLTable — 단일 테이블의 DDL 생성 및 파일 기록

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func scaffoldDDLTable(specsDir, dbDir, tableName string, ff *features.FeaturesFile, tableFeatMap map[string][]features.Feature, systemPrompt string, cfg Config, out io.Writer) error {
	outPath := filepath.Join(dbDir, tableName+".sql")
	if _, err := os.Stat(outPath); err == nil {
		fmt.Fprintf(out, "  scaffold ddl: skipped %s (exists)\n", tableName+".sql")
		return nil
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
	return nil
}
