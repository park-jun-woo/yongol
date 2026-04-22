//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeMainGo — import + body 만으로 cmd/main.go 생성 (top-level funcs 는 별도 파일)

package boot

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// writeMainGo renders the final main.go from deduplicated imports and the
// assembled body, then writes the result to artifactsDir/backend/cmd/main.go.
//
// Phase004: top-level helper functions (envInt, envDuration, …) are no
// longer appended after main(); they are emitted by writeEnvHelperFiles
// as sibling 1-file-1-func files so filefunc F1 passes on cmd/. Imports
// are pruned to those actually referenced in the main() body so the
// helpers' imports (e.g. "strconv") don't leak into main.go.
func writeMainGo(artifactsDir string, imports, body []string) error {
	outDir := filepath.Join(artifactsDir, "backend", "cmd")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	bodyText := strings.Join(body, "\n")
	used := filterImportsUsed(imports, bodyText, true)

	control := ffannot.DetectControl(body)
	sb := strings.Builder{}
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature:   "main",
			Type:      "command",
			Control:   control,
			Dimension: 1,
		},
		What: "main — 애플리케이션 엔트리포인트 (DB/JWT/authz/queue/cache/session/file/router/gin 초기화)",
	}))
	sb.WriteString("package main\n\nimport (\n")
	for _, imp := range used {
		sb.WriteString("\t" + imp + "\n")
	}
	sb.WriteString(")\n\nfunc main() {\n")
	for _, line := range body {
		if line == "" {
			sb.WriteString("\n")
		} else {
			sb.WriteString("\t" + line + "\n")
		}
	}
	sb.WriteString("}\n")

	return os.WriteFile(filepath.Join(outDir, "main.go"), []byte(sb.String()), 0o644)
}
