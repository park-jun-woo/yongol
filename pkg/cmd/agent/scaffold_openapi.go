//ff:func feature=agent type=command control=iteration dimension=1
//ff:what scaffoldOpenAPI — features.yaml ops로부터 OpenAPI path block 자동 생성

package agent

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// maxVerifyRetries is the maximum number of verify→retry cycles before
// giving up and saving the document with failing ops excluded.
const maxVerifyRetries = 3

// scaffoldOpenAPI generates specs/api/openapi.yaml from features.yaml.
// Each feature op produces one LLM call that returns a single path block.
// All blocks are merged into one OpenAPI document, verified with kin-openapi,
// and failing ops are retried up to maxVerifyRetries times. Existing file is skipped.
func scaffoldOpenAPI(specsDir string, ff *features.FeaturesFile, llmFn LLMCallFunc, cfg Config, out io.Writer) (int, error) {
	outPath := filepath.Join(specsDir, "api", "openapi.yaml")
	if _, err := os.Stat(outPath); err == nil {
		fmt.Fprintf(out, "  scaffold openapi: skipped (exists)\n")
		return 0, nil
	}

	if len(ff.Features) == 0 {
		return 0, nil
	}

	apiDir := filepath.Join(specsDir, "api")
	if err := os.MkdirAll(apiDir, 0755); err != nil {
		return 0, fmt.Errorf("create api dir: %w", err)
	}

	systemDoc, err := docs.FS.ReadFile("openapi.md")
	if err != nil {
		return 0, fmt.Errorf("read openapi.md docs: %w", err)
	}
	systemPrompt := string(systemDoc)

	// Build feature lookup by op name for retry.
	featByOp := make(map[string]features.Feature, len(ff.Features))
	for _, feat := range ff.Features {
		featByOp[feat.Op] = feat
	}

	// Collect all path blocks keyed by path string.
	pathBlocks := make(map[string]any)
	// Track which ops contributed to each path key.
	pathToOps := make(map[string][]string)
	// Track which path key each op produced (for replacement).
	opToPath := make(map[string]string)
	pathCount := 0

	for _, feat := range ff.Features {
		ddlContent := readDDLForTable(specsDir, feat.Table)
		userPrompt := buildOpenAPIUserPrompt(feat, ddlContent)

		numCtx := int(float64(len(systemPrompt)+len(userPrompt))/4*1.5) + 2048

		reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
		if err != nil {
			return 0, fmt.Errorf("scaffold openapi %s: %w", feat.Op, err)
		}

		content := stripCodeBlock(reply)
		if content == "" {
			return 0, fmt.Errorf("scaffold openapi %s: empty LLM response", feat.Op)
		}

		if mergePathBlock(pathBlocks, content) {
			pathCount++
			fmt.Fprintf(out, "  scaffold openapi: generated path for %s\n", feat.Op)
			// Record the path key(s) this op contributed
			recordPathOps(content, feat.Op, pathToOps, opToPath)
		} else {
			fmt.Fprintf(out, "  scaffold openapi: failed to parse path for %s (skipped)\n", feat.Op)
		}
	}

	// Assemble and verify loop
	projectName := "API"
	yamlDoc, offsets := assembleOpenAPI(projectName, pathBlocks, pathToOps)

	for attempt := 0; attempt < maxVerifyRetries; attempt++ {
		verifyErr := verifyOpenAPI([]byte(yamlDoc))
		if verifyErr == nil {
			// Valid — write and return
			if err := os.WriteFile(outPath, []byte(yamlDoc), 0644); err != nil {
				return 0, fmt.Errorf("scaffold openapi write: %w", err)
			}
			fmt.Fprintf(out, "  scaffold openapi: created openapi.yaml (%d paths, verified)\n", pathCount)
			return pathCount, nil
		}

		// Identify failing ops
		failedOps := extractErrorOps(verifyErr, offsets, ff.Features)
		if len(failedOps) == 0 {
			// Cannot attribute error — log and break to save as-is
			fmt.Fprintf(out, "  scaffold openapi: verify error (cannot attribute): %v\n", verifyErr)
			break
		}

		fmt.Fprintf(out, "  scaffold openapi: verify failed (attempt %d/%d), retrying ops: %v\n",
			attempt+1, maxVerifyRetries, failedOps)

		// Retry each failing op
		for _, opName := range failedOps {
			feat, ok := featByOp[opName]
			if !ok {
				continue
			}
			ddlContent := readDDLForTable(specsDir, feat.Table)
			userPrompt := buildRetryPrompt(feat, ddlContent, verifyErr.Error())

			numCtx := int(float64(len(systemPrompt)+len(userPrompt))/4*1.5) + 2048

			reply, err := llmCallWithNumCtx(cfg.Backend, cfg.Model, systemPrompt, userPrompt, numCtx)
			if err != nil {
				fmt.Fprintf(out, "  scaffold openapi: retry %s LLM error: %v\n", opName, err)
				continue
			}

			content := stripCodeBlock(reply)
			if content == "" {
				fmt.Fprintf(out, "  scaffold openapi: retry %s empty response\n", opName)
				continue
			}

			// Remove old path block for this op before merging new one
			if oldPath, ok := opToPath[opName]; ok {
				delete(pathBlocks, oldPath)
			}

			if mergePathBlock(pathBlocks, content) {
				recordPathOps(content, opName, pathToOps, opToPath)
				fmt.Fprintf(out, "  scaffold openapi: retried path for %s\n", opName)
			} else {
				fmt.Fprintf(out, "  scaffold openapi: retry parse failed for %s (skipped)\n", opName)
			}
		}

		// Reassemble with updated blocks
		yamlDoc, offsets = assembleOpenAPI(projectName, pathBlocks, pathToOps)
	}

	// Final verify after all retries
	if verifyErr := verifyOpenAPI([]byte(yamlDoc)); verifyErr != nil {
		// Remove failing ops and save the rest
		failedOps := extractErrorOps(verifyErr, offsets, ff.Features)
		if len(failedOps) > 0 {
			for _, opName := range failedOps {
				if p, ok := opToPath[opName]; ok {
					delete(pathBlocks, p)
					pathCount--
					fmt.Fprintf(out, "  scaffold openapi: excluded failing op %s\n", opName)
				}
			}
			yamlDoc, _ = assembleOpenAPI(projectName, pathBlocks, pathToOps)
		} else {
			fmt.Fprintf(out, "  scaffold openapi: verify error persists: %v\n", verifyErr)
		}
	}

	if err := os.WriteFile(outPath, []byte(yamlDoc), 0644); err != nil {
		return 0, fmt.Errorf("scaffold openapi write: %w", err)
	}

	fmt.Fprintf(out, "  scaffold openapi: created openapi.yaml (%d paths)\n", pathCount)
	return pathCount, nil
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

// buildOpenAPIUserPrompt builds the user prompt for a single path block.
func buildOpenAPIUserPrompt(feat features.Feature, ddlContent string) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Feature:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  table: %s\n", feat.Table)
	fmt.Fprintf(&b, "  public: %v\n", feat.Public)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)

	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL for table %s:\n%s\n", feat.Table, ddlContent)
	}

	b.WriteString("\nGenerate ONLY a single OpenAPI path block for this feature.")
	b.WriteString("\nThe output must be a valid YAML fragment starting with the path key (e.g. /resources/{id}:).")
	b.WriteString("\nInclude operationId, request/response schemas.")
	b.WriteString("\nUse 2-space indentation. No surrounding 'paths:' key. No markdown fences.")

	return b.String()
}

// mergePathBlock parses a YAML path block and merges it into pathBlocks.
// If parsing fails, the block is skipped with a warning on stderr and returns false.
func mergePathBlock(pathBlocks map[string]any, content string) bool {
	// Try to parse as YAML map
	var parsed map[string]any
	if err := yaml.Unmarshal([]byte(content), &parsed); err == nil && len(parsed) > 0 {
		for k, v := range parsed {
			pathBlocks[k] = v
		}
		return true
	}

	// Fallback: try wrapping with paths: prefix then extracting
	wrapped := "paths:\n" + indentText(content, "  ")
	var wrappedParsed struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal([]byte(wrapped), &wrappedParsed); err == nil && len(wrappedParsed.Paths) > 0 {
		for k, v := range wrappedParsed.Paths {
			pathBlocks[k] = v
		}
		return true
	}

	// Skip unparseable block — validate loop will catch the missing path
	fmt.Fprintf(os.Stderr, "  scaffold openapi: warning: skipped unparseable path block\n")
	return false
}

// assembleOpenAPI builds the final openapi.yaml document and returns both
// the YAML string and per-path line offset information for error attribution.
// pathToOps maps each path key (e.g. "/workflows/{id}") to its originating op names.
func assembleOpenAPI(projectName string, pathBlocks map[string]any, pathToOps map[string][]string) (string, []pathOffset) {
	// Build header
	var b strings.Builder
	b.WriteString("openapi: \"3.1.0\"\n")
	b.WriteString("info:\n")
	fmt.Fprintf(&b, "  title: %s\n", projectName)
	b.WriteString("  version: \"1.0.0\"\n")
	b.WriteString("security:\n")
	b.WriteString("  - bearerAuth: []\n")
	b.WriteString("paths:\n")

	// Track current line (1-based). Header above is 7 lines.
	currentLine := 8 // next line after "paths:\n"

	// Marshal each path block individually for offset tracking.
	// Sort keys for deterministic output.
	pathKeys := make([]string, 0, len(pathBlocks))
	for k := range pathBlocks {
		pathKeys = append(pathKeys, k)
	}
	sort.Strings(pathKeys)

	var offsets []pathOffset

	for _, pathKey := range pathKeys {
		block := map[string]any{pathKey: pathBlocks[pathKey]}
		blockYAML, err := yaml.Marshal(block)
		if err != nil {
			continue
		}
		indented := indentText(string(blockYAML), "  ")
		lineCount := strings.Count(indented, "\n")

		// Record offset for each op associated with this path
		ops := pathToOps[pathKey]
		if len(ops) == 0 {
			ops = []string{pathKey} // fallback: use path as op identifier
		}
		for _, op := range ops {
			offsets = append(offsets, pathOffset{
				Op:        op,
				Path:      pathKey,
				StartLine: currentLine,
				EndLine:   currentLine + lineCount - 1,
			})
		}

		b.WriteString(indented)
		currentLine += lineCount
	}

	// Footer: components
	b.WriteString("components:\n")
	b.WriteString("  securitySchemes:\n")
	b.WriteString("    bearerAuth:\n")
	b.WriteString("      type: http\n")
	b.WriteString("      scheme: bearer\n")
	b.WriteString("      bearerFormat: JWT\n")
	b.WriteString("  schemas:\n")
	b.WriteString("    Error:\n")
	b.WriteString("      type: object\n")
	b.WriteString("      required: [error, code]\n")
	b.WriteString("      properties:\n")
	b.WriteString("        error:\n")
	b.WriteString("          type: string\n")
	b.WriteString("        code:\n")
	b.WriteString("          type: string\n")

	return b.String(), offsets
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
	// Parse the full document
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal([]byte(openapiContent), &doc); err != nil || doc.Paths == nil {
		return ""
	}

	// Search for the path containing this operationId
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
				// Found it — marshal this path block back
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
