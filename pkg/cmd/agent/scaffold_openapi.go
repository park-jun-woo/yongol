//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldOpenAPI — features.yaml ops로부터 OpenAPI path block 분할 생성 (parameters + requestBody + schema200 → 조립)

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// maxVerifyRetries is the maximum number of verify→retry cycles before
// giving up and saving the document with failing ops excluded.
const maxVerifyRetries = 3

// splitSystemPrompt is the shared system prompt for all split LLM calls.
const splitSystemPrompt = "Output ONLY valid YAML. No explanations. No markdown fences."

// existingOperationIDs reads specs/api/openapi.yaml and returns the set of
// operationIds already defined. Returns an empty map if the file does not
// exist or cannot be parsed.
func existingOperationIDs(specsDir string) map[string]bool {
	data, err := os.ReadFile(filepath.Join(specsDir, "api", "openapi.yaml"))
	if err != nil {
		return nil
	}
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil || doc.Paths == nil {
		return nil
	}
	ids := make(map[string]bool)
	for _, methods := range doc.Paths {
		methodMap, ok := methods.(map[string]any)
		if !ok {
			continue
		}
		for _, detail := range methodMap {
			detailMap, ok := detail.(map[string]any)
			if !ok {
				continue
			}
			if opID, ok := detailMap["operationId"]; ok {
				ids[fmt.Sprintf("%v", opID)] = true
			}
		}
	}
	return ids
}

// existingPathBlocks reads specs/api/openapi.yaml and returns the paths map
// (path key → method map) and a pathToOps mapping. Returns nil maps if the
// file does not exist or cannot be parsed.
func existingPathBlocks(specsDir string) (map[string]any, map[string][]string) {
	data, err := os.ReadFile(filepath.Join(specsDir, "api", "openapi.yaml"))
	if err != nil {
		return nil, nil
	}
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil || doc.Paths == nil {
		return nil, nil
	}
	pathToOps := make(map[string][]string)
	for pathKey, methods := range doc.Paths {
		methodMap, ok := methods.(map[string]any)
		if !ok {
			continue
		}
		for _, detail := range methodMap {
			detailMap, ok := detail.(map[string]any)
			if !ok {
				continue
			}
			if opID, ok := detailMap["operationId"]; ok {
				pathToOps[pathKey] = appendUnique(pathToOps[pathKey], fmt.Sprintf("%v", opID))
			}
		}
	}
	return doc.Paths, pathToOps
}

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
		verifyErr := verifyOpenAPI([]byte(yamlDoc))
		if verifyErr == nil {
			if err := os.WriteFile(outPath, []byte(yamlDoc), 0644); err != nil {
				return 0, fmt.Errorf("scaffold openapi write: %w", err)
			}
			if incremental {
				fmt.Fprintf(out, "  scaffold openapi: updated openapi.yaml (%d paths)\n", pathCount)
			} else {
				fmt.Fprintf(out, "  scaffold openapi: created openapi.yaml (%d paths, verified)\n", pathCount)
			}
			return pathCount, nil
		}

		// Identify failing ops (in incremental mode, only retry newly generated ops)
		allFailed, relativeLines := extractErrorOps(verifyErr, offsets, ff.Features, yamlDoc)
		var failedOps []string
		for _, op := range allFailed {
			if !incremental || newOps[op] {
				failedOps = append(failedOps, op)
			}
		}
		if len(failedOps) == 0 {
			fmt.Fprintf(out, "  scaffold openapi: verify error (cannot attribute): %v\n", verifyErr)
			break
		}

		fmt.Fprintf(out, "  scaffold openapi: verify failed (attempt %d/%d), retrying ops: %v\n",
			attempt+1, maxVerifyRetries, failedOps)

		// Retry each failing op with split generation
		for _, opName := range failedOps {
			feat, ok := featByOp[opName]
			if !ok {
				continue
			}
			ddlContent := readDDLForTable(specsDir, feat.Table)
			rl := -1
			if relativeLines != nil {
				if v, ok := relativeLines[opName]; ok {
					rl = v
				}
			}

			// On retry, use the full-block retry prompt (from openapi_verify.go)
			// since the error context is needed for correction.
			retryPrompt := buildRetryPrompt(feat, ddlContent, verifyErr.Error(), rl)
			numCtx := len(splitSystemPrompt) + len(retryPrompt) + 2048

			reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, retryPrompt, numCtx)
			if err != nil {
				fmt.Fprintf(out, "  scaffold openapi: retry %s LLM error: %v\n", opName, err)
				continue
			}

			content := stripCodeBlock(reply)
			if content == "" {
				fmt.Fprintf(out, "  scaffold openapi: retry %s empty response\n", opName)
				continue
			}

			// Parse the retry response as a path block
			var parsed map[string]any
			if err := yaml.Unmarshal([]byte(content), &parsed); err != nil || len(parsed) == 0 {
				fmt.Fprintf(out, "  scaffold openapi: retry parse failed for %s (skipped)\n", opName)
				continue
			}

			// Remove old path block for this op before merging new one
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

		// Reassemble with updated blocks
		yamlDoc, offsets = assembleFullOpenAPI(projectName, pathBlocks, pathToOps)
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

// generateSplitPathBlock generates a single op's path block through split LLM calls:
// Step 1: parameters, Step 2: requestBody (POST/PUT only), Step 3: response 200 schema.
// Steps 4-5 are mechanical (error responses + assembly).
func generateSplitPathBlock(feat features.Feature, ddlContent string, cfg Config) (map[string]any, error) {
	// Step 1: parameters
	params, err := callStepWithRetry(cfg, buildParamsPrompt(feat))
	if err != nil {
		return nil, fmt.Errorf("parameters: %w", err)
	}
	var paramsParsed []any
	if params != "none" && strings.TrimSpace(params) != "" {
		if err := yaml.Unmarshal([]byte(params), &paramsParsed); err != nil {
			// LLM may wrap the array in a "parameters:" key — unwrap it.
			var wrapped map[string]any
			if err2 := yaml.Unmarshal([]byte(params), &wrapped); err2 == nil {
				if inner, ok := wrapped["parameters"]; ok {
					if arr, ok := inner.([]any); ok {
						paramsParsed = arr
					} else {
						return nil, fmt.Errorf("parse parameters YAML: %w", err)
					}
				} else {
					return nil, fmt.Errorf("parse parameters YAML: %w", err)
				}
			} else {
				return nil, fmt.Errorf("parse parameters YAML: %w", err)
			}
		}
	}

	// Step 2: requestBody (POST/PUT only)
	method := httpMethodFromOp(feat.Op)
	var reqBodyParsed map[string]any
	if needsRequestBody(method) {
		reqBodyRaw, err := callStepWithRetry(cfg, buildRequestBodyPrompt(feat, ddlContent))
		if err != nil {
			return nil, fmt.Errorf("requestBody: %w", err)
		}
		if reqBodyRaw != "none" && strings.TrimSpace(reqBodyRaw) != "" {
			if err := yaml.Unmarshal([]byte(reqBodyRaw), &reqBodyParsed); err != nil {
				return nil, fmt.Errorf("parse requestBody YAML: %w", err)
			}
			// LLM may wrap the object in a "requestBody:" key — unwrap it.
			if inner, ok := reqBodyParsed["requestBody"]; ok && len(reqBodyParsed) == 1 {
				if m, ok := inner.(map[string]any); ok {
					reqBodyParsed = m
				}
			}
		}
	}

	// Step 3: response 200 schema
	schema200Raw, err := callStepWithRetry(cfg, buildSchema200Prompt(feat, ddlContent))
	if err != nil {
		return nil, fmt.Errorf("schema200: %w", err)
	}
	var schema200Parsed map[string]any
	if err := yaml.Unmarshal([]byte(schema200Raw), &schema200Parsed); err != nil {
		return nil, fmt.Errorf("parse schema200 YAML: %w", err)
	}
	// LLM may wrap the object in a "schema:" key — unwrap it.
	if inner, ok := schema200Parsed["schema"]; ok && len(schema200Parsed) == 1 {
		if m, ok := inner.(map[string]any); ok {
			schema200Parsed = m
		}
	}

	// Step 4: error responses (mechanical)
	errorResps := buildErrorResponses(feat)

	// Step 5: assemble path block (mechanical)
	block := assemblePathBlock(feat, paramsParsed, reqBodyParsed, schema200Parsed, errorResps)
	return block, nil
}

// callStepWithRetry calls the LLM for a single step and retries once on parse failure.
// Returns the cleaned response content.
func callStepWithRetry(cfg Config, userPrompt string) (string, error) {
	numCtx := len(splitSystemPrompt) + len(userPrompt) + 2048

	reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, userPrompt, numCtx)
	if err != nil {
		// Retry once
		reply, err = llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, userPrompt, numCtx)
		if err != nil {
			return "", err
		}
	}

	content := stripCodeBlock(reply)
	if content == "" {
		// Retry once on empty response
		reply, err = llmCallWithNumCtx(cfg.Backend, cfg.Model, splitSystemPrompt, userPrompt, numCtx)
		if err != nil {
			return "", err
		}
		content = stripCodeBlock(reply)
		if content == "" {
			return "", fmt.Errorf("empty LLM response after retry")
		}
	}

	return content, nil
}

// buildParamsPrompt builds the user prompt for Step 1: parameters.
func buildParamsPrompt(feat features.Feature) string {
	var b strings.Builder
	b.WriteString("OpenAPI parameters for this endpoint:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)
	b.WriteString("\nRules:\n")
	b.WriteString("- Path parameters need matching parameters[] entry\n")
	b.WriteString("- All integers: type: integer, format: int64\n")
	b.WriteString("\nOutput ONLY the parameters array in YAML. Output \"none\" if no parameters.")
	return b.String()
}

// buildRequestBodyPrompt builds the user prompt for Step 2: requestBody.
func buildRequestBodyPrompt(feat features.Feature, ddlContent string) string {
	var b strings.Builder
	b.WriteString("OpenAPI requestBody for this endpoint:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)
	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL:\n%s\n", ddlContent)
	}
	b.WriteString("\nOutput \"none\" if no requestBody needed, or the requestBody YAML block.")
	return b.String()
}

// buildSchema200Prompt builds the user prompt for Step 3: response 200 schema.
func buildSchema200Prompt(feat features.Feature, ddlContent string) string {
	var b strings.Builder
	b.WriteString("OpenAPI 200 response schema for this endpoint:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)
	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL:\n%s\n", ddlContent)
	}
	b.WriteString("\nAll integers: format: int64. Timestamps: format: date-time.\n")
	b.WriteString("\nOutput ONLY the schema (type, required, properties) in YAML. No wrapping 'schema:' key.")
	return b.String()
}

// mergeMethodBlock merges a newly generated method block into pathBlocks for
// the given path key. If the path already exists, the new method entries are
// added alongside existing ones (so GET + POST on the same path coexist).
// If the path does not exist yet, the block is stored directly.
func mergeMethodBlock(pathBlocks map[string]any, pathKey string, block map[string]any) {
	existing, ok := pathBlocks[pathKey]
	if !ok {
		pathBlocks[pathKey] = block
		return
	}
	existingMap, ok := existing.(map[string]any)
	if !ok {
		pathBlocks[pathKey] = block
		return
	}
	for method, detail := range block {
		existingMap[method] = detail
	}
	pathBlocks[pathKey] = existingMap
}

// recordPathOps parses a YAML path block to extract path keys and records the
// mapping from path key to op name in both pathToOps and opToPath.
func recordPathOps(content, opName string, pathToOps map[string][]string, opToPath map[string]string) {
	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(content), &parsed); err == nil {
		for k := range parsed {
			pathToOps[k] = appendUnique(pathToOps[k], opName)
			opToPath[opName] = k
		}
		return
	}
	// Fallback: try with paths: wrapper
	wrapped := "paths:\n" + indentText(content, "  ")
	var wrappedParsed struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal([]byte(wrapped), &wrappedParsed); err == nil {
		for k := range wrappedParsed.Paths {
			pathToOps[k] = appendUnique(pathToOps[k], opName)
			opToPath[opName] = k
		}
	}
}

// appendUnique appends s to slice only if not already present.
func appendUnique(slice []string, s string) []string {
	for _, v := range slice {
		if v == s {
			return slice
		}
	}
	return append(slice, s)
}

// readDDLForTable reads the DDL file for a given table name. Returns empty if not found.
func readDDLForTable(specsDir, tableName string) string {
	if tableName == "" {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(specsDir, "db", tableName+".sql"))
	if err != nil {
		return ""
	}
	return string(data)
}

// indentText prepends a given prefix to each non-empty line.
func indentText(text, prefix string) string {
	lines := strings.Split(strings.TrimRight(text, "\n"), "\n")
	var b strings.Builder
	for _, line := range lines {
		if line == "" {
			b.WriteByte('\n')
		} else {
			b.WriteString(prefix)
			b.WriteString(line)
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// extractPathBlockForOp extracts the OpenAPI path block content for a given operationId
// from the full openapi.yaml content string. Returns the path block or empty string.
func extractPathBlockForOp(openapiContent, operationId string) string {
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal([]byte(openapiContent), &doc); err != nil || doc.Paths == nil {
		return ""
	}

	for pathKey, methods := range doc.Paths {
		methodMap, ok := methods.(map[string]any)
		if !ok {
			continue
		}
		for method, detail := range methodMap {
			detailMap, ok := detail.(map[string]any)
			if !ok {
				continue
			}
			if opID, ok := detailMap["operationId"]; ok && fmt.Sprintf("%v", opID) == operationId {
				block := map[string]any{pathKey: map[string]any{method: detail}}
				out, err := yaml.Marshal(block)
				if err != nil {
					return ""
				}
				return string(out)
			}
		}
	}
	return ""
}
