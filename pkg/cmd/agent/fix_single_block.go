//ff:func feature=agent type=command control=selection
//ff:what fixSingleBlock — operationId 블록 추출 → LLM 수정 → 머지

package agent

import (
	"fmt"
	"io"
	"path/filepath"
)

// fixSingleBlock extracts a block for an operationId, sends it to LLM, and merges
// the result back into fileContent. Returns true if the block was fixed.
func fixSingleBlock(w io.Writer, cfg Config, l layer, relFile, absPath string, fileContent *string, opID, desc, path string, diagMsgs []string) bool {
	var block string
	var startLine, endLine int
	var extractErr error

	switch l {
	case layerOpenAPI:
		block, startLine, endLine, extractErr = extractOpenAPIBlock(*fileContent, opID)
	case layerRego:
		block, startLine, endLine, extractErr = extractRegoBlock(*fileContent, opID)
	case layerHurl:
		block, startLine, endLine, extractErr = extractHurlBlock(*fileContent, opID)
	default:
		return false
	}

	if extractErr != nil {
		return generateNewBlock(w, cfg, l, relFile, absPath, fileContent, opID, desc, path)
	}

	systemPrompt := buildSystemPrompt(l, diagMsgs)
	userPrompt := buildBlockUserPrompt(desc, path, filepath.Base(relFile), opID, block, diagMsgs)

	reply, err := llmCall(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		fmt.Fprintf(w, "  skipped block: %s/%s (LLM error: %v)\n", relFile, opID, err)
		return false
	}

	fixedBlock := stripMarkdownFences(reply)
	if fixedBlock == "" {
		fmt.Fprintf(w, "  skipped block: %s/%s (empty LLM response)\n", relFile, opID)
		return false
	}

	var merged string
	switch l {
	case layerOpenAPI:
		merged, err = mergeOpenAPIBlock(*fileContent, startLine, endLine, fixedBlock)
	case layerRego:
		merged, err = mergeRegoBlock(*fileContent, startLine, endLine, fixedBlock)
	case layerHurl:
		merged, err = mergeHurlBlock(*fileContent, startLine, endLine, fixedBlock)
	}

	if err != nil {
		fmt.Fprintf(w, "  skipped block: %s/%s (merge error: %v)\n", relFile, opID, err)
		return false
	}

	*fileContent = merged
	return true
}
