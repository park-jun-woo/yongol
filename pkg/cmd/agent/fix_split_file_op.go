//ff:func feature=agent type=helper control=selection
//ff:what fixSplitFileOp — 단일 operationId의 블록 추출→수정→머지

package agent

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func fixSplitFileOp(specsDir string, ff *features.FeaturesFile, fileContent *string, opID string, diags []diagnostic.Diagnostic, msgs []string, lookup map[string]features.Feature, llmFn LLMCallFunc, cfg Config, l layer) error {
	opMsgs := filterMessagesByOp(msgs, opID)
	if len(opMsgs) == 0 {
		opMsgs = msgs
	}

	opDiags := filterDiagsByOp(diags, opID)

	desc := ""
	if feat, ok := lookup[opID]; ok {
		desc = feat.Desc
	}

	systemPrompt := buildFixSystemPrompt(l, opDiags)

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
	}

	if extractErr != nil {
		return extractErr
	}

	userPrompt := buildBlockFixUserPrompt(specsDir, ff, opID, desc, block, opDiags, l)

	reply, err := llmFn(cfg.Backend, cfg.Model, systemPrompt, userPrompt)
	if err != nil {
		return err
	}

	fixedBlock := stripMarkdownFences(reply)
	if fixedBlock == "" {
		return fmt.Errorf("empty response")
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
		return err
	}
	*fileContent = merged
	return nil
}
