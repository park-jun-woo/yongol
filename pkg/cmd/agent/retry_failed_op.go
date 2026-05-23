//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what retryFailedOp — 실패한 op의 path 블록 재생성

package agent

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

func retryFailedOp(opName string, featByOp map[string]features.Feature, relativeLines map[string]int, verifyErr error, pathBlocks map[string]any, pathToOps map[string][]string, opToPath map[string]string, cfg Config, out io.Writer) {
	feat, ok := featByOp[opName]
	if !ok {
		return
	}
	ddlContent := readDDLForTable(cfg.SpecsDir, feat.Table)
	rl := -1
	if relativeLines != nil {
		if v, ok := relativeLines[opName]; ok {
			rl = v
		}
	}

	retryPrompt := buildRetryPrompt(feat, ddlContent, verifyErr.Error(), rl)
	numCtx := len(splitSystemPrompt) + len(retryPrompt) + 2048

	reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, retryPrompt, numCtx)
	if err != nil {
		fmt.Fprintf(out, "  scaffold openapi: retry %s LLM error: %v\n", opName, err)
		return
	}

	content := stripCodeBlock(reply)
	if content == "" {
		fmt.Fprintf(out, "  scaffold openapi: retry %s empty response\n", opName)
		return
	}

	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(content), &parsed); err != nil || len(parsed) == 0 {
		fmt.Fprintf(out, "  scaffold openapi: retry parse failed for %s (skipped)\n", opName)
		return
	}

	if oldPath, ok := opToPath[opName]; ok {
		delete(pathBlocks, oldPath)
	}

	for k, v := range parsed {
		pathBlocks[k] = v
		pathToOps[k] = appendUnique(pathToOps[k], opName)
		opToPath[opName] = k
	}
	fmt.Fprintf(out, "  scaffold openapi: retried path for %s\n", opName)
}
