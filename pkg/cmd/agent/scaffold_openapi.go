//ff:func feature=agent type=command control=iteration dimension=4
//ff:what scaffoldOpenAPI — features.yaml ops로부터 OpenAPI path block 분할 생성 (parameters + requestBody + schema200 → 조립)

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// maxVerifyRetries is the maximum number of verify→retry cycles before
// giving up and saving the document with failing ops excluded.
const maxVerifyRetries = 3

// splitSystemPrompt is the shared system prompt for all split LLM calls.
const splitSystemPrompt = "Output ONLY valid YAML. No explanations. No markdown fences."

// scaffoldOpenAPI generates specs/api/openapi.yaml from features.yaml.
// Each feature op produces up to 3 LLM calls (parameters, requestBody, schema200)
// plus 2 mechanical steps (error responses, path block assembly).
// All blocks are merged into one OpenAPI document, verified with kin-openapi,
// and failing ops are retried up to maxVerifyRetries times.
// When the file already exists, only missing operationIds are generated (incremental mode).
func scaffoldOpenAPI(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) (int, error) {
	outPath := filepath.Join(specsDir, "api", "openapi.yaml")

	if len(ff.Features) == 0 {
		return 0, nil
	}

	apiDir := filepath.Join(specsDir, "api")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		return 0, fmt.Errorf("create api dir: %w", err)
	}

	// Check for existing operationIDs (incremental mode).
	existing := existingOperationIDs(specsDir)
	incremental := len(existing) > 0

	// In incremental mode, determine which ops are missing.
	var missingFeats []features.Feature
	if incremental {
		for _, feat := range ff.Features {
			if !existing[feat.Op] {
				missingFeats = append(missingFeats, feat)
			}
		}
		if len(missingFeats) == 0 {
			fmt.Fprintf(out, "  scaffold openapi: all paths exist (%d ops)\n", len(existing))
			return 0, nil
		}
		fmt.Fprintf(out, "  scaffold openapi: incremental — %d ops missing, generating...\n", len(missingFeats))
	} else {
		missingFeats = ff.Features
	}

	// Build feature lookup by op name for retry.
	featByOp := make(map[string]features.Feature, len(ff.Features))
	for _, feat := range ff.Features {
		featByOp[feat.Op] = feat
	}

	// Seed path blocks from existing file when in incremental mode.
	var pathBlocks map[string]any
	var pathToOps map[string][]string
	if incremental {
		pathBlocks, pathToOps = existingPathBlocks(specsDir)
		if pathBlocks == nil {
			pathBlocks = make(map[string]any)
		}
		if pathToOps == nil {
			pathToOps = make(map[string][]string)
		}
	} else {
		pathBlocks = make(map[string]any)
		pathToOps = make(map[string][]string)
	}
	// Track which path key each op produced.
	opToPath := make(map[string]string)
	// Pre-populate opToPath from existing pathToOps.
	for pathKey, ops := range pathToOps {
		for _, op := range ops {
			opToPath[op] = pathKey
		}
	}
	pathCount := len(pathBlocks)

	for _, feat := range missingFeats {
		ddlContent := readDDLForTable(specsDir, feat.Table)

		block, err := generateSplitPathBlock(feat, ddlContent, cfg)
		if err != nil {
			fmt.Fprintf(out, "  scaffold openapi: failed to generate path for %s: %v\n", feat.Op, err)
			continue
		}

		// Merge new method into existing path block (preserves sibling methods).
		mergeMethodBlock(pathBlocks, feat.Path, block)
		pathToOps[feat.Path] = appendUnique(pathToOps[feat.Path], feat.Op)
		opToPath[feat.Op] = feat.Path
		pathCount++
		fmt.Fprintf(out, "  scaffold openapi: generated path for %s\n", feat.Op)
	}

	// Build set of newly generated ops so retries/exclusions never touch existing paths.
	newOps := make(map[string]bool, len(missingFeats))
	for _, feat := range missingFeats {
		newOps[feat.Op] = true
	}

	// Assemble and verify loop
	projectName := "API"
	yamlDoc, offsets := assembleFullOpenAPI(projectName, pathBlocks, pathToOps)

	for attempt := 0; attempt < maxVerifyRetries; attempt++ {
		result := scaffoldOpenAPIVerifyRetry(&yamlDoc, &offsets, pathBlocks, pathToOps, opToPath, featByOp, newOps, incremental, attempt, cfg, out, ff)
		if result.verified {
			if err := os.WriteFile(outPath, []byte(yamlDoc), 0644); err != nil {
				return 0, fmt.Errorf("scaffold openapi write: %w", err)
			}
			label := "created"
			if incremental {
				label = "updated"
			}
			fmt.Fprintf(out, "  scaffold openapi: %s openapi.yaml (%d paths)\n", label, pathCount)
			return pathCount, nil
		}
		if result.stopped {
			break
		}
	}

	// Final verify after all retries (only exclude newly generated ops, never existing)
	if verifyErr := verifyOpenAPI([]byte(yamlDoc)); verifyErr != nil {
		allFinalFailed, _ := extractErrorOps(verifyErr, offsets, ff.Features, yamlDoc)
		var finalFailed []string
		for _, op := range allFinalFailed {
			if !incremental || newOps[op] {
				finalFailed = append(finalFailed, op)
			}
		}
		if len(finalFailed) > 0 {
			for _, opName := range finalFailed {
				if p, ok := opToPath[opName]; ok {
					delete(pathBlocks, p)
					pathCount--
					fmt.Fprintf(out, "  scaffold openapi: excluded failing op %s\n", opName)
				}
			}
			yamlDoc, _ = assembleFullOpenAPI(projectName, pathBlocks, pathToOps)
		} else {
			fmt.Fprintf(out, "  scaffold openapi: verify error persists: %v\n", verifyErr)
		}
	}

	if err := os.WriteFile(outPath, []byte(yamlDoc), 0644); err != nil {
		return 0, fmt.Errorf("scaffold openapi write: %w", err)
	}

	if incremental {
		fmt.Fprintf(out, "  scaffold openapi: updated openapi.yaml (%d paths)\n", pathCount)
	} else {
		fmt.Fprintf(out, "  scaffold openapi: created openapi.yaml (%d paths)\n", pathCount)
	}
	return pathCount, nil
}
