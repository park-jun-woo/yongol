//ff:func feature=agent type=adapter control=sequence
//ff:what openapi_verify — kin-openapi loader 검증 + 에러→op 매핑

package agent

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// pathOffset records the line range of a single op's path block in the
// assembled OpenAPI YAML. Used by extractErrorOps to map YAML-level errors
// back to the originating feature op.
type pathOffset struct {
	Op        string
	Path      string // e.g. /workflows/{id}
	StartLine int    // 1-based inclusive
	EndLine   int    // 1-based inclusive
}

// verifyOpenAPI loads yamlData through kin-openapi and returns the first
// validation error, or nil when the document is structurally valid.
func verifyOpenAPI(yamlData []byte) error {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(yamlData)
	if err != nil {
		return err
	}
	if err := doc.Validate(loader.Context); err != nil {
		return err
	}
	return nil
}

// extractErrorOps inspects the error message from verifyOpenAPI and tries to
// identify which ops produced the error.
//
// Strategy:
//  1. YAML line errors ("yaml: line N:") — map N to an offset range.
//  2. $ref errors containing "#/components/schemas/..." — exact match to op.
//  3. Path key errors containing a path like "/resources/{id}" — match to offset.
//  4. Quoted keyword grep — search keywords in YAML content and map to ops.
//
// Returns matched ops and a map of op → relative line within the op's block
// (populated only by strategy 4; nil when no grep matches).
func extractErrorOps(err error, offsets []pathOffset, feats []features.Feature, yamlContent string) ([]string, map[string]int) {
	if err == nil {
		return nil, nil
	}
	msg := err.Error()
	seen := make(map[string]bool)

	// --- strategy 1: yaml line number ---
	if ops := matchByLine(msg, offsets); len(ops) > 0 {
		for _, op := range ops {
			seen[op] = true
		}
	}

	// --- strategy 2: $ref / schema name ---
	if ops := matchBySchema(msg, offsets, feats); len(ops) > 0 {
		for _, op := range ops {
			seen[op] = true
		}
	}

	// --- strategy 3: path key in error ---
	if ops := matchByPath(msg, offsets); len(ops) > 0 {
		for _, op := range ops {
			seen[op] = true
		}
	}

	// --- strategy 4: keyword grep in YAML content ---
	var relativeLines map[string]int
	if grepOps, rl := matchByGrep(msg, yamlContent, offsets); len(grepOps) > 0 {
		relativeLines = rl
		for _, op := range grepOps {
			seen[op] = true
		}
	}

	out := make([]string, 0, len(seen))
	for op := range seen {
		out = append(out, op)
	}
	return out, relativeLines
}

// reYAMLLine matches "yaml: line 15:" or "line 15:" patterns.
var reYAMLLine = regexp.MustCompile(`(?:yaml: )?line (\d+)`)

// matchByLine extracts line numbers from a YAML error and maps them to ops.
func matchByLine(msg string, offsets []pathOffset) []string {
	matches := reYAMLLine.FindAllStringSubmatch(msg, -1)
	if len(matches) == 0 {
		return nil
	}
	var ops []string
	for _, m := range matches {
		lineNo, err := strconv.Atoi(m[1])
		if err != nil {
			continue
		}
		for _, off := range offsets {
			if lineNo >= off.StartLine && lineNo <= off.EndLine {
				ops = append(ops, off.Op)
				break
			}
		}
	}
	return ops
}

// reSchemaRef matches #/components/schemas/SomeName
var reSchemaRef = regexp.MustCompile(`#/components/schemas/(\w+)`)

// matchBySchema finds schema names in the error and matches them to feature
// ops by exact (case-insensitive) equality only.
func matchBySchema(msg string, offsets []pathOffset, feats []features.Feature) []string {
	refs := reSchemaRef.FindAllStringSubmatch(msg, -1)
	if len(refs) == 0 {
		return nil
	}
	var ops []string
	for _, ref := range refs {
		schema := strings.ToLower(ref[1])
		for _, feat := range feats {
			if schema == strings.ToLower(feat.Op) {
				ops = append(ops, feat.Op)
				break
			}
		}
	}
	return ops
}

// matchByPath looks for path keys (e.g. "/workflows/{id}") in the error
// message and maps them to ops via the offset table.
func matchByPath(msg string, offsets []pathOffset) []string {
	var ops []string
	for _, off := range offsets {
		if off.Path != "" && strings.Contains(msg, off.Path) {
			ops = append(ops, off.Op)
		}
	}
	return ops
}

// reQuotedKeyword matches single- or double-quoted words in error messages.
var reQuotedKeyword = regexp.MustCompile(`["'](\w+)["']`)

//ff:func feature=agent type=adapter control=sequence
//ff:what matchByGrep — 에러 메시지 키워드를 YAML 본문에서 grep하여 op 특정

// matchByGrep extracts quoted keywords from the error message, searches them
// in the assembled YAML content, and maps matching lines to ops via the offset
// table. Returns the matched ops and a map of op → relative line number within
// the op's path block.
func matchByGrep(msg string, yamlContent string, offsets []pathOffset) ([]string, map[string]int) {
	keywords := reQuotedKeyword.FindAllStringSubmatch(msg, -1)
	if len(keywords) == 0 {
		return nil, nil
	}

	// Deduplicate keywords.
	kwSet := make(map[string]bool, len(keywords))
	for _, m := range keywords {
		kwSet[m[1]] = true
	}

	// Search YAML content line by line for keyword hits.
	lines := strings.Split(yamlContent, "\n")
	var hitLines []int
	for i, line := range lines {
		for kw := range kwSet {
			if strings.Contains(line, kw) {
				hitLines = append(hitLines, i+1) // 1-based
				break
			}
		}
	}

	if len(hitLines) == 0 {
		return nil, nil
	}

	// Map hit lines to ops via offset ranges.
	seen := make(map[string]bool)
	relativeLines := make(map[string]int)
	for _, lineNo := range hitLines {
		for _, off := range offsets {
			if lineNo >= off.StartLine && lineNo <= off.EndLine {
				if !seen[off.Op] {
					seen[off.Op] = true
					relativeLines[off.Op] = lineNo - off.StartLine
				}
				break
			}
		}
	}

	ops := make([]string, 0, len(seen))
	for op := range seen {
		ops = append(ops, op)
	}
	return ops, relativeLines
}

// buildRetryPrompt builds the user prompt for retrying a single op's path
// block after a kin-openapi validation failure. If relativeLine >= 0, it is
// included to pinpoint the error location within the op's block.
func buildRetryPrompt(feat features.Feature, ddlContent, prevError string, relativeLine int) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Feature:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  table: %s\n", feat.Table)
	fmt.Fprintf(&b, "  public: %v\n", feat.Public)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)

	if ddlContent != "" {
		fmt.Fprintf(&b, "\nTable DDL:\n%s\n", ddlContent)
	}

	fmt.Fprintf(&b, "\nPrevious attempt had this error:\n%s\n", prevError)
	if relativeLine >= 0 {
		fmt.Fprintf(&b, "\nThe error is near line %d of your path block.\n", relativeLine)
	}
	b.WriteString("\nWrite the corrected OpenAPI path block.")
	b.WriteString("\nThe output must be a valid YAML fragment starting with the path key (e.g. /resources/{id}:).")
	b.WriteString("\nInclude operationId, request/response schemas.")
	b.WriteString("\nUse 2-space indentation. No surrounding 'paths:' key. No markdown fences.")

	return b.String()
}
