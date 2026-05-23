//ff:func feature=agent type=helper control=selection
//ff:what buildGeneratePrompt — SSaC + feature로 새 블록 생성 프롬프트 구성

package agent

import (
	"fmt"
	"strings"
)

// buildGeneratePrompt creates a prompt for generating new content (not fixing existing).
func buildGeneratePrompt(l layer, operationId, featureDesc, featurePath, ssacContent string) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Feature: %s\nPath: %s\n\n", featureDesc, featurePath)
	fmt.Fprintf(&b, "SSaC file (%s.ssac):\n%s\n\n", operationId, ssacContent)

	switch l {
	case layerOpenAPI:
		fmt.Fprintf(&b, "Generate a new OpenAPI path block for this endpoint.\n")
		fmt.Fprintf(&b, "The block must include the path key (e.g. /resources/{id}:), HTTP method, operationId: %s, and all request/response schemas.\n", operationId)
		fmt.Fprintf(&b, "Use 2-space indentation starting at indent level 2 (4 spaces for path key).\n")
		b.WriteString("Output ONLY the path block content (from the path key line). No surrounding keys like 'paths:'.\n\n")
		b.WriteString("Example:\n")
		b.WriteString(openapiExample)
	case layerRego:
		fmt.Fprintf(&b, "Generate a new Rego allow rule for operationId %s.\n", operationId)
		b.WriteString("The rule must use 'allow if {' syntax with input.action == \"<operationId>\" check.\n")
		b.WriteString("Output ONLY the allow block. No package declaration.\n\n")
		b.WriteString("Example:\n")
		b.WriteString(regoExample)
	case layerHurl:
		fmt.Fprintf(&b, "Generate a new Hurl request block for operationId %s.\n", operationId)
		fmt.Fprintf(&b, "Start with a comment '# %s', then the HTTP method and URL, headers, body, and response assertions.\n", operationId)
		b.WriteString("Output ONLY the request block.\n\n")
		b.WriteString("Example:\n")
		b.WriteString(hurlExample)
	}

	return b.String()
}
