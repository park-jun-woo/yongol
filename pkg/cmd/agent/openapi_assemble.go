//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what assembleFullOpenAPI — path blocks를 완전한 openapi.yaml 문서로 조립

package agent

import (
	"fmt"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func assembleFullOpenAPI(projectName string, pathBlocks map[string]any, pathToOps map[string][]string) (string, []pathOffset) {
	var b strings.Builder
	b.WriteString("openapi: \"3.1.0\"\n")
	b.WriteString("info:\n")
	fmt.Fprintf(&b, "  title: %s\n", projectName)
	b.WriteString("  version: \"1.0.0\"\n")
	b.WriteString("security:\n")
	b.WriteString("  - bearerAuth: []\n")
	b.WriteString("paths:\n")

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
