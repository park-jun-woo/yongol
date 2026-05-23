//ff:func feature=agent type=command control=sequence
//ff:what fixFile — 단일 파일의 validate 에러를 LLM으로 수정 (레이어별 컨텍스트 주입)

package agent

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// fixFile reads the target file, builds layer-specific context, calls LLM to
// fix validate errors, and writes the result back. For OpenAPI and Rego the
// fix is done per-block (operationId granularity) using existing split logic.
func fixFile(specsDir string, ff *features.FeaturesFile, filename string, diags []diagnostic.Diagnostic, llmFn LLMCallFunc, cfg Config) error {
	l := classifyFile(filename)
	absPath := filepath.Join(specsDir, filename)

	content, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("read %s: %w", filename, err)
	}

	// For split-based layers (OpenAPI, Rego, Hurl) delegate to block-level fix.
	if l == layerOpenAPI || l == layerRego || l == layerHurl {
		return fixSplitFile(specsDir, ff, filename, absPath, string(content), diags, llmFn, cfg, l)
	}

	// Whole-file fix for other layers.
	systemPrompt := buildFixSystemPrompt(l, diags)
	userPrompt := buildFixUserPrompt(specsDir, ff, filename, string(content), diags, l)

	reply, err := llmFn(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		return fmt.Errorf("LLM %s: %w", filename, err)
	}

	fixed := stripMarkdownFences(reply)
	if fixed == "" {
		return fmt.Errorf("empty LLM response for %s", filename)
	}

	return os.WriteFile(absPath, []byte(fixed), 0644)
}
