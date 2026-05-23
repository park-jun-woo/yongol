//ff:func feature=agent type=helper control=sequence
//ff:what scaffoldSSaCFeature — 단일 feature의 SSaC 서비스 파일 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func scaffoldSSaCFeature(specsDir string, feat features.Feature, openapiContent, systemPrompt string, cfg Config, out io.Writer) (bool, error) {
	domain := domainFromPath(feat.Path)
	serviceDir := filepath.Join(specsDir, "service", domain)
	outPath := filepath.Join(serviceDir, feat.Op+".ssac")

	if _, err := os.Stat(outPath); err == nil {
		fmt.Fprintf(out, "  scaffold ssac: skipped %s/%s.ssac (exists)\n", domain, feat.Op)
		return false, nil
	}

	if err := os.MkdirAll(serviceDir, 0755); err != nil {
		return false, fmt.Errorf("create service/%s dir: %w", domain, err)
	}

	ddlContent := readDDLForTable(specsDir, feat.Table)
	queryNames := readSQLcQueryNames(specsDir, feat.Table)
	pathBlock := extractPathBlockForOp(openapiContent, feat.Op)
	userPrompt := buildSSaCUserPrompt(feat, ddlContent, queryNames, pathBlock)
	numCtx := len(systemPrompt) + len(userPrompt) + 2048

	reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
	if err != nil {
		return false, fmt.Errorf("scaffold ssac %s: %w", feat.Op, err)
	}

	content := stripCodeBlock(reply)
	if content == "" {
		return false, fmt.Errorf("scaffold ssac %s: empty LLM response", feat.Op)
	}

	if err := os.WriteFile(outPath, []byte(content+"\n"), 0644); err != nil {
		return false, fmt.Errorf("scaffold ssac %s: write: %w", feat.Op, err)
	}

	fmt.Fprintf(out, "  scaffold ssac: created %s/%s.ssac\n", domain, feat.Op)
	return true, nil
}
