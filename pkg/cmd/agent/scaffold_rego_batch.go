//ff:func feature=agent type=helper control=sequence
//ff:what scaffoldRegoBatch — feature 배치의 Rego 룰 블록 생성

package agent

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func scaffoldRegoBatch(batch []features.Feature, batchIdx int, systemPrompt string, cfg Config) (string, error) {
	userPrompt := buildRegoUserPrompt(batch)
	numCtx := len(systemPrompt) + len(userPrompt) + 2048

	reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
	if err != nil {
		return "", fmt.Errorf("scaffold rego batch %d: %w", batchIdx, err)
	}

	content := stripCodeBlock(reply)
	if content == "" {
		return "", fmt.Errorf("scaffold rego batch %d: empty LLM response", batchIdx)
	}

	return content, nil
}
