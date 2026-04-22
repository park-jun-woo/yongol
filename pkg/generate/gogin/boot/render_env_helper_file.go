//ff:func feature=gen-gogin type=generator control=sequence
//ff:what renderEnvHelperFile — 환경변수 파싱 헬퍼 1개를 완전한 Go source 로 조립

package boot

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// renderEnvHelperFile produces a full Go source for one helper file.
// Imports are filtered by a substring match against the helper body so
// each file only pulls the packages it uses, keeping goimports happy.
// Control type is detected from the helper body via ffannot.DetectControl so
// helpers that carry switch/for at depth 1 get the correct annotation
// (filefunc A10~A14).
func renderEnvHelperFile(name, body string, imports []string) string {
	bodyLines := extractHelperBodyLines(body)
	control := ffannot.DetectControl(bodyLines)
	annot := ffannot.FuncAnnot{
		Feature: "main",
		Type:    "util",
		Control: control,
	}
	if control == ffannot.ControlIteration {
		annot.Dimension = 1
	}
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: annot,
		What: name + " — 환경변수 파싱 헬퍼 (실패 시 default 반환)",
	}))
	sb.WriteString("package main\n\n")
	used := filterImportsUsed(imports, body, false)
	if len(used) > 0 {
		sb.WriteString("import (\n")
		for _, imp := range used {
			sb.WriteString("\t" + imp + "\n")
		}
		sb.WriteString(")\n\n")
	}
	sb.WriteString(body)
	if !strings.HasSuffix(body, "\n") {
		sb.WriteString("\n")
	}
	return sb.String()
}
