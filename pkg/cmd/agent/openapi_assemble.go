//ff:func feature=agent type=helper control=sequence
//ff:what openapi_assemble — OpenAPI path block 기계적 조립 (parameters + requestBody + responses → path block → full document)

package agent

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"gopkg.in/yaml.v3"
)

// httpMethodFromOp derives the HTTP method from the operationId naming convention.
// Returns lowercase method string (e.g. "get", "post", "put", "delete").
func httpMethodFromOp(op string) string {
	lower := strings.ToLower(op)
	switch {
	case strings.HasPrefix(lower, "list"):
		return "get"
	case strings.HasPrefix(lower, "get"):
		return "get"
	case strings.HasPrefix(lower, "create"):
		return "post"
	case strings.HasPrefix(lower, "update"):
		return "put"
	case strings.HasPrefix(lower, "delete"), strings.HasPrefix(lower, "remove"):
		return "delete"
	default:
		return "post"
	}
}

// needsRequestBody returns true if the HTTP method typically carries a request body.
func needsRequestBody(method string) bool {
	return method == "post" || method == "put" || method == "patch"
}

// buildErrorResponses returns hardcoded error response entries based on the feature.
// All errors reference $ref: '#/components/schemas/Error'.
func buildErrorResponses(feat features.Feature) map[string]any {
	errRef := map[string]any{
		"description": "Error",
		"content": map[string]any{
			"application/json": map[string]any{
				"schema": map[string]any{
					"$ref": "#/components/schemas/Error",
				},
			},
		},
	}

	method := httpMethodFromOp(feat.Op)
	resps := make(map[string]any)

	if feat.Public {
		// Public ops (e.g. Login): 400 + 500
		resps["400"] = copyMap(errRef)
		resps["500"] = copyMap(errRef)
	} else if method == "delete" {
		// DELETE: 401 + 403 + 404
		resps["401"] = copyMap(errRef)
		resps["403"] = copyMap(errRef)
		resps["404"] = copyMap(errRef)
	} else {
		// Non-public: 401 + 403 + 404 + 500
		resps["401"] = copyMap(errRef)
		resps["403"] = copyMap(errRef)
		resps["404"] = copyMap(errRef)
		resps["500"] = copyMap(errRef)
	}

	return resps
}

// copyMap performs a shallow copy of a map[string]any.
func copyMap(m map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

// assemblePathBlock combines parameters, requestBody, and response schema into
// a single OpenAPI path block for one operation.
// Returns a map suitable for embedding under the path key in the paths section.
func assemblePathBlock(feat features.Feature, params []any, reqBody map[string]any, schema200 map[string]any, errorResps map[string]any) map[string]any {
	method := httpMethodFromOp(feat.Op)

	op := map[string]any{
		"operationId": feat.Op,
		"summary":     feat.Desc,
	}

	// Security: non-public ops require bearerAuth
	if !feat.Public {
		op["security"] = []map[string][]string{
			{"bearerAuth": {}},
		}
	}

	// Parameters
	if len(params) > 0 {
		op["parameters"] = params
	}

	// Request body
	if reqBody != nil && needsRequestBody(method) {
		op["requestBody"] = reqBody
	}

	// Responses: 200 + errors
	responses := make(map[string]any)
	responses["200"] = map[string]any{
		"description": "OK",
		"content": map[string]any{
			"application/json": map[string]any{
				"schema": schema200,
			},
		},
	}
	for code, resp := range errorResps {
		responses[code] = resp
	}
	op["responses"] = responses

	// method block under the path key
	return map[string]any{method: op}
}

// assembleFullOpenAPI builds the complete openapi.yaml document from Go struct
// path blocks and returns both the YAML string and per-path line offsets for
// error attribution.
// pathBlocks maps path keys (e.g. "/workflows/{id}") to method maps.
// pathToOps maps path keys to originating op names.
func assembleFullOpenAPI(projectName string, pathBlocks map[string]any, pathToOps map[string][]string) (string, []pathOffset) {
	var b strings.Builder
	b.WriteString("openapi: \"3.1.0\"\n")
	b.WriteString("info:\n")
	fmt.Fprintf(&b, "  title: %s\n", projectName)
	b.WriteString("  version: \"1.0.0\"\n")
	b.WriteString("security:\n")
	b.WriteString("  - bearerAuth: []\n")
	b.WriteString("paths:\n")

	// Header above is 7 lines; next line is 8 (1-based).
	currentLine := 8

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

		ops := pathToOps[pathKey]
		if len(ops) == 0 {
			ops = []string{pathKey}
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
