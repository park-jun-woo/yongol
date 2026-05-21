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
//  2. $ref errors containing "#/components/schemas/..." — match schema name
//     to feature op/path.
//  3. Path key errors containing a path like "/resources/{id}" — match to offset.
//
// Returns an empty slice when the error cannot be attributed to a specific op.
func extractErrorOps(err error, offsets []pathOffset, feats []features.Feature) []string {
	if err == nil {
		return nil
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

	out := make([]string, 0, len(seen))
	for op := range seen {
		out = append(out, op)
	}
	return out
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

// matchBySchema finds schema names in the error and tries to match them to
// feature ops. Matching: the schema name (lowered) is a prefix/substring of
// the op or path, or the feature table name appears in the schema name.
func matchBySchema(msg string, offsets []pathOffset, feats []features.Feature) []string {
	refs := reSchemaRef.FindAllStringSubmatch(msg, -1)
	if len(refs) == 0 {
		return nil
	}
	var ops []string
	for _, ref := range refs {
		schema := strings.ToLower(ref[1])
		for _, feat := range feats {
			table := strings.ToLower(feat.Table)
			op := strings.ToLower(feat.Op)
			if table != "" && strings.Contains(schema, table) {
				ops = append(ops, feat.Op)
				break
			}
			if strings.Contains(schema, op) || strings.Contains(op, schema) {
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

// buildRetryPrompt builds the user prompt for retrying a single op's path
// block after a kin-openapi validation failure.
func buildRetryPrompt(feat features.Feature, ddlContent, prevError string) string {
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
	b.WriteString("\nWrite the corrected OpenAPI path block.")
	b.WriteString("\nThe output must be a valid YAML fragment starting with the path key (e.g. /resources/{id}:).")
	b.WriteString("\nInclude operationId, request/response schemas.")
	b.WriteString("\nUse 2-space indentation. No surrounding 'paths:' key. No markdown fences.")

	return b.String()
}
