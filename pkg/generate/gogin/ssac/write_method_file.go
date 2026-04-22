//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what writeMethodFile — method Go source 파일 쓰기 (어노테이션 블록 포함)

package ssac

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

func writeMethodFile(dir, funcName, modulePath string, imports []string, sig string, body []string, what string) error {
	seen := make(map[string]bool)
	var deduped []string
	for _, imp := range imports {
		if imp != "" && !seen[imp] {
			seen[imp] = true
			deduped = append(deduped, imp)
		}
	}
	sort.Strings(deduped)

	control := ffannot.DetectControl(body)
	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{
			Feature:   "service",
			Type:      "handler",
			Control:   control,
			Dimension: 1,
		},
		What: funcName + " — " + what,
	})

	var sb strings.Builder
	sb.WriteString(header)
	sb.WriteString("package service\n\nimport (\n")
	for _, imp := range deduped {
		sb.WriteString("\t" + imp + "\n")
	}
	sb.WriteString(")\n\n")
	sb.WriteString(sig + " {\n")
	for _, line := range body {
		if line == "" {
			sb.WriteString("\n")
		} else {
			sb.WriteString("\t" + line + "\n")
		}
	}
	sb.WriteString("}\n")

	fileName := strcase.ToSnake(funcName) + ".go"
	return fffile.WriteIfNotPreserved(filepath.Join(dir, fileName), []byte(sb.String()))
}
