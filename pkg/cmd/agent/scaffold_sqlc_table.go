//ff:func feature=agent type=helper control=sequence
//ff:what scaffoldSQLcTable — 단일 테이블의 sqlc 쿼리 생성 및 파일 기록

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func scaffoldSQLcTable(specsDir, queriesDir, tableName, systemPrompt string, tableFeatMap map[string][]features.Feature, cfg Config, out io.Writer) error {
	outPath := filepath.Join(queriesDir, tableName+".sql")
	if _, err := os.Stat(outPath); err == nil {
		fmt.Fprintf(out, "  scaffold sqlc: skipped %s (exists)\n", tableName+".sql")
		return nil
	}

	ddlPath := filepath.Join(specsDir, "db", tableName+".sql")
	ddlContent, err := os.ReadFile(ddlPath)
	if err != nil {
		fmt.Fprintf(out, "  scaffold sqlc: skipped %s (DDL not found: %v)\n", tableName+".sql", err)
		return nil
	}

	feats := tableFeatMap[tableName]
	userPrompt := buildSQLcUserPrompt(tableName, string(ddlContent), feats)
	numCtx := len(systemPrompt) + len(userPrompt) + 2048

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
	return nil
}
