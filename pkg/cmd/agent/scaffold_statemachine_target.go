//ff:func feature=agent type=helper control=sequence
//ff:what scaffoldStateMachineTarget — 단일 테이블의 state diagram 생성 및 파일 기록

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func scaffoldStateMachineTarget(statesDir, tableName string, states []string, feats []features.Feature, systemPrompt string, cfg Config, out io.Writer) (bool, error) {
	outPath := filepath.Join(statesDir, tableName+".md")
	if _, err := os.Stat(outPath); err == nil {
		fmt.Fprintf(out, "  scaffold states: skipped %s.md (exists)\n", tableName)
		return false, nil
	}

	userPrompt := buildStateMachineUserPrompt(tableName, states, feats)
	numCtx := len(systemPrompt) + len(userPrompt) + 2048

	reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
	if err != nil {
		return false, fmt.Errorf("scaffold states %s: %w", tableName, err)
	}

	content := strings.TrimSpace(reply)
	if content == "" {
		return false, fmt.Errorf("scaffold states %s: empty LLM response", tableName)
	}

	if err := os.WriteFile(outPath, []byte(content+"\n"), 0644); err != nil {
		return false, fmt.Errorf("scaffold states %s: write: %w", tableName, err)
	}

	fmt.Fprintf(out, "  scaffold states: created %s.md\n", tableName)
	return true, nil
}
