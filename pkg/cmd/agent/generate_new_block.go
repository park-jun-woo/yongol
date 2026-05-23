//ff:func feature=agent type=command control=selection
//ff:what generateNewBlock — SSaC + features로 새 블록 생성 후 파일에 삽입

package agent

import (
	"fmt"
	"io"
)

// generateNewBlock generates a new block for an operationId using SSaC content + features,
// then inserts it into the file. Returns true if block was generated and inserted.
func generateNewBlock(w io.Writer, cfg Config, l layer, relFile, absPath string, fileContent *string, opID, desc, featurePath string) bool {
	specsDir := resolveSpecsRoot(absPath, l)

	ssacContent, ssacFound := findSSaCFile(specsDir, opID)
	if !ssacFound {
		fmt.Fprintf(w, "  skipped generate: %s/%s (SSaC file not found)\n", relFile, opID)
		return false
	}

	systemPrompt := "You generate yongol SSOT content. Output ONLY the requested block. No explanations. No markdown fences.\n\nExample for " + layerName(l) + ":\n" + layerExample(l)
	userPrompt := buildGeneratePrompt(l, opID, desc, featurePath, ssacContent)

	reply, err := llmCall(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		fmt.Fprintf(w, "  skipped generate: %s/%s (LLM error: %v)\n", relFile, opID, err)
		return false
	}

	newBlock := stripMarkdownFences(reply)
	if newBlock == "" {
		fmt.Fprintf(w, "  skipped generate: %s/%s (empty LLM response)\n", relFile, opID)
		return false
	}

	var inserted string
	switch l {
	case layerOpenAPI:
		inserted, err = insertOpenAPIBlock(*fileContent, newBlock)
	case layerRego:
		inserted, err = insertRegoBlock(*fileContent, newBlock)
	case layerHurl:
		inserted, err = insertHurlBlock(*fileContent, newBlock)
	default:
		return false
	}

	if err != nil {
		fmt.Fprintf(w, "  skipped generate: %s/%s (insert error: %v)\n", relFile, opID, err)
		return false
	}

	*fileContent = inserted
	fmt.Fprintf(w, "  generated: %s/%s\n", relFile, opID)
	return true
}
